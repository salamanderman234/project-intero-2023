package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDatabaseDsn() string {
	DB_USER := viper.GetString("DB_USER")
	DB_PASSWORD := viper.GetString("DB_PASSWORD")
	DB_HOST := viper.GetString("DB_HOST")
	DB_PORT := viper.GetString("DB_PORT")
	DB_NAME := viper.GetString("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
	return dsn
}

func ConnectDatabase() (*gorm.DB, error) {
	dsn := GetDatabaseDsn()

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		TranslateError: true,
	})
}