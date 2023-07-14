package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sejal/go-react-calorie-tracker-yt/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id", routes.GetEntryById)
	router.GET("/ingredients/:ingredients", routes.GetEntriesBYIngredient)

	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.PUT("/ingredients/update/:id", routes.UpdateIngredients)
	router.DELETE("'entry/delete/:id", routes.DeleteEntry)
	router.Run(":" + port)

}
