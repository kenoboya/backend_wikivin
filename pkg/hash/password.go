package hash

import (
	"crypto/sha256"
	"fmt"
	"wikivin/internal/model"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

type SHA256Hasher struct {
	salt string
}

func NewSHA256Hasher(salt string) *SHA256Hasher {
	return &SHA256Hasher{salt}
}

func (s *SHA256Hasher) Hash(password string) (string, error) {
	hash := sha256.New()
	if _, err := hash.Write([]byte(password)); err != nil{
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(s.salt))), nil
}

func (s *SHA256Hasher) Compare(hashedPassword, password string) error{
	var err error
	if password, err = s.Hash(password); err != nil{
		return err
	}
	if hashedPassword != password{
		return model.ErrInvalidPassword
	}
	return nil
}