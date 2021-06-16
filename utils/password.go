package utils

import "golang.org/x/crypto/bcrypt"

// Password ...
type Password struct {
	Value string
}

// Hash ...
func (p *Password) Hash() string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(p.Value), 14)
	return string(bytes)
}

// IsEqualsTo ...
func (p *Password) IsEqualsTo(hashedPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(p.Value))
	return err == nil
}
