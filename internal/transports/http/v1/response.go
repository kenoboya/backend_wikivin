package v1

import "github.com/gin-gonic/gin"

type IDResponse struct{
	ID int
}

type response struct{
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, response{Message: message})
}
