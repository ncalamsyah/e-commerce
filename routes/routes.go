package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	controllersOrder "github.com/ncalamsyah/e-commerce/controllers/order"
	controllersProduct "github.com/ncalamsyah/e-commerce/controllers/product"
	controllersUser "github.com/ncalamsyah/e-commerce/controllers/user"
	controllersWallet "github.com/ncalamsyah/e-commerce/controllers/wallet"
	_ "github.com/ncalamsyah/e-commerce/docs"
	"github.com/ncalamsyah/e-commerce/middlewares"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/docs/*", echoSwagger.WrapHandler)

	// user
	public := e.Group("/public")
	public.POST("/users/register", controllersUser.CreateUser)
	public.POST("/users/register-seller", controllersUser.CreateUserSeller)
	public.POST("/users/login", controllersUser.Login)
	public.GET("/users/:id", controllersUser.GetUser)
	public.GET("/products", controllersProduct.GetAllProduct)
	public.GET("/product/:id", controllersProduct.GetDetailProduct)
	public.GET("/products/:user_id", controllersProduct.GetProductByUserId)

	// product
	p := e.Group("/product")
	p.Use(middleware.JWT([]byte(middlewares.SECRET)))
	p.POST("", controllersProduct.CreateProduct)
	p.PUT("/:id", controllersProduct.UpdateProduct)
	p.DELETE("/:id", controllersProduct.DeleteProduct)

	// order
	o := e.Group("/order")
	o.Use(middleware.JWT([]byte(middlewares.SECRET)))
	o.POST("", controllersOrder.CreateOrder)
	o.GET("/:id", controllersOrder.GetOrderDetail)
	o.PUT("/:id", controllersOrder.UpdateOrder)
	o.GET("/order-list", controllersOrder.GetOrderByUserIdOrProductIds)

	// user
	u := e.Group("/user")
	u.Use(middleware.JWT([]byte(middlewares.SECRET)))
	u.GET("/wallet", controllersWallet.GetWalletByUserId)
	u.POST("/wallet", controllersWallet.CreateWallet)
	u.POST("/wallet-topup", controllersWallet.TopUpWallet)
	u.PUT("/wallet-confirm/:id", controllersWallet.ConfirmWalletTrans)
	return e
}
