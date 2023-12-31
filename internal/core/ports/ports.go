package ports

import (
	"goGinGormProject/internal/adapters/dto"
	"goGinGormProject/internal/core/domain"
)

type PostService interface {
	GetPostByUUID(uuid string) (dto.GetPostDto, error)
	GetPosts() ([]dto.GetPostDto, error)
	CreatePost(post dto.CreatePostDto) (uuid string, err error)
	UpdatePostByUUID(uuid string, post dto.UpdatePostDto) (dto.GetPostDto, error)
	DeletePostByUUID(uuid string) error
}

type PostRepository interface {
	GetPostByUUID(uuid string) (domain.Post, error)
	GetPosts() []domain.Post
	CreatePost(post domain.Post) (uuid string, err error)
	UpdatePostByUUID(uuid string, post domain.Post) (domain.Post, error)
	DeletePostByUUID(uuid string) error
}
