package controllers

import (
	"cakrawala-go/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MarketController struct {
	DB *gorm.DB
}

func NewMarketController(DB *gorm.DB) MarketController {
	return MarketController{DB}
}

func (mc *MarketController) FindMarket(ctx *gin.Context) {
	var q = ctx.Param("q")
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var markets []models.Market
	results := mc.DB.Limit(intLimit).Offset(offset)

	if q == nil {
		results.Where("name like ?", "%"+q+"%")
	}

	results.Find(&markets)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(markets), "data": markets})
}
