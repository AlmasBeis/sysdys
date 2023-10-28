// deprecated
package samplemigrate

import (
	"FixPrice/initializers"
	"FixPrice/models"
	"log"
	"math/rand"
	"time"
)

type Item struct {
	ID     uint   `json:"id" gorm:"primarykey"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Rating int    `json:"rating"`
}

// GenerateSampleItems Generate random items with random names, prices, and ratings.
func GenerateSampleItems(count int) []*Item {
	names := []string{"Apple", "Banana", "Cherry", "Durian", "Eggplant", "Fig", "Grape", "Honeydew", "Jackfruit", "Kiwi", "Lemon", "Mango", "Nectarine", "Orange", "Papaya", "Quince", "Raspberry", "Strawberry", "Tangerine", "Ugli fruit", "Watermelon"}
	rand.Seed(time.Now().UnixNano())

	items := make([]*Item, count)
	for i := 0; i < count; i++ {
		item := &Item{
			Name:   names[rand.Intn(len(names))] + " " + names[rand.Intn(len(names))],
			Price:  rand.Intn(100) + 1,
			Rating: rand.Intn(5) + 1,
		}
		items[i] = item
	}

	return items
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	err = initializers.DB.AutoMigrate(&models.Item{})
	if err != nil {
		panic(err)
	}

	items := GenerateSampleItems(10)
	for _, item := range items {
		err = initializers.DB.Create(item).Error
		if err != nil {
			panic(err)
		}
	}
}
