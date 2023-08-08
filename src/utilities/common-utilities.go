package utilities

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"log"
	"os/exec"
	"strings"
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
