package routes

import (
	"cakrawala-go/controllers"

	"github.com/gin-gonic/gin"
)

type MarketRouteController struct {
	marketController controllers.MarketController
}

func NewRouteMarketController(marketController controllers.MarketController) MarketRouteController {
	return MarketRouteController{marketController}
}

func (mc *MarketRouteController) MarketRoute(rg *gin.RouterGroup) {

	router := rg.Group("markets")
	router.GET("/", mc.marketController.FindMarket)
	// router.POST("/", pc.postController.CreatePost)
	// router.PUT("/:postId", pc.postController.UpdatePost)
	// router.GET("/:postId", pc.postController.FindPostById)
	// router.DELETE("/:postId", pc.postController.DeletePost)
}
