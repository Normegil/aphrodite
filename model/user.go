package model

import "github.com/normegil/aphrodite/security"

type User struct {
	id       ID
	name     string
	password Password
	disabled bool
}

type Password struct {
	hash string
	salt string
}

func NewPassword(passwordReel string) (Password, error) {
	salt, err := security.Salt()
	if nil != err {
		return err
	}

	var pepper string
	return Password{
		hash:,
		salt: salt,
	}
}