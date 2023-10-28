package main

import (
	"FixPrice/initializers"
	"FixPrice/models"
	"fmt"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	//if initializers.DB.Migrator().HasTable(&models.User{}) {
	//	initializers.DB.Migrator().DropTable(&models.User{})
	//}
	//if initializers.DB.Migrator().HasTable(&models.Item{}) {
	//	initializers.DB.Migrator().DropTable(&models.User{})
	//}
	//if initializers.DB.Migrator().HasTable(&models.ItemRating{}) {
	//	initializers.DB.Migrator().DropTable(&models.User{})
	//}
	initializers.DB.AutoMigrate(&models.User{}, models.Preference{}, models.Survey{})
	fmt.Println("? Migration complete")
}
