package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token, err := c.Cookie("access_token")
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
            c.Abort()
            return
        }

        err = h.services.Users.Verify(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
            c.Abort()
            return
        }
        c.Next()
    }
}