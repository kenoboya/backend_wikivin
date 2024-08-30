package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initArticlesRoutes(router *gin.RouterGroup){
	article:= router.Group("/articles")
	{
		article.GET("/:id", h.LoadArticle)
		article.GET("", h.LoadArticlesBriefInfo)
	}
}

func (h *Handler) LoadArticle(c *gin.Context){
	articlePage, err:= h.services.Articles.LoadArticle(c.Request.Context(), 12)
	if err != nil{
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, articlePage)
}
func (h *Handler) LoadArticlesBriefInfo(c *gin.Context){
	articleBriefInfo, err:= h.services.Articles.LoadArticlesBriefInfo(c.Request.Context())
	if err != nil{
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, articleBriefInfo)
}