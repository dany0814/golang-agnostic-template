package business

import (
	"golang-agnostic-template/src/application/domain/utils"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(hashed, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}

func HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ObfuscatePhoneNumber(phone string) (string, error) {
	asterisks := strings.Repeat("*", len(phone)-3)
	lastThree := phone[len(phone)-3:]
	return asterisks + lastThree, nil
}

func IsValidPhone(phone string) (string, error) {
	if len(phone) < 3 {
		return "", utils.ErrInvalidPhone
	}
	const phoneRegex = `^\d+$`
	re := regexp.MustCompile(phoneRegex)
	if !re.MatchString(phone) {
		return "", utils.ErrInvalidPhone
	}
	return phone, nil
}
