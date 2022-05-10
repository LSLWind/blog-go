package response

type CategoryResponse struct {
	Id              uint          `json:"id"`
	Name            string        `json:"name"`
	SvgIcon         string        `json:"svgIcon"`
	SubCategoryList []SubCategory `json:"subCategoryList"`
}
type SubCategory struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
