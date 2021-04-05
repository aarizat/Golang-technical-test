// Model to represent employees.
package models

import "time"

type Employee struct {
	Id_Card        string    `json:"id_card" gorm:"primary_key; size:20"`
	FirstName      string    `json:"first_name" gorm:"size:20; not null"`
	SecondName     string    `json:"second_name" gorm:"size:50"`
	LastName       string    `json:"last_name" gorm:"size:20; not null"`
	SecondLastname string    `json:"second_lastname" gorm:"size:20; not null"`
	Country        string    `json:"country" gorm:"not null; size:20"`
	IdType         string    `json:"id_type" gorm:"not null; size:20"`
	Email          string    `json:"email" gorm:"not null; unique; size:300"`
	CreatedAt      time.Time
	RegisterDate   time.Time `json:"register_date" gorm:"not null"`
	Area           string    `json:"area" gorm:"not null; size:20"`
	Status         string    `json:"status" gorm:"not null; size:6"`
}
