package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"pustaka-api/book"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "Khairon",
		"profesi": "Perawat Informatika",
	})
}
func AboutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":        "Khairon",
		"profesi":     "Perawat Informatika",
		"pendidikan":  "Keperawatan",
		"Konsentrasi": "Informatika Keperawatan",
		"Agama":       "Islam",
	})
}
func BookHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}
func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{"title": title})
}

func PostBooksHandler(c *gin.Context) {
	var bookInput book.BookRequest
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
}
