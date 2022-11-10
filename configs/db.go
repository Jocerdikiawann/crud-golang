package configs

import (
	"belajar-golang-rest-api/models/roles"
	"belajar-golang-rest-api/models/user"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection(host, port, dbname, username, password string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, dbname, port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	db.AutoMigrate(&user.User{}, &roles.Roles{})
	return db, err
}
