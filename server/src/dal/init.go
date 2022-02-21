package dal

import (
	"fmt"
	"simple-chat-app/server/src/models"
	"simple-chat-app/server/src/shared"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dnsStr = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
)

var (
	db *gorm.DB = nil
)

/**
https://github.com/go-gorm/postgres
*/
func Init() {
	// Don't setup if already connected
	if db != nil {
		return
	}
	// Setup connection string
	host, user, pwd, name, port := shared.GetDbVals()
	dsn := fmt.Sprintf(dnsStr, host, user, pwd, name, port)
	// Open connection
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Migrate GORM models
	conn.AutoMigrate(&models.User{}, &models.UserCreds{})
	// Init connection
	db = conn
}
