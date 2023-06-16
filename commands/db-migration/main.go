package main

import (
	"flag"
	"fmt"

	"github.com/bitmark-inc/config-loader"
	"github.com/bitmark-inc/logistics"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config.LoadConfig("LOGISTIC_SERVER")

	db, err := gorm.Open(postgres.Open(viper.GetString("store.dsn")), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect database, error: %s\n", err)
		flag.Usage()
	}

	if err := db.AutoMigrate(&logistics.Logistic{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&logistics.ArtworkOwnershipSnapshot{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&logistics.LogisticShipmentInformation{}); err != nil {
		panic(err)
	}
}
