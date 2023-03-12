package services

import (
	"github.com/Farheen-cell/ecom/infrastructure"
	"github.com/Farheen-cell/ecom/models"
)

func FetchAllTags() ([]models.Tag, error) {
	database := infrastructure.GetDb()
	var tags []models.Tag
	err := database.Preload("Images", "tag_id IS NOT NULL").Find(&tags).Error
	return tags, err
}
