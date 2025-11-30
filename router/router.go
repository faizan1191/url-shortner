package router

import (
	"github.com/faizan1191/url-shortener/handlers"
	"github.com/faizan1191/url-shortener/storage"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() // Create Gin router

	// Initialize memory store
	store := storage.NewMemoryStore()

	// Initialize URL handler
	urlHandler := handlers.NewURLHandler(store)

	// Register routes
	r.POST("/shorten", urlHandler.Shorten)
	r.GET("/:code", urlHandler.Redirect)

	return r
}
