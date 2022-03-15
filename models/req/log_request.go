package req

type LogRequest struct {
	Content string `json:"content"`
	Type    int    `json:"type"`
}
