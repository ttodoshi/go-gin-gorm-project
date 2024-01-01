package postgres

import (
	"fmt"
	"goGinGormProject/internal/core/domain"
	"goGinGormProject/internal/core/errors"
	"goGinGormProject/internal/core/ports"
	"goGinGormProject/pkg/database"
	"gorm.io/gorm"
)

type postRepository struct {
	DB *gorm.DB
}

func NewPostRepository() ports.PostRepository {
	return &postRepository{DB: database.ConnectToDb()}
}

func (p *postRepository) GetPostByUUID(uuid string) (post domain.Post, err error) {
	p.DB.First(&post, "uuid = ?", uuid)

	if post.UUID != "" {
		return post, nil
	} else {
		return post, &errors.NotFoundError{
			Message: fmt.Sprintf(`post by uuid '%s' not found`, uuid),
		}
	}
}

func (p *postRepository) GetPosts() (posts []domain.Post) {
	p.DB.Find(&posts)
	return
}

func (p *postRepository) CreatePost(post domain.Post) (uuid string, err error) {
	result := p.DB.Create(&post)

	if result.Error != nil {
		return "", fmt.Errorf(`post %v not created due to error: %v`, post, result.Error)
	}
	return post.UUID, nil
}

func (p *postRepository) UpdatePostByUUID(uuid string, post domain.Post) (updatedPost domain.Post, err error) {
	p.DB.First(&updatedPost, "uuid = ?", uuid)

	if updatedPost.UUID == "" {
		return updatedPost, &errors.NotFoundError{
			Message: fmt.Sprintf(`post by uuid '%s' not found`, uuid),
		}
	}

	p.DB.Model(&updatedPost).Updates(post)

	return updatedPost, nil
}

func (p *postRepository) DeletePostByUUID(uuid string) error {
	var post domain.Post

	p.DB.First(&post, "uuid = ?", uuid)

	if post.UUID != "" {
		p.DB.Delete(&domain.Post{}, "uuid = ?", uuid)
		return nil
	} else {
		return &errors.NotFoundError{
			Message: fmt.Sprintf(`post by uuid '%s' not found`, uuid),
		}
	}
}
