package handler

import (
	"github.com/gin-gonic/gin"
	"goGinGormProject/internal/adapters/dto"
	"goGinGormProject/internal/core/ports"
	"log"
)

type PostHandler struct {
	svc ports.PostService
}

func NewPostHandler(svc ports.PostService) *PostHandler {
	return &PostHandler{svc: svc}
}

func (h *PostHandler) GetPostByUUID(c *gin.Context) {
	uuid := c.Param("uuid")

	post, err := h.svc.GetPostByUUID(uuid)

	if err != nil {
		c.Status(404)
		return
	}
	c.JSON(200, post)
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	posts, err := h.svc.GetPosts()
	if err != nil {
		c.Status(500)
		return
	}
	c.JSON(200, posts)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var createPostDto dto.CreatePostDto

	err := c.Bind(&createPostDto)
	if err != nil {
		log.Print("error mapping body")
	}

	postUUID, err := h.svc.CreatePost(createPostDto)

	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(201, gin.H{
		"uuid": postUUID,
	})
}

func (h *PostHandler) UpdatePostByUUID(c *gin.Context) {
	uuid := c.Param("uuid")

	var updatePostDto dto.UpdatePostDto

	err := c.Bind(&updatePostDto)
	if err != nil {
		log.Print("error mapping body")
	}

	updatedPost, err := h.svc.UpdatePostByUUID(uuid, updatePostDto)
	if err != nil {
		c.Status(404)
		return
	}
	c.JSON(200, updatedPost)
}

func (h *PostHandler) DeletePostByUUID(c *gin.Context) {
	uuid := c.Param("uuid")

	err := h.svc.DeletePostByUUID(uuid)
	if err != nil {
		c.Status(404)
		return
	}
	c.JSON(204, nil)
}
