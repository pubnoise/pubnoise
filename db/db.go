package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Story struct {
	ID uint `gorm:"uniqueIndex;autoIncrement"`
	// Messages []Message
	UserId uint
	// CreateAt  time.Time
	// UpdatedAt time.Time
}

type Message struct {
	gorm.Model
	ID      uint `gorm:"uniqueIndex;autoIncrement"`
	Message *string
	// CreateAt  time.Time
	// UpdatedAt time.Time
}

func Connection() {

	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DBPORT")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Story{}, &Message{})

}
