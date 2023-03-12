package services

import (
	"github.com/Farheen-cell/ecom/infrastructure"
)

func CreateOne(data interface{}) error {
	database := infrastructure.GetDb()
	err := database.Create(data).Error
	return err
}

func SaveOne(data interface{}) error {
	database := infrastructure.GetDb()
	err := database.Save(data).Error
	return err
}
