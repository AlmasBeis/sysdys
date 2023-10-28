package routes

import (
	"FixPrice/controllers"
	//"FixPrice/middleware"
	"github.com/gin-gonic/gin"
)

type PopUpRouteController struct {
	popUpController controllers.PopUpController
}

func NewItemRouteController(popUpController controllers.PopUpController) PopUpRouteController {
	return PopUpRouteController{popUpController}
}

func (rc *PopUpRouteController) ItemRoute(rg *gin.RouterGroup) {
	router := rg.Group("/popup")
	router.POST("/send", rc.popUpController.SendNotification)
	router.PUT("/pref", rc.popUpController.StoreUserPreferences)
	router.POST("/survey", rc.popUpController.CreateSurvey)
	//router.POST("/rating/:id", middleware.DeserializeUser(), rc.popUpController.GiveRatingToItem)
	//router.POST("", middleware.DeserializeUser(), rc.popUpController.CreateItem)
}
