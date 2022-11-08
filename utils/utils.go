package utils

import (
	"errors"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var requestBodyMap map[string]interface{}

func GetRequestBodyMap(c *gin.Context) ([]byte, error) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error
		return nil, errors.New("invalid request body")

	}

	return jsonData, nil
}
