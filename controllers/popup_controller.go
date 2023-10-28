package controllers

import (
	"FixPrice/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
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

	if preference.NotificationType == "popUp" {
		// Update the user's PopUpActive field to true
		userID := preference.UserID // Assuming UserID is set in the preference
		if err := ic.DB.Model(&models.User{}).Where("id = ?", userID).Update("PopUpActive", preference.Enabled).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update PopUpActive"})
			return
		}
	} else if preference.NotificationType == "survey" {
		userID := preference.UserID // Assuming UserID is set in the preference
		if err := ic.DB.Model(&models.User{}).Where("id = ?", userID).Update("SurveyActive", preference.Enabled).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update SurveyActive"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Preference stored successfully"})
}

func (ic *PopUpController) SendNotification(c *gin.Context) {
	// Parse the request data
	request := models.Request
	var ugroup models.UserGroup
	var group models.Group

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the user's preference from the PostgreSQL database
	userID := request.UserID
	notificationType := request.NotificationType
	answer := request.Answer

	var user models.User
	if err := ic.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	fmt.Println(user)

	if notificationType == "survey" {
		if user.SurveyActive == "true" {
			ugroup.UserID = user.ID
			if err := ic.DB.Where("name = ?", answer.Interest).First(&group).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			ugroup.GroupID = group.ID
			if err := ic.DB.Create(&ugroup).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the survey"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Survey is filled! "})
		}
	} else if notificationType == "popUp" {

		if user.PopUpActive == "true" {
			if err := ic.DB.Where("user_id = ?", user.ID).First(&ugroup).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			if err := ic.DB.Where("id = ?", ugroup.GroupID).First(&group).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			if group.Name == "Computers" {
				htmlFilePath := "htmpTepmlates/Advertisment.html" // Update the path to your HTML file
				htmlContent, err := os.ReadFile(htmlFilePath)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read HTML file"})
					return
				}
				c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Notification not sent as user preference is 'false'"})
		}
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
