package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("geen authenticatie info gevonden")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("foute auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("foute eerste deel van de auth header")
	}
	return vals[1], nil
}
