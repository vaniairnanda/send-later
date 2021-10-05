package config

import (
	"fmt"
	"github.com/vaniairnanda/send-later/environment"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func BaseDB(dbName string) *gorm.DB {
	env := environment.Load()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", env.DBHost, env.DBUsername, env.DBPassword, dbName, env.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return db
}

func GetDBDisbursement() *gorm.DB {
	env := environment.Load()
	return BaseDB(env.DBName)
}
