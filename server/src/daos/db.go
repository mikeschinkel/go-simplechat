package daos

import (
	"fmt"
	"simple-chat-app/server/src/models"
	envUtil "simple-chat-app/server/src/util/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dnsStr = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
)

var (
	conn *gorm.DB = nil
)

/**
https://github.com/go-gorm/postgres
*/
func InitConn() {
	// Don't setup if already connected
	if conn != nil {
		return
	}
	// Setup connection string
	host, user, pwd, name, port := envUtil.GetDbVals()
	dsn := fmt.Sprintf(dnsStr, host, user, pwd, name, port)
	// Open connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Migrate GORM models
	db.AutoMigrate(&models.User{}, &models.UserCreds{})
	// Init connection
	conn = db
}

/**
Get the database connection.
*/
func GetDbConn() *gorm.DB {
	return conn
}
