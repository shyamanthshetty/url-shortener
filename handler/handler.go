package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shyamanthshetty/go-url-shortener/shortener"
	"github.com/shyamanthshetty/go-url-shortener/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "https://localhost:9808/"
	c.JSON(200, gin.H{
		"message":  "short url created successfully",
		"shortUrl": host + shortUrl,
	})

}
func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	intialUrl := store.GetUrlMapping(shortUrl)
	c.Redirect(302, intialUrl)
}
