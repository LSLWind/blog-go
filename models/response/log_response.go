package response

import "time"

type LogResponse struct {
	Content    string    `json:"content"`
	UpdateTime time.Time `json:"updateTime"`
}
