package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []book{
	{
		ID:     "1",
		Name:   "Harry Potter",
		Author: "J.K. Rowling",
		Price:  15.9,
	},
	{
		ID:     "2",
		Name:   "One Piece",
		Author: "Oda Eiichiro",
		Price:  2.99,
	},
	{
		ID:     "3",
		Name:   "Demon Slayer",
		Author: "Koyoharu Gotouge",
		Price:  2.99,
	},
}

// GET /books
func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

// GET /books/:id
func getBookById(c *gin.Context) {
	id := c.Param("id")

	for i := 0; i < len(books); i++ {
		if books[i].ID == id {
			c.JSON(http.StatusOK, books[i])
			return
		}
	}

	c.JSON(http.StatusOK, "Not Found")
}

// POST /books
func postBook(c *gin.Context) {
	var newBook book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Found"})
		return
	}

	books = append(books, newBook)
	c.JSON(http.StatusOK, newBook)
}

// PUT /books/:id
func updateBook(c *gin.Context) {
	id := c.Param("id")
	var newBook book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for i := 0; i < len(books); i++ {
		if books[i].ID == id {
			books[i].ID = newBook.ID
			books[i].Name = newBook.Name
			books[i].Author = newBook.Author
			books[i].Price = newBook.Price

			c.JSON(http.StatusOK, newBook)
			break
		}
	}
}

// DELETE /books/:id
func deleteBookById(c *gin.Context) {
	id := c.Param("id")

	for i := 0; i < len(books); i++ {
		if books[i].ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, "Delete Success")
			return
		}
	}

	c.JSON(http.StatusNotFound, "Data Not Found")
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", postBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBookById)

	router.Run(":4000")
}
