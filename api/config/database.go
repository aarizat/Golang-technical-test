// Configuration of the database, this case MySQL database and the ORM gorm.
package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)


var DB *gorm.DB

// struct to define DB parameters
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}


// Set configuretion to DB
func GetConfig() *DBConfig {
	dbConfig := &DBConfig{
			Host:     "localhost",
			Port:     3306,
			User:     "root",
			DBName:   "cidenet_db",
			Password: "aarizat",
		}
	return dbConfig
}


// connect to database
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
