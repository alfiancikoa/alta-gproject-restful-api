package photo

type PostPhoto struct {
	Title      string `json:"title" form:"title"`
	Size       string `json:"size" form:"size"`
	Resolution string `json:"resolution" form:"resolution"`
	Link       string `json:"link" form:"link"`
	Product_ID uint
}
