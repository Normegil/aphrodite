package model

type DataSource interface {
	ImageLoader
	UserLoader
	UserCreator
}

type ImageLoader interface {
	AllImages(offset, limit int) []Image
	Image(id ID) Image
}

type UserLoader interface {
	//AllUsers(offset, limit int) []User
	SpecificUserLoader
}

type SpecificUserLoader interface {
	User(name string) (*User, error)
}

type UserCreator interface {
	UserCreate(User) error
}
