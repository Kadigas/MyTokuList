package routers

import (
	"mytokulist/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/categories", controllers.GetAllCategory)
	router.POST("/categories", controllers.InsertCategory)
	router.PUT("/categories/:id", controllers.UpdateCategory)
	router.DELETE("/categories/:id", controllers.DeleteCategory)

	router.GET("/types", controllers.GetAllType)
	router.POST("/types", controllers.InsertType)
	router.PUT("/types/:id", controllers.UpdateType)
	router.DELETE("/types/:id", controllers.DeleteType)

	router.GET("/movies", controllers.GetAllMovie)
	router.POST("/movies", controllers.InsertMovie)
	router.PUT("/movies/:id", controllers.UpdateMovie)
	router.PATCH("/movies/:id/attach", controllers.AttachMovie)
	router.DELETE("/movies/:id", controllers.DeleteMovie)

	return router
}
