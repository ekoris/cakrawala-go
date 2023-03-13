package main

import (
	"cakrawala-go/controllers"
	"cakrawala-go/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	server *gin.Engine

	MarketController      controllers.MarketController
	MarketRouteController routes.MarketRouteController
)

func init() {
	dsn := "root:@tcp(127.0.0.1:3306)/cakrawala?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success!")

	MarketController = controllers.NewMarketController(db)
	MarketRouteController = routes.NewRouteMarketController(MarketController)

	server = gin.Default()
}

func main() {
	router := server.Group("/api")
	MarketRouteController.MarketRoute(router)
	server.Run()

}
