package models

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"` // receptionist/doctor
}

type Patient struct {
	ID        uint      `gorm:"primaryKey"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Age       int
	Gender    string
	Address   string
	Phone     string
	Details   string // medical detail
	UpdatedBy uint   
	CreatedAt time.Time
	UpdatedAt time.Time
}


// CREATE TABLE patients (
//     id SERIAL PRIMARY KEY,
//     first_name VARCHAR(255),
//     last_name VARCHAR(255),
//     age INT,
//     gender VARCHAR(50),
//     address VARCHAR(255),
//     phone VARCHAR(50),
//     details TEXT,
//     updated_by INT,
//     created_at TIMESTAMP,
//     updated_at TIMESTAMP
// );