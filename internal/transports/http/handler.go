package rest

import (
	"wikivin/internal/config"
	"wikivin/internal/service"
	v1 "wikivin/internal/transports/http/v1"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services 
}

func NewHandler(services *service.Services) *Handler{
	return &Handler{services}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine{
	router:= gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	h.initAPI(router)
	return router
}

func(h *Handler) initAPI(router *gin.Engine){
	handlerV1 := v1.NewHandler(h.services)
	app:= router.Group("/api")
	{
		handlerV1.Init(app)
	}
}