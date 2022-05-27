package response

// 返回的svg icon信息
type SvgIconInfo struct {
	Name   string `json:"name"`
	SvgStr string `json:"svgStr"`
	Size   uint   `json:"size"`
}
