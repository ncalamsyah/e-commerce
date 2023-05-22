package config

import (
	"fmt"

	"github.com/labstack/gommon/log"
	order "github.com/ncalamsyah/e-commerce/models/order/entity"
	product "github.com/ncalamsyah/e-commerce/models/product/entity"
	users "github.com/ncalamsyah/e-commerce/models/users/entity"
	wallet "github.com/ncalamsyah/e-commerce/models/wallet/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() {
	// config := os.Getenv("DB_DSN")
	dsn := PostgresURI("alam", "root", fmt.Sprintf(`%s:%s`, "localhost", "5432"), "ecommerce")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Print("db connected")

}

func InitMigrate() {
	DB.AutoMigrate(&users.Users{})
	DB.AutoMigrate(&product.Product{})
	DB.AutoMigrate(&order.Transactions{})
	DB.AutoMigrate(&wallet.Wallet{})
	log.Print("migrate successfully")
}

func PostgresURI(dbUserName, dbPassword, dbAddress, dbName string) string {
	return fmt.Sprintf(`postgres://%s:%s@%s/%s?sslmode=disable`, dbUserName, dbPassword, dbAddress, dbName)
}
