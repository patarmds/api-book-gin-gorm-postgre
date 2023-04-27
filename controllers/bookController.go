package controllers
 
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"rest-api-books-gin-gorm/database"
	"rest-api-books-gin-gorm/models"
)

type Book = models.Book

func CreateBook(ctx *gin.Context){
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()
	Book := models.Book{
		NameBook: newBook.NameBook,
		Author: newBook.Author,
	}
	err := db.Create(&Book).Error

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, Book)
}

func UpdateBook(ctx *gin.Context){
	bookID := ctx.Param("bookID")
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()
	Book := models.Book{}
	// res := db.Model(&Book).Where("id = ?", bookID).Updates(models.Book{NameBook: updatedBook.NameBook, Author: updatedBook.Author})
	db.First(&Book, "id = ?", bookID)
	Book.NameBook = updatedBook.NameBook
	Book.Author = updatedBook.Author
	res := db.Save(&Book)
	if res.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, res.Error)
		return
	}

	rowsAffected := res.RowsAffected

	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, Book)
}

func GetBook(ctx *gin.Context){
	bookID := ctx.Param("bookID")
	db := database.GetDB()
	
	Book := models.Book{}

	err := db.First(&Book, "id = ?", bookID).Error
	
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, Book)
}

func GetBooks(ctx *gin.Context){
	db := database.GetDB()
 
	var books []Book
	res := db.Find(&books)
	// rows, err := db.Query("SELECT * FROM books")
	if res.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, res.Error)
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func DeleteBook(ctx *gin.Context){
	bookID := ctx.Param("bookID")
	
	db := database.GetDB()

	Book := models.Book{}

	res := db.Where("id = ?", bookID).Delete(&Book);


	if res.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, res.Error)
		return
	}

	rowsAffected := res.RowsAffected

	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message" : "Book deleted successfully",
	})

}


