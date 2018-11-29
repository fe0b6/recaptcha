package recaptcha

import (
	"time"
)

// Ans - ответ гугла
type Ans struct {
	Success     bool          `json:"success"`
	Score       float64       `json:"score"`
	Action      string        `json:"action"`
	ChallengeTS time.Time     `json:"challenge_ts"`
	Hostname    string        `json:"hostname"`
	ErrorCodes  []interface{} `json:"error-codes"`
}
