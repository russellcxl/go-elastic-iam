package types

type Video struct {
	Title       string `json:"title" binding:"min=2,max=20"`
	Description string `json:"description"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"required,email"`
}
