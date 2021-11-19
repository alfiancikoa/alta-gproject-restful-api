package category

type PostCategory struct {
	Title string `json:"title" form:"title"`
}

type EditCategory struct {
	Title string `json:"title" form:"title"`
}
