package response

type NavigationInfoResponse struct {
	Id      uint   `json:"id"`
	Type    string `json:"type"`
	LinkUrl string `json:"link_url"`
	Name    string `json:"name"`
}
