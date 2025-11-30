package handlers

import (
	"net/http"

	"github.com/faizan1191/url-shortner/storage"
	"github.com/faizan1191/url-shortner/utils"
	"github.com/gin-gonic/gin"
)

// URLHandler stores the storage instance
type URLHandler struct {
	Store *storage.MemoryStore
}

// Constructor for URLHandler
func NewURLHandler(store *storage.MemoryStore) *URLHandler {
	return &URLHandler{
		Store: store,
	}
}

// Request payload for shortening URL and url be the json key as required binding
type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

// Shorten handler: POST /shorten
func (h *URLHandler) Shorten(c *gin.Context) {
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
func (h *URLHandler) Redirect(c *gin.Context) {
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
