package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"wikivin/internal/model"
	"wikivin/pkg/auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

func (h *Handler) initArticlesRoutes(router *gin.RouterGroup){
	article:= router.Group("/articles")
	{
		article.POST("/create",h.AuthMiddleware(), h.CreateArticle)
		article.GET("/:title", h.LoadArticle)
		article.GET("", h.LoadArticlesBriefInfo)
	}
}

func (h *Handler) CreateArticle(c *gin.Context) {
    token, err := c.Cookie("access_token")
    if err != nil {
        newResponse(c, http.StatusUnauthorized, err.Error())
        return
    }

    authorID, err := h.services.Users.GetUserIDFromToken(c.Request.Context(), token, auth.AccessToken)
    if err != nil {
        newResponse(c, http.StatusUnauthorized, err.Error())
        return
    }

    file, err := c.FormFile("image")
    if err != nil {
        newResponse(c, http.StatusBadRequest, model.ErrNotImageUploaded.Error())
        return
    }

    photoURL := fmt.Sprintf("%d_%s", rand.Int63(), file.Filename)
	uploadPath := filepath.Join("/root/uploads/images/articles", photoURL)
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		newResponse(c, http.StatusInternalServerError, model.ErrSaveFile.Error())
		return
	}
	photoURL = "/static/articles/" + photoURL
	

    form, err := c.MultipartForm()
    if err != nil {
        newResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    jsonData := form.Value["data"]
    if len(jsonData) == 0 {
        newResponse(c, http.StatusBadRequest, model.ErrJSONData.Error())
        return
    }

    var raw map[string]interface{}
    if err := json.Unmarshal([]byte(jsonData[0]), &raw); err != nil {
        newResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    infoBoxType, ok := raw["InfoBoxType"].(string)
    if !ok {
        newResponse(c, http.StatusBadRequest, model.ErrInfoBoxType.Error())
        return
    }
    factory, err := model.GetInfoBoxFactory(infoBoxType)
    if err != nil {
        newResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    infoBox := factory()
    infoBoxData, err := json.Marshal(raw[infoBoxType])
    if err != nil {
        newResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    if err := json.Unmarshal(infoBoxData, infoBox); err != nil {
        newResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    infoBoxDB := model.InfoBoxDB{
        InfoBoxType: infoBoxType,
        InfoBox:     infoBox,
    }

    var article model.Article
    articleData, ok := raw["Article"].(map[string]interface{})
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

    article.AuthorID = authorID
    article.Image = photoURL

    var chapters []model.Chapter
    chaptersData, ok := raw["Chapters"].([]interface{})
    if !ok {
        newResponse(c, http.StatusBadRequest, model.ErrChapterData.Error())
        return
    }
    chaptersBytes, err := json.Marshal(chaptersData)
    if err != nil {
        newResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    if err := json.Unmarshal(chaptersBytes, &chapters); err != nil {
        newResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := h.services.Articles.CreateArticle(c.Request.Context(), infoBoxDB, article, chapters); err != nil {
        newResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusOK, "Article created successfully")
}

func (h *Handler) LoadArticle(c *gin.Context){
	title:= c.Param("title")
	title = strings.Replace(title, "_", " ", -1)
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