package middleware

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"time"
	"path/filepath"
	"os"
	"strings"

	"golang_template_source/config"
	"golang_template_source/domain"

	"github.com/gin-gonic/gin"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // Capture response body
	return w.ResponseWriter.Write(b)
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Clone request body for logging
		var requestBody string
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			requestBody = string(bodyBytes)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Restore request body for further processing
		}

		// Clone query parameters
		requestQuery, _ := json.Marshal(c.Request.URL.Query())

		// Wrap the response writer to capture response body
		bw := &bodyWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = bw

		// Process the request
		c.Next()

		// Capture response details
		statusCode := c.Writer.Status()
		responseBody := bw.body.String()
		duration := time.Since(startTime).Seconds()
		realIP := GetRealIP(c)

		// Save log into SYS_LOG table
		log := domain.SysLog{
			ActionDatetime:   startTime,
			PathName:         c.Request.URL.Path,
			Method:           c.Request.Method,
			IP:               realIP,
			StatusResponse:   statusCode,
			Response:         responseBody,
			Description:      "Request logged by middleware",
			RequestBody:      requestBody,
			RequestQuery:     string(requestQuery),
			Duration:         duration,
		}
		fmt.Println( "Request logged by middleware")
		if err := config.DB.Create(&log).Error; err != nil {
			c.Error(err) // Log the error if saving the log fails
		}

		if err := writeLogToCSV(log); err != nil {
			fmt.Printf("Error writing log to CSV: %v\n", err)
		}
	}
}


func writeLogToCSV(log domain.SysLog) error {
	// Get the current week number and year
	year, week := time.Now().ISOWeek()
	filename := filepath.Join("logs", fmt.Sprintf("log_week_%d_%02d.csv", year, week))

	// Ensure the logs directory exists
	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create logs directory: %w", err)
	}

	// Check if the file exists
	fileExists := false
	if _, err := os.Stat(filename); err == nil {
		fileExists = true
	}

	// Open the file in append mode
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header if the file is new
	if !fileExists {
		if err := writer.Write([]string{
			"action_datetime", "path_name", "method", "ip",
			"status_response", "response", "description",
			"request_body", "request_query", "duration",
		}); err != nil {
			return fmt.Errorf("failed to write CSV header: %w", err)
		}
	}
	
	// Write the log entry
	if err := writer.Write([]string{
		log.ActionDatetime.Format("2006-01-02 15:04:05"),
		log.PathName,
		log.Method,
		log.IP,
		fmt.Sprintf("%d", log.StatusResponse),
		log.Response,
		log.Description,
		log.RequestBody,
		log.RequestQuery,
		fmt.Sprintf("%.3f", log.Duration),
	}); err != nil {
		return fmt.Errorf("failed to write log entry to CSV: %w", err)
	}

	return nil
}

func GetRealIP(c *gin.Context) string {
	// Headers to check
	headersToCheck := []string{
		"X-Forwarded-For",
		"X-Real-IP",
	}

	for _, header := range headersToCheck {
		if value := c.GetHeader(header); value != "" {
			// If the header exists, take the first IP from the comma-separated list
			ips := strings.Split(value, ",")
			return strings.TrimSpace(ips[0])
		}
	}

	// Fallback to the client's remote IP address
	clientIP := c.ClientIP()
	return clientIP
}