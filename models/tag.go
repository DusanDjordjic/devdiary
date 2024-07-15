package models

type Tag struct {
	ID              uint
	Text            string `gorm:"uniqueIndex"`
	Color           string
	BackgroundColor string
	Posts           []Post `gorm:"many2many:post_tags"`
}

type TagDTO struct {
	ID              uint   `json:"id"`
	Text            string `json:"text"`
	Color           string `json:"color"`
	BackgroundColor string `json:"background_color"`
}

func (t Tag) ToDTO() TagDTO {
	return TagDTO{
		ID:              t.ID,
		Text:            t.Text,
		Color:           t.Color,
		BackgroundColor: t.BackgroundColor,
	}
}

func DTOTags(tags []Tag) []TagDTO {
	dtos := make([]TagDTO, 0)
	for _, p := range tags {
		dtos = append(dtos, p.ToDTO())
	}
	return dtos
}
