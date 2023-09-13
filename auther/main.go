package main

import (
	"auther/controllers"
	"auther/initializer"
	"auther/initializer/migration"
	"auther/util"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnv()
	initializer.LoadDB()
	migration.MigrateDB()
}

func main() {
	serveApp()
}

func serveApp() {
	router := gin.Default()
	authRoutes := router.Group("/auth/user")
	// registration route
	authRoutes.POST("/register", controllers.Register)
	// login route
	authRoutes.POST("/login", controllers.Login)
	// logout
	authRoutes.GET("/logout", controllers.Logout)

	// Admin
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(util.JWTAuth())
	adminRoutes.GET("/users", controllers.GetUsers)
	adminRoutes.GET("/user/:id", controllers.GetUser)
	adminRoutes.PUT("/user/:id", controllers.UpdateUser)
	adminRoutes.POST("/user/role", controllers.CreateRole)
	adminRoutes.GET("/user/roles", controllers.GetRoles)
	adminRoutes.PUT("/user/role/:id", controllers.UpdateRole)

	router.Run()
}
