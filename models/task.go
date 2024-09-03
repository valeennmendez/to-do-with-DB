package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	ID   uint   `json: "id"`
	Task string `json:"task" validate:"requiered"`
}
