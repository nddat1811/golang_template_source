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
	"golang_template_source/repository"

	"github.com/gin-gonic/gin"
)

// bodyWriter is a custom response writer that captures the response body
type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write captures the response body while allowing it to be written normally
func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // Capture response body
	return w.ResponseWriter.Write(b)
}

// containsSensitivePath checks if the request path contains sensitive keywords
func containsSensitivePath(path string) bool {
	// List of sensitive paths that should have masked logging
	sensitivePaths := []string{"login", "signin", "password-reset", "register"}

	// Check if the request path contains any of the sensitive keywords
	for _, keyword := range sensitivePaths {
		if strings.Contains(path, keyword) {
			return true
		}
	}
	return false
}

// maskPassword replaces the password field in the request body with asterisks
func maskPassword(body string) string {
	// Parse JSON request body to replace the password field
	var requestData map[string]interface{}
	if err := json.Unmarshal([]byte(body), &requestData); err != nil {
		return body // If JSON parsing fails, return the original body
	}

	// If the "password" field exists, replace its value with asterisks
	if _, exists := requestData["password"]; exists {
		requestData["password"] = "******"
	}

	// Convert the modified JSON object back to a string
	maskedBody, err := json.Marshal(requestData)
	if err != nil {
		return body // If JSON encoding fails, return the original body
	}

	return string(maskedBody)
}

// getRealIP retrieves the real IP address of the client making the request
func getRealIP(c *gin.Context) string {
	// Headers to check for the client's real IP
	headersToCheck := []string{
		"X-Forwarded-For",
		"X-Real-IP",
	}

	// Check if any of the headers contain an IP address
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

// writeLogToCSV writes log entries to a CSV file based on the current week and year
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

	// Write the CSV header if the file is new
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

// LogMiddleware is a Gin middleware for logging HTTP requests and responses
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/swagger/") {
            c.Next() // Not log swagger
            return
        }

		startTime := time.Now()

		// Check Content-Type to determine if the request contains file upload
		contentType := c.Request.Header.Get("Content-Type")
		var requestBody string

		if !strings.HasPrefix(contentType, "multipart/form-data") && c.Request.Body != nil {
			// Clone request body if it's not a file upload
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			requestBody = string(bodyBytes)

			// Mask password if the request path contains sensitive information
			if containsSensitivePath(c.Request.URL.Path) {
				requestBody = maskPassword(requestBody)
			}

			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Restore request body for further processing
		} else {
			requestBody = "file upload"
		}

		// Clone query parameters for logging
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
		if strings.Contains(c.Writer.Header().Get("Content-Type"), "application/octet-stream") {
			responseBody = "file export"
		}
		duration := time.Since(startTime).Seconds()
		realIP := getRealIP(c)

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

		// Initialize database connection
		con := config.InitPostgreSQL()
		defer config.CloseConnectDB(con)
		sysLogRepo := repository.NewSysLogRepository(con)

		// Insert log into the database
		err := sysLogRepo.InsertLog(&log)
		if err != nil {
			c.Error(err)
		}

		// Write log to CSV file
		if err := writeLogToCSV(log); err != nil {
			fmt.Printf("Error writing log to CSV: %v\n", err)
		}
	}
}

