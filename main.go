package main

import (
	"example.com/m/handler"
	"github.com/gin-gonic/gin"
)

// 1. membuat endpoint menggunakan Gin yang mengembalikan random quote anime https://animechan.vercel.app/api/random/anime?title=naruto

// 1a. Membuat HTTP Server
func main() {
	// Membuat HTTP Server
	r := gin.Default()
	r.GET("/quote", handler.GetQuote)
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
