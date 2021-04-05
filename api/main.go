// Entry point for API.
package main

import (
	"api/config"
	"api/controllers"
	"api/models"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
)

// main entry point for API.
func main() {

	var err error

	config.DB, err = gorm.Open("mysql", config.DbURL(config.GetConfig()))
	router := mux.NewRouter()
	// endpoints
	router.HandleFunc("/api/v1/employees", controllers.CreateEmployee).Methods("POST")
	router.HandleFunc("/api/v1/employees", controllers.GetEmployees).Methods("GET")
	router.HandleFunc("/api/v1/employees/{id_card}", controllers.GetEmployee).Methods("GET")
	router.HandleFunc("/api/v1/employees/{id_card}", controllers.DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/api/v1/employees/{id_card}", controllers.UpdateEmployee).Methods("PUT")

	if err != nil {
	 	fmt.Println("Status:", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Employee{})
	log.Fatal(http.ListenAndServe(":5000", router))
}
