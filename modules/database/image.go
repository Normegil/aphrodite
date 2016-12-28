package database

import (
	"github.com/normegil/aphrodite/model"
)

func (db DB) AllImages(offset, limit int) []model.Image {
	var images []model.Image
	for i := 0; i < limit; i++ {
		images = append(images, model.Image{}.WithID(model.NewID()).WithName("Test"))
	}
	return images
}

func (db DB) Image(id model.ID) model.Image {
	return model.Image{}.WithID(id).WithName("Test")
}
