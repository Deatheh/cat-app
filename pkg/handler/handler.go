package handler

import (
	"github.com/Deatheh/cat-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.singUp)
		auth.POST("/sing-in", h.singIn)
	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updeteList)
			lists.DELETE("/:id", h.deleteList)
			cats := lists.Group(":id/cats")
			{
				cats.POST("/", h.createCat)
				cats.GET("/", h.getAllCat)
				cats.GET("/:cat_id", h.getCatById)
				cats.PUT("/:cat_id", h.updeteCat)
				cats.DELETE("/:cat_id", h.deleteCat)
			}
		}
	}

	return router
}
