package controllers

import (
	"errors"
	"ignbmd/blogging-platform-api/internal/app/models"
	"ignbmd/blogging-platform-api/internal/app/services"
	"ignbmd/blogging-platform-api/internal/app/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostController struct {
	service *services.PostService
}

func NewPostController(service *services.PostService) *PostController {
	return &PostController{
		service: service,
	}
}

func (ctl *PostController) FindAll(c *gin.Context) {
	searchTerm := c.Query("term")

	posts, err := ctl.service.FindAll(c, searchTerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "data": nil, "message": "Failed to retrieved posts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": posts, "message": "Posts retrieved"})
}

func (cli *PostController) FindByID(c *gin.Context) {
	id := c.Param("id")
	post, err := cli.service.FindByID(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "data": nil, "message": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": post, "message": "Post found"})
}

func (ctl *PostController) Create(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "data": nil, "message": "Failed to parse JSON body payload"})
		return
	}

	if err := ctl.service.Create(c, &post); err != nil {
		var validationErr *validators.ValidationError
		if errors.As(err, &validationErr) {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "data": nil, "message": "Validation failed", "errors": validationErr.Errors})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "data": nil, "message": "Failed to create a post"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": "true", "data": post, "message": "Post created"})
}

func (ctl *PostController) Update(c *gin.Context) {
	id := c.Param("id")

	var updatedData models.Post
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "data": nil, "message": "Failed to parse JSON body payload"})
		return
	}

	updatedPost, err := ctl.service.Update(c, id, &updatedData)
	if err != nil {
		var validationErr *validators.ValidationError
		if errors.As(err, &validationErr) {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "data": nil, "message": "Validation failed", "errors": validationErr.Errors})
			return
		}

		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "data": nil, "message": "Post not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "data": nil, "message": "Failed to update a post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": updatedPost, "message": "Post updated successfully"})
}

func (ctl *PostController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := ctl.service.Delete(c, id); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "data": nil, "message": "Post not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "data": nil, "message": "Failed to delete a post"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
