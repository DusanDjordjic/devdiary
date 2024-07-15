package services

import (
	"dev-diary/db"
	"dev-diary/models"
	"dev-diary/utils"

	"gorm.io/gorm"
)

func SortTags(db *gorm.DB) *gorm.DB {
	return db.Order("tags.id ASC")
}

func GetAllPosts() ([]models.Post, int64, error) {
	var posts []models.Post
	err := db.DB.Order("created_at DESC").Preload("Tags", SortTags).Find(&posts).Error
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
	err := db.DB.Where("id = ?", id).Preload("Tags", SortTags).First(&post).Error
	return post, err
}

type CreatePostData struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	Content         string `json:"content"`
	ImageURL        string `json:"image_url"`
	ImageCaption    string `json:"image_caption"`
	ImageCaptionURL string `json:"image_caption_url"`
	Published       bool   `json:"published"`
	TagIDs          []uint `json:"tag_ids"`
}

func CreatePost(data CreatePostData) (models.Post, error) {
	post := models.Post{
		Title:           data.Title,
		ImageURL:        data.ImageURL,
		ImageCaption:    data.ImageCaption,
		ImageCaptionURL: data.ImageCaptionURL,
		Content:         data.Content,
		Description:     data.Description,
		Published:       data.Published,
		Tags:            make([]models.Tag, 0),
	}

	data.TagIDs = utils.Dedup(data.TagIDs)

	for _, tagID := range data.TagIDs {
		tag, err := GetTagByID(tagID)
		if err != nil {
			return post, err
		}
		post.Tags = append(post.Tags, tag)
	}

	err := db.DB.Create(&post).Error
	return post, err
}

func DeletePost(id uint) error {
	post := models.Post{ID: id}
	return db.DB.Delete(&post).Error
}
