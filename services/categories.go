package services

import (
	"github.com/Farheen-cel/ecom/infrastructure"
	"github.com/Farheen-cell/ecom/models"
)

func FetchAllCategories() ([]models.Category, error) {
	database := infrastructure.GetDb()
	var categories []models.Category
	err := database.Preload("Images", "category_id IS NOT NULL").Find(&categories).Error
	return categories, err
}
