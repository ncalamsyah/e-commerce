package main

import (
	"github.com/ncalamsyah/e-commerce/config"
	"github.com/ncalamsyah/e-commerce/routes"
)

// @title E-Commerce API
// @description This is a docs for e-commerce API.
// @securityDefinitions.apikey JWTAuth
// @in header
// @name Authorization
// @host localhost:9200
func main() {
	config.InitDB()
	config.InitMigrate()
	e := routes.New()
	e.Logger.Fatal(e.Start(":9200"))
}
