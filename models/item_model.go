package models

// Item structure

type Preference struct {
	UserID           int    `json:"userId" gorm:"not null"`
	NotificationType string `json:"notificationType" gorm:"not null"`
	Enabled          string `json:"enabled" gorm:"not null"`
}

var Request struct {
	UserID           int    `json:"userId" binding:"required"`
	NotificationType string `json:"notificationTypeId" binding:"required"`
	Answer           Answer `json:"answer"`
}

type Answer struct {
	Rating   int    `json:"rating"`
	Interest string `json:"interest"`
	Comment  string `json:"comment"`
}

type Survey struct {
	UserID uint       // User ID associated with the survey
	Answer JSONString `json:"answer"`
}

type JSONString string

func (j *JSONString) UnmarshalJSON(data []byte) error {
	// Here, we store the JSON object as a string without unmarshaling it
	*j = JSONString(data)
	return nil
}

//type ItemList struct {
//	ID        uint    `json:"id"`
//	Price     int     `json:"price"`
//	Name      string  `json:"name"`
//	AvgRating float64 `json:"avg_rating"`
//}
//
//type ItemFilters struct {
//	Price  int
//	Rating int
//	Search string
//}
//
//type ItemInput struct {
//	Price int    `json:"price" binding:"required"`
//	Name  string `json:"name" binding:"required"`
//}
//
//// ItemRating structure
//type ItemRating struct {
//	ID     uint    `gorm:"primaryKey"`
//	UserID uint    `gorm:"not null"`
//	ItemID uint    `gorm:"not null"`
//	Rating float64 `gorm:"not null"`
//	User   User    `gorm:"foreignKey:UserID"`
//	Item   Item    `gorm:"foreignKey:ItemID"`
//}
//
//type ItemRatingInput struct {
//	Rating float64 `json:"rating" binding:"required"`
//}
//
//type ItemRatingResponse struct {
//	ID     uint    `gorm:"primaryKey"`
//	UserID uint    `gorm:"not null"`
//	ItemID uint    `gorm:"not null"`
//	Rating float64 `json:"rating" binding:"required"`
//}
