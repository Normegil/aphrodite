package model

import (
	"crypto/rand"

	"golang.org/x/crypto/scrypt"
	"encoding/json"
)

type user struct {
	Name     string `json:"name"`
	Password Password `json:"password"`
	Disabled bool `json:"disabled"`
}

type User struct {
	user user
}

func NewUser(name, password string) (*User, error) {
	pass, err := NewPassword(password)
	if nil != err {
		return nil, err
	}
	return &User{user{
		Name:     name,
		Password: *pass,
		Disabled: false,
	}}, nil
}

func (u User) Name() string {
	return u.user.Name
}

func (u User) Password() Password {
	return u.user.Password
}

func (u *User) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &u.user)
}

func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.user)
}

type Password struct {
	hash            []byte
	salt            []byte

	scryptN         int
	scryptR         int
	scryptP         int
	scryptKeyLength int
}

func NewPassword(passwordReel string) (*Password, error) {
	salt, err := salt()
	if nil != err {
		return nil, err
	}

	pepper := pepper()

	const SCRYPT_N = 16384
	const SCRYPT_R = 8
	const SCRYPT_P = 1
	const KEY_LENGTH = 32
	s := append(salt, []byte(pepper)...)
	hash, err := scrypt.Key([]byte(passwordReel), s, SCRYPT_N, SCRYPT_R, SCRYPT_P, KEY_LENGTH)
	if nil != err {
		return nil, err
	}
	return &Password{
		hash: hash,
		salt: salt,

		scryptN:         SCRYPT_N,
		scryptR:         SCRYPT_R,
		scryptP:         SCRYPT_P,
		scryptKeyLength: KEY_LENGTH,
	}, nil
}

func (p Password) Check(password string) (bool, error) {
	pepper := pepper()
	s := append(p.salt, []byte(pepper)...)
	hash, err := scrypt.Key([]byte(password), s, p.scryptN, p.scryptR, p.scryptP, p.scryptKeyLength)
	if nil != err {
		return false, err
	}
	correspond := string(p.hash) == string(hash)
	return correspond, nil
}

const SALT_SIZE = 32

func salt() ([]byte, error) {
	salt := make([]byte, SALT_SIZE)
	_, err := rand.Read(salt)
	if nil != err {
		return nil, err
	}
	return salt, nil
}

func pepper() []byte {
	return []byte("")
}
