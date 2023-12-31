package handler

import (
	"github.com/gin-gonic/gin"
	"goGinGormProject/internal/core/domain"
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
	c.JSON(200, h.svc.GetPosts())
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var body struct { // TODO: Dto
		Body  string
		Title string
	}

	err := c.Bind(&body)
	if err != nil {
		log.Print("error mapping body")
	}

	newPost := domain.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	postUUID, err := h.svc.CreatePost(newPost)

	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(201, gin.H{
		"UUID": postUUID,
	})
}

func (h *PostHandler) UpdatePostByUUID(c *gin.Context) {
	uuid := c.Param("uuid")

	var body struct { // TODO: Dto
		Body  string
		Title string
	}

	err := c.Bind(&body)
	if err != nil {
		log.Print("error mapping body")
	}

	post := domain.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	updatedPost, err := h.svc.UpdatePostByUUID(uuid, post)
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
