package seguranca

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(senha string) ([]byte, error) {
	senhahash, erro := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return senhahash, erro
}

func VerificaSenha(hash string, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
}
