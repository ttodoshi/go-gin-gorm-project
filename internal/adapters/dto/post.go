package dto

import (
	"time"
)

type PostDto struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type GetPostDto struct {
	UUID string `json:"uuid"`
	PostDto
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type CreatePostDto struct {
	PostDto
}
type UpdatePostDto struct {
	PostDto
}
