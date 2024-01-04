package hasher

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var enabled = true

func init() { // nolint: gochecknoinits
	disableHas := os.Getenv("UNSAFE_DISABLE_PASSWORD_PROJECTION") == "yes_i_am_sure"

	if disableHas {
		fmt.Println("WARNING: Password protection is disabled. This is unsafe in production.")
		enabled = false
	}
}

func HashPassword(password string) (string, error) {
	if !enabled {
		return password, nil
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	if !enabled {
		return password == hash
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
