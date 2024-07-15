package services

import (
	"dev-diary/db"
	"dev-diary/models"
)

func GetAllTags() ([]models.Tag, int64, error) {
	var tags []models.Tag
	err := db.DB.Find(&tags).Error
	if err != nil {
		return nil, 0, err
	}

	var total int64
	err = db.DB.Model(models.Tag{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return tags, total, nil
}

func GetTagByID(id uint) (models.Tag, error) {
	var tag models.Tag
	err := db.DB.Where("id = ?", id).First(&tag).Error
	return tag, err
}

func GetTagByNameWithPosts(name string) (models.Tag, error) {
	var tag models.Tag
	err := db.DB.Where("text = ?", name).Preload("Posts.Tags").First(&tag).Error
	return tag, err
}

type CreateTagData struct {
	Text            string `json:"text"`
	Color           string `json:"color"`
	BackgroundColor string `json:"background_color"`
}

func CreateTag(data CreateTagData) (models.Tag, error) {
	tag := models.Tag{
		Text:            data.Text,
		Color:           data.Color,
		BackgroundColor: data.BackgroundColor,
	}

	err := db.DB.Create(&tag).Error
	return tag, err
}

func DeleteTag(id uint) error {
	tag := models.Tag{ID: id}
	return db.DB.Delete(&tag).Error
}
