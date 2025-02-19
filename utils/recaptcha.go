package utils

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
)


const recaptchaSecretKey = ""

type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes,omitempty"`
}

// VerifyRecaptcha kiểm tra token reCAPTCHA
func VerifyRecaptcha(token string) (bool, float64, error) {
	client := resty.New()
	resp, err := client.R().
		SetFormData(map[string]string{
			"secret":   recaptchaSecretKey,
			"response": token,
		}).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Post("https://www.google.com/recaptcha/api/siteverify")

	if err != nil {
		return false, 0, err
	}

	var result RecaptchaResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return false, 0, err
	}

	if !result.Success {
		return false, 0, errors.New("failed to verify reCAPTCHA")
	}

	return true, result.Score, nil
}
