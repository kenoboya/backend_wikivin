package v1

import "github.com/gin-gonic/gin"

type iDResponse struct{
	ID int `json:"id"`
}

type response struct{
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, response{Message: message})
}