package datasource

import "github.com/normegil/aphrodite/model"

type DataSource interface {
	ImageLoader
	UserLoader
	UserCreator
}

type ImageLoader interface {
	AllImages(offset, limit int) []model.Image
	Image(id model.ID) model.Image
}

type UserLoader interface {
	//AllUsers(offset, limit int) []User
	SpecificUserLoader
}

type SpecificUserLoader interface {
	User(name string) (*model.User, error)
}

type UserCreator interface {
	UserCreate(model.User) error
}
