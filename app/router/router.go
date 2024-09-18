package router

import (
	"online-store/app/controller"
	"online-store/app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	//cors
	r.Use(cors.Default())
	v1 := r.Group("/api/v1")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	//register
	v1.POST("/register", controller.Register)

	//login
	v1.POST("/login", controller.Login)

	authRoutes := v1.Group("/auth")
	authRoutes.Use(middleware.RequireAuth)
	{
		//products
		authRoutes.GET("/products", controller.GetAllProducts)
		authRoutes.POST("/products", controller.CreateProduct)

		//category
		authRoutes.POST("/category", controller.CreateCategory)
		authRoutes.GET("/category", controller.GetAllCategories)

		//cart
		authRoutes.GET("/cart", controller.GetAllCart)
		authRoutes.POST("/cart", controller.CreateCart)
		authRoutes.DELETE("/cart/:id", controller.DeleteCartByID)
	}

	return r
}
