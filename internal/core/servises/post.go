package servises

import (
	"goGinGormProject/internal/core/domain"
	"goGinGormProject/internal/core/ports"
	"log"
)

type PostService struct {
	repo ports.PostRepository
}

func NewPostService(repo ports.PostRepository) ports.PostService {
	return &PostService{repo: repo}
}

func (p *PostService) GetPostByUUID(uuid string) (post domain.Post, err error) {
	post, err = p.repo.GetPostByUUID(uuid)
	if err != nil {
		log.Printf(`error getting post by uuid: "%s" due to error: %v`, uuid, err)
		return post, err
	}
	return post, nil
}

func (p *PostService) GetPosts() []domain.Post {
	return p.repo.GetPosts()
}

func (p *PostService) CreatePost(post domain.Post) (uuid string, err error) {
	uuid, err = p.repo.CreatePost(post)
	if err != nil {
		log.Printf("%v", err)
		return uuid, err
	}
	return uuid, nil
}

func (p *PostService) UpdatePostByUUID(uuid string, post domain.Post) (updatedPost domain.Post, err error) {
	updatedPost, err = p.repo.UpdatePostByUUID(uuid, post)
	if err != nil {
		log.Printf(`error updating post by uuid: "%s" due to error: %v`, uuid, err)
		return updatedPost, err
	}
	return updatedPost, nil
}

func (p *PostService) DeletePostByUUID(uuid string) error {
	err := p.repo.DeletePostByUUID(uuid)
	if err != nil {
		log.Printf(`error deleting post by uuid: "%s" due to error: %v`, uuid, err)
		return err
	}
	return nil
}
