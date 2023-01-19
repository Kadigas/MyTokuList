package routers

import (
	"mytokulist/controllers"
	"mytokulist/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to MyTokuList!",
		})
	})
	router.POST("/signup", controllers.Register)
	router.POST("/login", controllers.Login)

	router.POST("/logout", middleware.RequiredAuth("isLogin"), controllers.Logout)

	router.GET("/admin/home", middleware.RequiredAuth("Admin"), controllers.Validate)
	router.GET("/admin/list-admin", middleware.RequiredAuth("Admin"), controllers.GetAdmins)
	router.GET("/admin/list-user", middleware.RequiredAuth("Admin"), controllers.GetUsers)
	router.PUT("/admin/new-password", middleware.RequiredAuth("Admin"), controllers.NewPassword)

	router.GET("/home", middleware.RequiredAuth("User"), controllers.Validate)
	router.PUT("/new-password", middleware.RequiredAuth("User"), controllers.NewPassword)
	router.GET("/list-user", middleware.RequiredAuth("User"), controllers.GetUsers)

	router.GET("/categories", controllers.GetAllCategory)
	router.POST("/categories", middleware.RequiredAuth("Admin"), controllers.InsertCategory)
	router.PUT("/categories/:id", middleware.RequiredAuth("Admin"), controllers.UpdateCategory)
	router.DELETE("/categories/:id", middleware.RequiredAuth("Admin"), controllers.DeleteCategory)

	router.GET("/types", controllers.GetAllType)
	router.POST("/types", middleware.RequiredAuth("Admin"), controllers.InsertType)
	router.PUT("/types/:id", middleware.RequiredAuth("Admin"), controllers.UpdateType)
	router.DELETE("/types/:id", middleware.RequiredAuth("Admin"), controllers.DeleteType)

	router.GET("/movies", controllers.GetAllMovie)
	router.POST("/movies", middleware.RequiredAuth("Admin"), controllers.InsertMovie)
	router.PUT("/movies/:id", middleware.RequiredAuth("Admin"), controllers.UpdateMovie)
	router.PATCH("/movies/:id/attach", middleware.RequiredAuth("Admin"), controllers.AttachMovie)
	router.DELETE("/movies/:id", middleware.RequiredAuth("Admin"), controllers.DeleteMovie)

	router.GET("/watchlist", middleware.RequiredAuth("User"), controllers.GetAllWatchlist)
	router.POST("/watchlist", middleware.RequiredAuth("User"), controllers.InsertWatchlist)
	router.DELETE("/watchlist/:id", middleware.RequiredAuth("User"), controllers.DeleteWatchlist)

	return router
}
