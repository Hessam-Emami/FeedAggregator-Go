package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAuthHeader(request *http.Request) (string, error) {
	authHeader := request.Header.Get("Authorization")
	if len(authHeader) == 0 {
		return "", errors.New("missing authorization header")
	}

	splitted := strings.Split(authHeader, " ")
	if len(splitted) != 2 || splitted[0] != "ApiKey" {
		return "", errors.New("missing authorization header")
	}

	return splitted[1], nil
}
