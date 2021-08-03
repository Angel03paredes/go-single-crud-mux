package database

import (
	"fmt"
	"test-mux/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	server   = "localhost"
	port     = 53676
	user     = "SA"
	password = "root"
	dbName   = "gorm_users"
)

func Connection() *gorm.DB {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, dbName)
	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")

	}
	user := new(models.User)
	db.AutoMigrate(&user)
	return db
}
