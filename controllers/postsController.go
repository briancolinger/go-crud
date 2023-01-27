package controllers

import (
	"github.com/briancolinger/go-crud/initializers"
	"github.com/briancolinger/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data from request body.
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post.
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it.
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts.
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get the id.
	id := c.Param("id")

	// Get the posts.
	var post []models.Post
	initializers.DB.First(&post, id)

	// Return them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id.
	id := c.Param("id")

	// Get data from request body.
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Get the post.
	var post []models.Post
	initializers.DB.First(&post, id)

	// Update the post.
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Return.
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id.
	id := c.Param("id")

	// Delete the post.
	initializers.DB.Delete(&models.Post{}, id)

	// Respond.
	c.Status(200)
}
