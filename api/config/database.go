// Configuration of the database, this case MySQL database and the ORM gorm.
package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)


var DB *gorm.DB

// struct to define DB parameters
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}


// Set configuretion to DB
func GetConfig() *DBConfig {
	dbConfig := &DBConfig{
			Host:     os.Getenv("HOST"),
			Port:     os.Getenv("PORT"),
			User:     os.Getenv("USER"),
			DBName:   os.Getenv("DATABASE"),
			Password: os.Getenv("PASSWORD"),
		}
	return dbConfig
}


// connect to database
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
