package v1

import (
	"net/http"
	"wikivin/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserRoutes(router *gin.RouterGroup){
	user := router.Group("")
	{
		user.POST("/sign-up", h.signUp)
		user.POST("/sign-in", h.signIn)
		user.POST("/sign-out", h.signOut)
		user.GET("/refresh", h.refresh)

		user.GET("/profile", h.userProfile)
		user.PUT("/profile", h.updateUserProfile)

		user.POST("/favorite-article", h.addToFavoriteArticle)
		user.DELETE("/favorite-article/:title", h.removeFromFavoriteArticle)

		user.GET("/favorite-articles", h.getFavoriteArticles)
		user.GET("/history-articles", h.historyArticles)
	}
}

func (h *Handler) signUp(c *gin.Context){
	var requestSignUp  model.RequestSignUp 
	if err:= c.BindJSON(&requestSignUp); err!= nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	token, err:= h.services.User.SignUp(c.Request.Context(), requestSignUp)
	if err != nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}