package services

import (
	"dev-diary/db"
	"dev-diary/models"
)

func GetAllBlogs() ([]models.Blog, int64, error) {
	var blogs []models.Blog
	err := db.DB.Find(&blogs).Error
	if err != nil {
		return nil, 0, err
	}

	var total int64
	err = db.DB.Model(models.Blog{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return blogs, total, nil
}

func GetBlogByID(id uint) (models.Blog, error) {
	var blog models.Blog
	err := db.DB.Where("id = ?", id).First(&blog).Error
	return blog, err
}
