package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	pbconf2 "github.com/mozyy/protoc-gen-gorm/example/base"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := os.Getenv("postgresql_dsn")
	db, err := gorm.Open(postgres.Open(dsn+"test22"), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		}),
	})
	if err != nil {
		panic("failed to connect database1")
	}
	db.AutoMigrate(pbconf2.ConfigGORM{})
	db.AutoMigrate(pbconf2.Config2GORM{})
	db.Create(&pbconf2.ConfigGORM{Name: "name: " + strconv.Itoa(rand.Int()), Config: &pbconf2.Config2GORM{Type: "type223"}})
	db.Unscoped().Delete(&pbconf2.ConfigGORM{ID: 2})
	c := pbconf2.ConfigGORM{ID: 2}
	e := db.Unscoped().First(&c)
	if e.Error != nil {
		fmt.Println("err", e.Error)
	} else {
		fmt.Println(c)
	}
	// db.AutoMigrate(pbconf2.CreditCardGORM{})
	// db.AutoMigrate(pbconf2.UserGORM{})
	// db.Create(&pbconf2.UserGORM{Name: "name22", CreditCard: []*pbconf2.CreditCardGORM{{Number: "number1"}, {Number: "number2"}}})
}
