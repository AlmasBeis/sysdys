package main

import (
	"FixPrice/controllers"
	"FixPrice/initializers"
	"FixPrice/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	ItemController      controllers.PopUpController
	ItemRouteController routes.PopUpRouteController

	//redisClient *redis.Client
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	//redisClient = redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // No password by default
	//	DB:       0,  // Default DB
	//})

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	ItemController = controllers.NewPopUpController(initializers.DB)
	ItemRouteController = routes.NewItemRouteController(ItemController)
	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	router := server.Group("/api")
	router.GET("/health-checker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	ItemRouteController.ItemRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
