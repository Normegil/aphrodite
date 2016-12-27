package db

import "github.com/normegil/aphrodite/model"

var users = make([]model.User, 0)

func (db DB) User(name string) (*model.User, error) {
	for _, user := range users {
		if user.Name() == name {
			return &user, nil
		}
	}
	return nil, nil
}

func (db DB) UserCreate(user model.User) error {
	users = append(users, user)
	return nil
}
