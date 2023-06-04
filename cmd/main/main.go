package main

import (
	"fmt"
	"freekos/crud/configs"
	"freekos/crud/pkg/controllers"
	"freekos/crud/pkg/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.FindPosts)
	r.GET("/posts/:id", controllers.FindPost)
	r.PUT("/posts", controllers.UpdatePost)

	r.Run(fmt.Sprintf("%v:%v", configs.Config.HOST, configs.Config.PORT))
}
