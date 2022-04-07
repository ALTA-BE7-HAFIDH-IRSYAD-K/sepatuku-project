package driver

import (
	"fmt"
	"sepatuku-project/configs"
	"sepatuku-project/entity/cart"
	"sepatuku-project/entity/order"
	"sepatuku-project/entity/product"
	"sepatuku-project/entity/user"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(configs *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		configs.Database.Username,
		configs.Database.Password,
		configs.Database.Address,
		configs.Database.Port,
		configs.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	InitialMigration(db)
	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&cart.Cart{})
	db.AutoMigrate(&order.Order{})
	db.AutoMigrate(&order.HistoryOrder{})
}
