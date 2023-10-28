package controllers

import (
	"FixPrice/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type PopUpController struct {
	DB *gorm.DB
}

func NewPopUpController(DB *gorm.DB) PopUpController {
	return PopUpController{DB: DB}
}

func (ic *PopUpController) StoreUserPreferences(c *gin.Context) {
	var preference models.Preference

	if err := c.ShouldBindJSON(&preference); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create or update the user's preference in the PostgreSQL database
	db := ic.DB.Create(&preference)
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store preference"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Preference stored successfully"})
}

func (ic *PopUpController) SendNotification(c *gin.Context) {
	// Parse the request data
	request := models.Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the user's preference from the PostgreSQL database
	userID := request.UserID
	notificationType := request.NotificationType

	var preference models.Preference
	if err := ic.DB.Where("user_id = ? AND notification_type = ?", userID, notificationType).First(&preference).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve preference"})
		return
	}
	fmt.Println(preference)

	if preference.Enabled {
		if preference.NotificationType == "survey" {
			htmlBody := `<html>
			<head></head>
			<body>
				<h1>Hello User</h1>
				<p>This is a sample HTML survey notification.</p>
			</body>
			</html>`
			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlBody))
		} else if preference.NotificationType == "ad" {
			htmlBody := `<html>
			<head></head>
			<body>
				<h1>Hello User</h1>
				<p>This is a sample HTML ad notification.</p>
			</body>
			</html>`
			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlBody))
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Notification not sent as user preference is 'false'"})
	}
}

func (ic *PopUpController) CreateSurvey(c *gin.Context) {
	var survey models.Survey

	if err := c.ShouldBindJSON(&survey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(survey)
	if err := ic.DB.Create(&survey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the survey"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Survey created successfully"})
}

//func (ic *PopUpController) GetItems(ctx *gin.Context) {
//	var items []models.ItemList
//
//	query := ic.DB.Table("items").
//		Select("items.id, items.name, items.price, AVG(item_ratings.rating) as avg_rating").
//		Joins("LEFT JOIN item_ratings ON items.id = item_ratings.item_id").
//		Group("items.id").
//		Order("items.id")
//
//	ratingGTEFilter, err := strconv.Atoi(ctx.Query("rating_gte"))
//	if err == nil {
//		query.Where("rating >= ?", ratingGTEFilter).Find(&items)
//	}
//	ratingLTEFilter, err := strconv.Atoi(ctx.Query("rating_gte"))
//	if err == nil {
//		query.Where("rating <= ?", ratingLTEFilter).Find(&items)
//	}
//	priceGTEFilter, err := strconv.Atoi(ctx.Query("price_gte"))
//	if err == nil {
//		query.Where("price >= ?", priceGTEFilter)
//	}
//	priceLTEFilter, err := strconv.Atoi(ctx.Query("price_lte"))
//	if err == nil {
//		query.Where("price <= ?", priceLTEFilter)
//	}
//
//	searchFilter := strings.ToLower(ctx.Query("search"))
//	if searchFilter != "" {
//		query.Where("name ILIKE ?", "%"+searchFilter+"%")
//	}
//	query.Scan(&items)
//	ctx.JSON(http.StatusOK, gin.H{"status": "success", "items": items})
//}
//
//func (ic *PopUpController) CreateItem(ctx *gin.Context) {
//
//	var payload *models.ItemInput
//
//	if err := ctx.ShouldBindJSON(&payload); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"status": "fail", "message": err.Error(),
//		})
//		return
//	}
//	newItem := &models.Item{
//		Name:  payload.Name,
//		Price: payload.Price,
//	}
//	result := ic.DB.Create(&newItem)
//
//	if result.Error != nil {
//		ctx.JSON(http.StatusBadGateway, gin.H{
//			"status": "error", "message": result.Error.Error(),
//		})
//		return
//	}
//
//	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "item": payload})
//}
//
//func (ic *PopUpController) GetItemByID(ctx *gin.Context) {
//	item := models.ItemList{}
//	fmt.Println(ctx.Param("id"))
//	id, _ := strconv.Atoi(ctx.Param("id"))
//	initializers.DB.Table("items").
//		Select("items.id, items.name, items.price, AVG(item_ratings.rating) as avg_rating").
//		Joins("LEFT JOIN item_ratings ON items.id = item_ratings.item_id").
//		Group("items.id").
//		Order("items.id").Find(&item, "items.id = ?", id)
//	if item.ID != 0 {
//		ctx.JSON(http.StatusOK, gin.H{"status": "success", "item": item})
//	} else {
//		ctx.JSON(http.StatusOK, gin.H{"status": "success", "item": nil})
//	}
//}
//
//func (ic *PopUpController) GiveRatingToItem(ctx *gin.Context) {
//	payload := models.ItemRatingInput{}
//	id, _ := strconv.Atoi(ctx.Param("id"))
//	fmt.Println(id)
//	if err := ctx.ShouldBindJSON(&payload); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"status": "fail", "message": err.Error(),
//		})
//		return
//	}
//	currentUser := ctx.MustGet("currentUser").(models.User)
//	newItemRating := models.ItemRating{
//		ItemID: uint(id),
//		UserID: currentUser.ID,
//		Rating: payload.Rating,
//	}
//	initializers.DB.Create(&newItemRating)
//	itemResponse := models.ItemRatingResponse{
//		ID:     newItemRating.ID,
//		ItemID: newItemRating.ItemID,
//		UserID: newItemRating.UserID,
//		Rating: newItemRating.Rating,
//	}
//	ctx.JSON(http.StatusCreated, gin.H{
//		"status": "success", "item_rating": itemResponse,
//	})
//}
