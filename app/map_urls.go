package app

import (
	auth "Book-Rental-Service/JWT-Auth"
	"Book-Rental-Service/controllers"
)

func mapUrls() {
	router.POST("/token", controllers.Token)
	router.POST("/token/v2", controllers.GetTokenV2)
	router.Use(auth.Middleware())

	router.GET("/ping", controllers.Ping)
	router.GET("/user", controllers.GetUser)
	router.POST("/Signup", controllers.CreateUser)
	router.DELETE("/user", controllers.DeleteUser)
}
