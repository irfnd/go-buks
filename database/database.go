package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(uri string) {
	db_config := postgres.New(postgres.Config{DSN: uri})
	db, err := gorm.Open(db_config, &gorm.Config{NowFunc: func() time.Time { return time.Now().Local() }})

	if err != nil {
		fmt.Println("[DATABASE]::CONNECTION_ERROR")
		panic(err)
	}

	DB = db
	fmt.Println("[DATABASE]::CONNECTED")
}

func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
