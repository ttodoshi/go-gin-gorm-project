package servises

import (
	"github.com/jinzhu/copier"
	"goGinGormProject/internal/adapters/dto"
	"goGinGormProject/internal/core/domain"
	"goGinGormProject/internal/core/errors"
	"goGinGormProject/internal/core/ports"
	"goGinGormProject/pkg/logging"
)

type PostService struct {
	repo ports.PostRepository
	log  logging.Logger
}

func NewPostService(repo ports.PostRepository, log logging.Logger) ports.PostService {
	return &PostService{
		repo: repo,
		log:  log,
	}
}

func (s *PostService) GetPostByUUID(uuid string) (getPostDto dto.GetPostDto, err error) {
	post, err := s.repo.GetPostByUUID(uuid)
	if err != nil {
		s.log.Infof(`error getting post by uuid: '%s' due to error: %v`, uuid, err)
		return getPostDto, err
	}

	err = copier.Copy(&getPostDto, &post)
	if err != nil {
		return getPostDto, &errors.MappingError{Message: `struct mapping error`}
	}
	return getPostDto, nil
}

func (s *PostService) GetPosts() (postsDto []dto.GetPostDto, err error) {
	posts := s.repo.GetPosts()
	err = copier.Copy(&postsDto, &posts)
	if err != nil {
		return postsDto, &errors.MappingError{Message: `struct mapping error`}
	}
	return postsDto, nil
}

func (s *PostService) CreatePost(createPostDto dto.CreatePostDto) (uuid string, err error) {
	var post domain.Post
	err = copier.Copy(&post, &createPostDto)
	if err != nil {
		return "", &errors.MappingError{Message: `struct mapping error`}
	}

	uuid, err = s.repo.CreatePost(post)
	if err != nil {
		s.log.Infof("%v", err)
		return "", err
	}
	return uuid, nil
}

func (s *PostService) UpdatePostByUUID(uuid string, updatePostDto dto.UpdatePostDto) (getPostDto dto.GetPostDto, err error) {
	var post domain.Post
	err = copier.Copy(&post, &updatePostDto)
	if err != nil {
		return getPostDto, &errors.MappingError{Message: `struct mapping error`}
	}

	var updatedPost domain.Post
	updatedPost, err = s.repo.UpdatePostByUUID(uuid, post)
	if err != nil {
		s.log.Infof(`error updating post by uuid: '%s' due to error: %v`, uuid, err)
		return getPostDto, err
	}

	err = copier.Copy(&getPostDto, &updatedPost)
	if err != nil {
		return getPostDto, &errors.MappingError{Message: `struct mapping error`}
	}

	return getPostDto, nil
}

func (s *PostService) DeletePostByUUID(uuid string) error {
	err := s.repo.DeletePostByUUID(uuid)
	if err != nil {
		s.log.Infof(`error deleting post by uuid: '%s' due to error: %v`, uuid, err)
		return err
	}
	return nil
}
