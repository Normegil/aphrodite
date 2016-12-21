package db

import "github.com/normegil/aphrodite/model"

func (db *DB) Image(id model.ID) model.Image {
	return model.Image{}.WithID(id).WithName("Test")
}