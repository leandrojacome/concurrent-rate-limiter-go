package domain

import (
	"errors"
	"strings"
)

type ClientKey string

func NewClientKey(value string) (ClientKey, error) {
	normalized := strings.TrimSpace(value)
	if len(normalized) < 3 || len(normalized) > 128 {
		return "", errors.New("client key must contain 3 to 128 characters")
	}
	return ClientKey(normalized), nil
}
