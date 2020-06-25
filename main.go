package main

import (
	"github.com/chillicoder/bookstore/controllers"
	"github.com/chillicoder/bookstore/models"
	"github.com/gin-gonic/gin"
	_ "net/http"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
