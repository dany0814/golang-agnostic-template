package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type UIDGen interface {
	New() string
}

type uidgen struct{}

func UID() UIDGen {
	return &uidgen{}
}

func (u uidgen) New() string {
	return uuid.New().String()
}

func Parse(value string) (string, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

var ErrInvalidID = errors.New("invalid ID")

type ObjectID struct {
	value string
}

func NewObjectID(value string) (ObjectID, error) {
	v, err := Parse(value)
	if err != nil {
		return ObjectID{}, fmt.Errorf("%w: %s", ErrInvalidID, value)
	}
	return ObjectID{
		value: v,
	}, nil
}

func (id ObjectID) String() string {
	return id.value
}

func ObfuscatePhoneNumber(phone string) (string, error) {
	asterisks := strings.Repeat("*", len(phone)-3)
	lastThree := phone[len(phone)-3:]
	return asterisks + lastThree, nil
}

func IsValidPhone(phone string) (string, error) {
	if len(phone) < 3 {
		return "", ErrInvalidPhone
	}
	const phoneRegex = `^\d+$`
	re := regexp.MustCompile(phoneRegex)
	if !re.MatchString(phone) {
		return "", ErrInvalidPhone
	}
	return phone, nil
}
