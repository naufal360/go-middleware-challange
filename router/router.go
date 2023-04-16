package router

import (
	"go-middleware-challange/controllers"
	"go-middleware-challange/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/", controllers.GetAllProduct)
		productRouter.GET("/:productId", middlewares.ProductAuthorization(), controllers.GetProductById)
		productRouter.PUT("/:productId", middlewares.ProductUpdateDeleteAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:productId", middlewares.ProductUpdateDeleteAuthorization(), controllers.DeleteProductById)
	}

	return r
}
