package services

import (
	"dev-diary/db"
	"dev-diary/models"
)

func GetAllPosts() ([]models.Post, int64, error) {
	var posts []models.Post
	err := db.DB.Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	var total int64
	err = db.DB.Model(models.Post{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

func GetPostByID(id uint) (models.Post, error) {
	var post models.Post
	err := db.DB.Where("id = ?", id).First(&post).Error
	return post, err
}

type CreatePostData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	ImageURL    string `json:"image_url"`
	Published   bool   `json:"published"`
}

func CreatePost(data CreatePostData) (models.Post, error) {
	post := models.Post{
		Title:       data.Title,
		ImageURL:    data.ImageURL,
		Content:     data.Content,
		Description: data.Description,
		Published:   data.Published,
	}

	err := db.DB.Create(&post).Error
	return post, err
}

func DeletePost(id uint) error {
	post := models.Post{ID: id}
	return db.DB.Delete(&post).Error
}
