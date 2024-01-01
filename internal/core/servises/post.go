package servises

import (
	"github.com/jinzhu/copier"
	"goGinGormProject/internal/adapters/dto"
	"goGinGormProject/internal/core/domain"
	"goGinGormProject/internal/core/errors"
	"goGinGormProject/internal/core/ports"
	"log"
)

type PostService struct {
	repo ports.PostRepository
}

func NewPostService(repo ports.PostRepository) ports.PostService {
	return &PostService{repo: repo}
}

func (p *PostService) GetPostByUUID(uuid string) (getPostDto dto.GetPostDto, err error) {
	post, err := p.repo.GetPostByUUID(uuid)
	if err != nil {
		log.Printf(`error getting post by uuid: '%s' due to error: %v`, uuid, err)
		return getPostDto, err
	}

	err = copier.Copy(&getPostDto, &post)
	if err != nil {
		return getPostDto, &errors.MappingError{Message: `struct mapping error`}
	}
	return getPostDto, nil
}

func (p *PostService) GetPosts() (postsDto []dto.GetPostDto, err error) {
	posts := p.repo.GetPosts()
	err = copier.Copy(&postsDto, &posts)
	if err != nil {
		return postsDto, &errors.MappingError{Message: `struct mapping error`}
	}
	return postsDto, nil
}

func (p *PostService) CreatePost(createPostDto dto.CreatePostDto) (uuid string, err error) {
	var post domain.Post
	err = copier.Copy(&post, &createPostDto)
	if err != nil {
		return "", &errors.MappingError{Message: `struct mapping error`}
	}

	uuid, err = p.repo.CreatePost(post)
	if err != nil {
		log.Printf("%v", err)
		return "", err
	}
	return uuid, nil
}

func (p *PostService) UpdatePostByUUID(uuid string, updatePostDto dto.UpdatePostDto) (getPostDto dto.GetPostDto, err error) {
	var post domain.Post
	err = copier.Copy(&post, &updatePostDto)
	if err != nil {
		return getPostDto, &errors.MappingError{Message: `struct mapping error`}
	}

	var updatedPost domain.Post
	updatedPost, err = p.repo.UpdatePostByUUID(uuid, post)
	if err != nil {
		log.Printf(`error updating post by uuid: '%s' due to error: %v`, uuid, err)
		return getPostDto, err
	}

	err = copier.Copy(&getPostDto, &updatedPost)
	if err != nil {
		return getPostDto, &errors.MappingError{Message: `struct mapping error`}
	}

	return getPostDto, nil
}

func (p *PostService) DeletePostByUUID(uuid string) error {
	err := p.repo.DeletePostByUUID(uuid)
	if err != nil {
		log.Printf(`error deleting post by uuid: '%s' due to error: %v`, uuid, err)
		return err
	}
	return nil
}
