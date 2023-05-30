package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// 1. membuat endpoint menggunakan Gin yang mengembalikan random quote anime https://animechan.vercel.app/api/random/anime?title=naruto

type Quote struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

// 1a. Membuat HTTP Server
func main() {
	// Membuat HTTP Server
	r := gin.Default()
	r.GET("/quote", func(c *gin.Context) {
		quote, err := getAnimeQuote()
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
	})
	r.Run(":8081")
}

// 1a1. Membuat design endpoint
//(API contract:
// - HTTP method GET
// - Path /quote
// - response
// {
// 	anime: "Naruto",
// 	character: "...",
// 	quote: "..."
// } )

// 1b. Membuat HTTP Client
func getAnimeQuote() (*Quote, error) {
	resp, err := http.Get("https://animechan.vercel.app/api/random/anime?title=naruto")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch anime quote")
	}

	body, err := ioutil.ReadAll(resp.Body)

	var quote Quote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil, err
	}

	return &quote, nil
}
