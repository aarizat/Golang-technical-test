// Define all of functions incharge of hadlering http methods.
package controllers

import (
	"api/config"
	"api/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateEmployee creates a new employee and maps to database.
func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	var employee models.Employee

	decoder :=  json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := config.DB.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, employee)
}


// GetEmployees fecths all of employees from database.
func GetEmployees(w http.ResponseWriter, r *http.Request) {

	var employees []models.Employee

	config.DB.Find(&employees)
	respondJSON(w, http.StatusOK, employees)
}


// GetEmployee fecths an employee with id_card specified.
func GetEmployee(w http.ResponseWriter, r *http.Request) {

	var employee models.Employee

	id := mux.Vars(r)["id_card"]
	if err := config.DB.First(&employee, "id_card = ?", id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}


// DeleteEmployee deletes an employee specified by his ID.
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	var employee models.Employee

	id := mux.Vars(r)["id_card"]
	if err := config.DB.Where("Id_Card = ?", id).Delete(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, nil)
}


// UpdateEmployee updates a employee given his id_card.
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {

	var employee models.Employee
	id := mux.Vars(r)["id_card"]
	if err := config.DB.First(&employee, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	if err := config.DB.Model(&employee).Update(employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}
