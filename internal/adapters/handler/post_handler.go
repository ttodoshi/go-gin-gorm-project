package handler

import (
	"github.com/gin-gonic/gin"
	"goGinGormProject/internal/adapters/dto"
	"goGinGormProject/internal/core/errors"
	"goGinGormProject/internal/core/ports"
	"goGinGormProject/pkg/logging"
)

type PostHandler struct {
	svc ports.PostService
	log logging.Logger
}

func NewPostHandler(svc ports.PostService, log logging.Logger) *PostHandler {
	return &PostHandler{
		svc: svc,
		log: log,
	}
}

func (h *PostHandler) GetPostByUUID(c *gin.Context) {
	h.log.Debug("received get post by uuid request")

	uuid := c.Param("uuid")

	post, err := h.svc.GetPostByUUID(uuid)

	if err != nil {
		err = c.Error(err)
		return
	}
	c.JSON(200, post)
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	h.log.Debug("received get posts request")

	posts, err := h.svc.GetPosts()
	if err != nil {
		err = c.Error(err)
		return
	}
	c.JSON(200, posts)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	h.log.Debug("received create post request")

	var createPostDto dto.CreatePostDto

	err := c.ShouldBindJSON(&createPostDto)
	if err != nil {
		err = c.Error(&errors.BodyMappingError{Message: "error while mapping body"})
		h.log.Warn("error while mapping body")
		return
	}

	postUUID, err := h.svc.CreatePost(createPostDto)

	if err != nil {
		err = c.Error(err)
		return
	}

	c.JSON(201, gin.H{
		"uuid": postUUID,
	})
}

func (h *PostHandler) UpdatePostByUUID(c *gin.Context) {
	h.log.Debug("received update post request")

	uuid := c.Param("uuid")

	var updatePostDto dto.UpdatePostDto

	err := c.ShouldBindJSON(&updatePostDto)
	if err != nil {
		err = c.Error(&errors.BodyMappingError{Message: "error while mapping body"})
		h.log.Warn("error while mapping body")
		return
	}

	updatedPost, err := h.svc.UpdatePostByUUID(uuid, updatePostDto)
	if err != nil {
		err = c.Error(err)
		return
	}
	c.JSON(200, updatedPost)
}

func (h *PostHandler) DeletePostByUUID(c *gin.Context) {
	h.log.Debug("received delete post request")

	uuid := c.Param("uuid")

	err := h.svc.DeletePostByUUID(uuid)
	if err != nil {
		err = c.Error(err)
		return
	}
	c.JSON(204, nil)
}
