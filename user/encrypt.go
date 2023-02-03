package user

import "golang.org/x/crypto/bcrypt"

type Encrypt interface {
	HashPassword(string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
type EncryptStruct struct {
}

func (a *EncryptStruct) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (a *EncryptStruct) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func NewEncrypt() Encrypt {
	return &EncryptStruct{}
}
