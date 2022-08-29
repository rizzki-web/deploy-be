package bcrypt

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashedByte), nil
}

func CheckPasswordHash(password, HashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(password))
	return err == nil
}
