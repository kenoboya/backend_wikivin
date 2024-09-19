package v1

import (
	"net/http"
	"wikivin/internal/model"
	"wikivin/pkg/auth"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserRoutes(router *gin.RouterGroup){
	user := router.Group("")
	{
		user.POST("/sign-up", h.signUp)
		user.POST("/sign-in", h.signIn)
		user.GET("/refresh", h.refresh)
		user.GET("/sign-out", h.signOut)
		user.GET("/verify", h.AuthMiddleware(),func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Authenticated"})
		})

		user.POST("/favorite/articles", h.addFavorite)
		user.DELETE("/favorite/articles", h.deleteFavorite)
		user.GET("/user/profile", h.getBriefInfoProfile)
		user.GET("/favorite/articles", h.getFavoriteArticles)

		// user.GET("/profile/:username", h.userProfile)
		// user.PUT("/profile", h.updateUserProfile)

		// user.POST("/favorite-article", h.addToFavoriteArticle)
		// user.DELETE("/favorite-article/:title", h.removeFromFavoriteArticle)

		// user.GET("/favorite-articles", h.getFavoriteArticles)
		// user.GET("/history-articles", h.historyArticles)
	}
}

func (h *Handler) signUp(c *gin.Context){
	var requestSignUp  model.RequestSignUp 
	if err:= c.BindJSON(&requestSignUp); err!= nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	tokens, err:= h.services.Users.SignUp(c.Request.Context(), requestSignUp)
	if err != nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.SetCookie("access_token", tokens.AccessToken, int(h.services.Users.GetAccessTokenTTL().Seconds()), "/", "localhost", false, true)
	c.SetCookie("refresh_token", tokens.RefreshToken, int(h.services.Users.GetRefreshTokenTTL().Seconds()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *Handler) signIn(c *gin.Context){
	var userSignIn  model.UserSignIn
	if err:= c.BindJSON(&userSignIn); err!= nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	tokens, err:= h.services.Users.SignIn(c.Request.Context(), userSignIn)
	if err != nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.SetCookie("access_token", tokens.AccessToken, int(h.services.Users.GetAccessTokenTTL().Seconds()), "/", "localhost", false, true)
	c.SetCookie("refresh_token", tokens.RefreshToken, int(h.services.Users.GetRefreshTokenTTL().Seconds()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "User signed in successfully"})
}

func (h *Handler) refresh(c *gin.Context){
	refreshToken, err:= c.Cookie("refresh_token")
	if err != nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	tokens, err:= h.services.Users.RefreshToken(c.Request.Context(), refreshToken)
	if err != nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.SetCookie("access_token", tokens.AccessToken, int(h.services.Users.GetAccessTokenTTL().Seconds()), "/", "localhost", false, true)
	c.SetCookie("refresh_token", tokens.RefreshToken, int(h.services.Users.GetRefreshTokenTTL().Seconds()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully"})
}


func(h *Handler) getBriefInfoProfile(c *gin.Context){
	token, err:= c.Cookie("access_token")
	if err != nil{
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	userID, err:= h.services.Users.GetUserIDFromToken(c.Request.Context(), token, auth.AccessToken)
	if err != nil{
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	briefInfoProfile, err:= h.services.Profiles.GetBriefInfoProfile(c.Request.Context(), userID)
	if err != nil{
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, briefInfoProfile)
}

func(h *Handler) signOut(c *gin.Context){
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Sign out was successfully"})
}

func (h *Handler) getFavoriteArticles(c *gin.Context){
	token, err:= c.Cookie("access_token")
	if err != nil{
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	userID, err:= h.services.Users.GetUserIDFromToken(c.Request.Context(), token, auth.AccessToken)
	if err != nil{
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	favoriteArticles, err:= h.services.Favorites.GetFavoriteArticlesByUserID(c.Request.Context(), userID)
	if err != nil{
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, favoriteArticles)
}

func (h *Handler) addFavorite(c *gin.Context){
	var articleID int
	if err:= c.BindJSON(&articleID); err!= nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err:= c.Cookie("access_token")
	if err != nil{
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	userID, err:= h.services.Users.GetUserIDFromToken(c.Request.Context(), token, auth.AccessToken)
	if err != nil{
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if err:= h.services.Favorites.AddFavorite(c.Request.Context(), userID, articleID); err != nil{
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Favorite was created successfully"})
}

func (h *Handler) deleteFavorite(c *gin.Context){
	var articleID int
	if err:= c.BindJSON(&articleID); err!= nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err:= c.Cookie("access_token")
	if err != nil{
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	userID, err:= h.services.Users.GetUserIDFromToken(c.Request.Context(), token, auth.AccessToken)
	if err != nil{
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if err:= h.services.Favorites.DeleteFavorite(c.Request.Context(), userID, articleID); err != nil{
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Favorite was deleted successfully"})
}
