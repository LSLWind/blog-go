package response

import "time"

type LogResponse struct {
	Content    string    `json:"content"`
	UpdateTime time.Time `json:"timestamp"`
	Type       int       `json:"type"`
}
