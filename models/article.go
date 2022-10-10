package models

import "time"


type Person struct {
	Firstname string  `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"John"`
	Lastname  string  `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Doe"`
}

type Content struct {
	Title string	`json:"title" binding:"required"`
	Body  string	`json:"body" binding:"required"`
}

type Article struct {
	ID          string `json:"id"`
	Content          
	AuthorID    string `json:"author_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt *time.Time `json:"updated_at"`
	DeleteAt *time.Time `json:"-"`
}	
type CreateArticleModel struct {
	Content          // Promoted fields
	AuthorID    string `json:"author_id" binding:"required"`
}

//PackedArticleModel...
type PackedArticleModel struct {
	ID        string `json:"id"`
	Content          
	Author    Author `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt *time.Time `json:"updated_at"`
	DeleteAt *time.Time `json:"-"`
}	
