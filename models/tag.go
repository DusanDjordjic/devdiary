package models

type Tag struct {
	ID              uint
	Text            string
	Color           string
	BackgroundColor string
	Posts           []Post `gorm:"many2many:post_tags"`
}

type TagDTO struct {
	ID              uint
	Text            string
	Color           string
	BackgroundColor string
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
	var dtos []TagDTO
	for _, p := range tags {
		dtos = append(dtos, p.ToDTO())
	}
	return dtos
}
