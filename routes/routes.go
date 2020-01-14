package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kishankpatel/simplify/db"
	"github.com/kishankpatel/simplify/models"
	"github.com/kishankpatel/simplify/services"
)

// Handler - handles routes
func Handler() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.templ.html")
	bookService, err := services.NewBookService()
	if err != nil {
		fmt.Println(bookService)
		panic(err)
	}
	db := db.InitDb()

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.templ.html", gin.H{
			"name": "Kishan",
		})
	})

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Ok",
		})
	})

	router.POST("/books", func(c *gin.Context) {

		params := &models.Book{}
		c.BindJSON(params)
		book := models.NewBook(params)
		bookRes := bookService.AddNewBook(book, db)

		c.JSON(http.StatusOK, bookRes)
	})

	router.GET("/books", func(c *gin.Context) {
		res := bookService.AllBooks(db)
		c.JSON(http.StatusOK, res)
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusBadRequest, "404.templ.html", gin.H{})
		// c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router.Run(":8080")
}
