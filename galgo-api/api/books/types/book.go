package types

type AddBookRequestBody struct {
	Title   string `json:"title" form:"title" xml:"title" binding:"required"`
	Content string `json:"content" form:"content" xml:"content" binding:"required"`
}

type UpdateBookRequestBody struct {
	Title   string `json:"title" form:"title" xml:"title" binding:"required"`
	Content string `json:"content" form:"content" xml:"content" binding:"required"`
}

type ParamIdBook struct {
	ID string `uri:"id" binding:"required"`
}
