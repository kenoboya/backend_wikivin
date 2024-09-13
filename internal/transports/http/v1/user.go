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
	
	tokens, err:= h.services.User.SignUp(c.Request.Context(), requestSignUp)
	if err != nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.SetCookie("access_token", tokens.AccessToken, int(h.services.User.GetAccessTokenTTL().Seconds()), "/", "localhost", false, true)
	c.SetCookie("refresh_token", tokens.RefreshToken, int(h.services.User.GetRefreshTokenTTL().Seconds()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *Handler) signIn(c *gin.Context){
	var userSignIn  model.UserSignIn
	if err:= c.BindJSON(&userSignIn); err!= nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	tokens, err:= h.services.User.SignIn(c.Request.Context(), userSignIn)
	if err != nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.SetCookie("access_token", tokens.AccessToken, int(h.services.User.GetAccessTokenTTL().Seconds()), "/", "localhost", false, true)
	c.SetCookie("refresh_token", tokens.RefreshToken, int(h.services.User.GetRefreshTokenTTL().Seconds()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "User signed in successfully"})
}

func (h *Handler) refresh(c *gin.Context){
	refreshToken, err:= c.Cookie("refresh_token")
	if err != nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	tokens, err:= h.services.User.RefreshToken(c.Request.Context(), refreshToken)
	if err != nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.SetCookie("access_token", tokens.AccessToken, int(h.services.User.GetAccessTokenTTL().Seconds()), "/", "localhost", false, true)
	c.SetCookie("refresh_token", tokens.RefreshToken, int(h.services.User.GetRefreshTokenTTL().Seconds()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully"})
}