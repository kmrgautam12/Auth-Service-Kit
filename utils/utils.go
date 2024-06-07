package utils

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	c.JSON(200, gin.H{"status": "healthy"})
}

func MustParseUrl(u string) *url.URL {

	url, err := url.Parse(u)
	if err != nil {
		return nil
	}
	return url
}
