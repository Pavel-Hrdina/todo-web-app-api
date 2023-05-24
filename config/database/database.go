package database

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/Pavel-Hrdina/todo/backend/config"
)

// DB is the database connection
var DB *gorm.DB

// Initiates the connection to sqlite database and migrates all the tables
func Connect() {
	db, err := gorm.Open(sqlite.Open(config.DB), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("[DATABASE]::CONNECTION_ERROR")
		panic(err)
	}

	// start a database connection
	DB = db
	fmt.Println("[DATABASE]::CONNECTED")
}

func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
