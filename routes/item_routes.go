package routes

import (
	"FixPrice/controllers"
	"FixPrice/middleware"
	"github.com/gin-gonic/gin"
)

type ItemRouteController struct {
	itemController controllers.ItemController
}

func NewItemRouteController(itemController controllers.ItemController) ItemRouteController {
	return ItemRouteController{itemController}
}

func (rc *ItemRouteController) ItemRoute(rg *gin.RouterGroup) {
	router := rg.Group("/items")
	router.GET("", rc.itemController.GetItems)
	router.GET("/:id", rc.itemController.GetItemByID)
	router.POST("/rating/:id", middleware.DeserializeUser(), rc.itemController.GiveRatingToItem)
	router.POST("", middleware.DeserializeUser(), rc.itemController.CreateItem)
}
