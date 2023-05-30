package handler

import (
	"net/http"

	"example.com/m/repository"
	"github.com/gin-gonic/gin"
)

func GetQuote(c *gin.Context) {
	quote, err := repository.GetAnimeQuote()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Found",
		"data":    quote,
	})
}
