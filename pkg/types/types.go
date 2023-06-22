package types

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title       string `json:"title" binding:"min=2,max=20" validate:"is-title-ok" gorm:"unique;not null"`
	Description string `json:"description"`
	URL         string `json:"url" binding:"required,url"`
	AuthorID    uint   `json:"author_id"`
	Author      Author `json:"author,omitempty" gorm:"foreignKey:AuthorID;references:ID"`
}

type Author struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique;not null"`
}

type GetAuthorRequest struct {
	ID uint `json:"id" bind:"required"`
}
