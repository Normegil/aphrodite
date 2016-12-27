package model

import (
	"golang.org/x/crypto/scrypt"
	"crypto/rand"
)

type User struct {
	name     string
	password Password
	disabled bool
}

func NewUser(name, password string) (*User, error) {
	pass, err := NewPassword(password)
	if nil != err {
		return nil, err
	}
	return &User{
		name:     name,
		password: *pass,
		disabled: false,
	}, nil
}

func (u User) Name() string {
	return u.name
}

func (u User) Password() Password {
	return u.password
}

type Password struct {
	hash []byte
	salt []byte

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