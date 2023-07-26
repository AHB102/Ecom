package main

import (
	// "github.com/AHB102/Ecom/controllers"
	// "github.com/AHB102/Ecom/database"

	"github.com/AHB102/Ecom/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(midddleware.Auth)

	//Product Service routes
	//router.GET("/AddProduct", AddProduct)
	router.GET("/ViewProduct", AddToCart)

	//Order Service Routes
	router

}
