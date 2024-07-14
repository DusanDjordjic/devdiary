package models

import "time"

type Post struct {
	ID              uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ImageURL        string
	ImageCaption    string
	ImageCaptionURL string
	Title           string
	Description     string
	Content         string
	Published       bool
	Tags            []Tag `gorm:"many2many:post_tags"`
}

type PostDTO struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	ImageURL        string    `json:"image_url"`
	ImageCaption    string    `json:"image_caption"`
	ImageCaptionURL string    `json:"image_caption_url"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Content         string    `json:"content"`
	Published       bool      `json:"published"`
	Tags            []TagDTO  `json:"tags"`
}

func DTOPosts(posts []Post) []PostDTO {
	var dtos []PostDTO
	for _, p := range posts {
		dtos = append(dtos, p.ToDTO())
	}
	return dtos
}

func (p Post) ToDTO() PostDTO {
	return PostDTO{
		ID:              p.ID,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
		ImageURL:        p.ImageURL,
		ImageCaption:    p.ImageCaption,
		ImageCaptionURL: p.ImageCaptionURL,
		Title:           p.Title,
		Description:     p.Description,
		Content:         p.Content,
		Published:       p.Published,
		Tags:            DTOTags(p.Tags),
	}
}
