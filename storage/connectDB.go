package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	var (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "123456"
		dbname   = "aveonlinedb"
	)

	var db *gorm.DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection error to DB %v", err)
		return nil, err
	}
	if db.Error != nil {
		log.Fatalln("Any Error in connect the DB " + err.Error())
		return nil, db.Error
	}
	log.Println("DB connected")
	return db, nil
}
