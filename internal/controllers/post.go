package controllers

import (
	"github.com/gin-gonic/gin"
	"goGinGormProject/internal/models"
	"goGinGormProject/pkg/initializers"
	"log"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	err := c.Bind(&body)
	if err != nil {
		log.Print("error mapping body")
	}

	newPost := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := initializers.DB.Create(&newPost)

	if result.Error != nil {
		c.Status(500)
		log.Printf("error creating post with body %v", body)
		return
	}
	c.JSON(201, gin.H{
		"post": newPost,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPostByID(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	initializers.DB.First(&post, id)

	if post.ID != 0 {
		c.JSON(200, gin.H{
			"post": post,
		})
	} else {
		c.Status(404)
		log.Printf("error getting post by id %s", id)
	}
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	err := c.Bind(&body)
	if err != nil {
		log.Print("error mapping body")
	}

	var post models.Post

	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(
		models.Post{
			Title: body.Title,
			Body:  body.Body,
		},
	)

	if post.ID != 0 {
		c.JSON(200, gin.H{
			"post": post,
		})
	} else {
		c.Status(404)
		log.Printf("error getting post by id %s", id)
	}
}

func DeletePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	initializers.DB.First(&post, id)

	if post.ID != 0 {
		initializers.DB.Delete(&models.Post{}, id)

		c.JSON(204, nil)
	} else {
		c.Status(404)
		log.Printf("error getting post by id %s", id)
	}
}
