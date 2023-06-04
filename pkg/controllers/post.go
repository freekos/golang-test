package controllers

import (
	"freekos/crud/pkg/services"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	services.CreatePost(c)
}

func FindPosts(c *gin.Context) {
	services.FindPosts(c)
}

func FindPost(c *gin.Context) {
	services.FindPost(c)
}

func UpdatePost(c *gin.Context) {
	services.UpdatePost(c)
}
