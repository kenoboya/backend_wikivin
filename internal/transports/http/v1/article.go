package v1

import (
	"encoding/json"
	"net/http"
	"strings"
	"wikivin/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initArticlesRoutes(router *gin.RouterGroup){
	article:= router.Group("/articles")
	{
		article.POST("", h.CreateArticle)
		article.GET("/:title", h.LoadArticle)
		article.GET("", h.LoadArticlesBriefInfo)
	}
}

func (h *Handler) CreateArticle(c *gin.Context){
	var raw map[string]interface{}
	if err:= json.NewDecoder(c.Request.Body).Decode(&raw); err!= nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	infoBoxType, ok := raw["infoBoxType"].(string)
	if !ok{
		newResponse(c, http.StatusBadRequest, model.ErrInfoBoxType.Error())
		return
	}
	factory, err:= model.GetInfoBoxFactory(infoBoxType)
	if err != nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	infoBox := factory()
	infoBoxData, err:= json.Marshal(raw[infoBoxType])
	if err != nil{
        newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err:= json.Unmarshal(infoBoxData, infoBox); err!= nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	infoBoxDB:= model.InfoBoxDB{
		InfoBoxType: infoBoxType,
		InfoBox: infoBox,
	}

	var article model.Article
    articleData, ok := raw["article"].(map[string]interface{})
    if !ok {
        newResponse(c, http.StatusBadRequest, model.ErrArticleData.Error())
        return
    }
    articleBytes, err := json.Marshal(articleData)
    if err != nil {
        newResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    if err := json.Unmarshal(articleBytes, &article); err != nil {
        newResponse(c, http.StatusBadRequest, err.Error())
        return
    }

	var chapters []model.Chapter
	chaptersData, ok := raw["article"].(map[string]interface{})
	if !ok {
		newResponse(c, http.StatusBadRequest, model.ErrChapterData.Error())
		return
	}
	chaptersBytes, err:= json.Marshal(chaptersData)
	if err != nil{
		newResponse(c, http.StatusBadRequest, err.Error())
        return
	}
	if err:= json.Unmarshal(chaptersBytes, &chapters); err != nil{
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err:= h.services.Articles.CreateArticle(c.Request.Context(), infoBoxDB, article, chapters); err!= nil{
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) LoadArticle(c *gin.Context){
	title:= c.Param("title")
	title = strings.Replace(title, "_", "", -1)
	articlePage, err:= h.services.Articles.LoadArticle(c.Request.Context(), title)
	if err != nil{
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, articlePage)
}
func (h *Handler) LoadArticlesBriefInfo(c *gin.Context){
	articles, err:= h.services.Articles.LoadArticles(c.Request.Context())
	if err != nil{
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, articles)
}