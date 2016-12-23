package security

import "crypto/rand"

const SALT_SIZE = 32

func Salt() ([]byte, error) {
	salt := make([]byte, SALT_SIZE)
	_, err := rand.Read(salt)
	if nil != err {
		return nil, err
	}
	return salt, nil
}