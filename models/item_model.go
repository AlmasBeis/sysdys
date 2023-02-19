package models

type Item struct {
	ID    uint   `gorm:"primary_key"`
	Price int    `gorm:"not null"`
	Name  string `gorm:"not null"`
}

type ItemList struct {
	ID        uint    `json:"id"`
	Price     int     `json:"price"`
	Name      string  `json:"name"`
	AvgRating float64 `json:"avg_rating"`
}

type ItemFilters struct {
	Price  int
	Rating int
	Search string
}

type ItemInput struct {
	Price int    `json:"price" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

type ItemRating struct {
	ID     uint    `gorm:"primaryKey"`
	UserID uint    `gorm:"not null"`
	ItemID uint    `gorm:"not null"`
	Rating float64 `gorm:"not null"`
	User   User    `gorm:"foreignKey:UserID"`
	Item   Item    `gorm:"foreignKey:ItemID"`
}

type ItemRatingInput struct {
	Rating float64 `json:"rating" binding:"required"`
}
