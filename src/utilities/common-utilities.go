package utilities

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func GenerateUUID() (string, error) {
	out, err := exec.Command("uuidgen").Output()
	uuid := fmt.Sprintf("%s", out)
	uuid = strings.TrimSpace(uuid)
	if err != nil {
		log.Fatal("Error", err)
		return "", err
	}
	return uuid, nil
}

func EnctyptText(plainText string) (string, error) {
	return "", nil
}

func Decrypt(hash string) (string, error) {
	return "", nil
}

func GenerateHash(plainText string) string {
	hash := hmac.New(sha512.New, []byte(plainText))
	bs := hash.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func CompareHash(computedHash, expectedHash []byte) bool {
	return hmac.Equal(computedHash, expectedHash)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
