package routers

import (
	"rest-api-books-gin-gorm/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)

	router.PUT("/books/:bookID", controllers.UpdateBook)

	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:bookID", controllers.GetBook)


	router.DELETE("books/:bookID", controllers.DeleteBook)
	return router
}