package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string `gorm:"not null" json:"firstname"`
	Lastname  string `gorm:"not null" json:"lastname"`
	Email     string `gorm:"not null;unique_index" json:"email"`
	Tasks     []Task `json:"tasks"`
}
