package controllers

import (
	"FixPrice/initializers"
	"FixPrice/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type ItemController struct {
	DB *gorm.DB
}

func NewItemController(DB *gorm.DB) ItemController {
	return ItemController{DB}
}

func (ic *ItemController) GetItems(ctx *gin.Context) {
	var items []models.ItemList

	ic.DB.Preload("ItemRatings").Find(&items)

	query := ic.DB.Table("items").
		Select("items.id, items.name, items.price, AVG(item_ratings.rating) as avg_rating").
		Joins("LEFT JOIN item_ratings ON items.id = item_ratings.item_id").
		Group("items.id").
		Order("items.id")

	ratingGTEFilter, err := strconv.Atoi(ctx.Query("rating_gte"))
	if err == nil {
		query.Where("rating >= ?", ratingGTEFilter).Find(&items)
	}
	ratingLTEFilter, err := strconv.Atoi(ctx.Query("rating_gte"))
	if err == nil {
		query.Where("rating <= ?", ratingLTEFilter).Find(&items)
	}
	priceGTEFilter, err := strconv.Atoi(ctx.Query("price_gte"))
	if err == nil {
		query.Where("price >= ?", priceGTEFilter)
	}
	priceLTEFilter, err := strconv.Atoi(ctx.Query("price_lte"))
	if err == nil {
		query.Where("price <= ?", priceLTEFilter)
	}

	searchFilter := strings.ToLower(ctx.Query("search"))
	if searchFilter != "" {
		query.Where("name ILIKE ?", "%"+searchFilter+"%")
	}
	query.Scan(&items)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "items": items})
}

func (ic *ItemController) CreateItem(ctx *gin.Context) {

	var payload *models.ItemInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	newItem := &models.Item{
		Name:  payload.Name,
		Price: payload.Price,
	}
	result := ic.DB.Create(&newItem)

	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "item": payload})
}

func (ic *ItemController) GetItemByID(ctx *gin.Context) {
	item := models.Item{}
	fmt.Println(ctx.Param("id"))
	id, _ := strconv.Atoi(ctx.Param("id"))
	initializers.DB.Find(&item, "id = ?", id)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "item": item})
}

func (ic *ItemController) GiveRatingToItem(ctx *gin.Context) {
	payload := models.ItemRatingInput{}
	id, _ := strconv.Atoi(ctx.Param("id"))
	fmt.Println(id)
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	currentUser := ctx.MustGet("currentUser").(models.User)
	newItemRating := models.ItemRating{
		ItemID: uint(id),
		UserID: currentUser.ID,
		Rating: payload.Rating,
	}
	initializers.DB.Create(&newItemRating)
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "item_rating": newItemRating})
}
