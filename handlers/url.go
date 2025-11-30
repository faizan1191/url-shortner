package handlers

import (
	"net/http"

	"github.com/faizan1191/url-shortener/utils"
	"github.com/faizan1191/url-shortner/storage"
	"github.com/gin-gonic/gin"
)

// URLHanlder stores the storage instance
type URLHanlder struct {
	Store *storage.MemoryStore
}

// Constructor for URLHanlder
func NewURLHandler(store *storage.MemoryStore) *URLHanlder {
	return &URLHanlder{
		Store: store,
	}
}

// Request payload for shortening URL and url be the json key as required binding
type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

// Shorten handler: POST /shorten
func (h *URLHanlder) Shorten(c *gin.Context) {
	var req ShortenRequest

	// Bind JSON Input
	// Bind JSON input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
		return
	}

	// Generate a random short code
	code := utils.GenerateCode(6)

	// Save mapping in memory store
	h.Store.Save(code, req.URL)

	// Build Short URL
	shortURL := "http://localhost:8080/" + code

	// Return JSON Response
	c.JSON(http.StatusOK, gin.H{
		"short_url": shortURL,
	})
}

// Redirect handler Get /:code
func (h *URLHanlder) Redirect(c *gin.Context) {
	code := c.Param("code")

	// Lookup Original URL
	url, ok := h.Store.Get(code)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Redirect to Original URL
	c.Redirect(http.StatusMovedPermanently, url)
}
