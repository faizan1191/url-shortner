package router

import (
	"github.com/faizan1191/url-shortner/handlers"
	"github.com/faizan1191/url-shortner/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() // Create Gin router
	r.Use(cors.Default())

	// Initialize memory store
	// store := storage.NewMemoryStore()

	// Initialize redis store
	store := storage.NewRedisStore("localhost:6379")

	// Initialize URL handler
	urlHandler := handlers.NewURLHandler(store)

	// Register routes
	r.POST("/shorten", urlHandler.Shorten)
	r.GET("/:code", urlHandler.Redirect)

	return r
}
