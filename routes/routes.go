package routes

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

var MaxReqSize int64 = 32000000

func init() {
	env_mrs := os.Getenv("APP_MAX_REQ_SIZE")

	if env_mrs != "" {
		mrs, err := strconv.ParseInt(env_mrs, 10, 64)
		if err != nil && mrs > 0 {
			MaxReqSize = mrs
		}
	}
}

func formSayi(deger string) (int, error) {
	var sayi int

	str, err := formStr(deger)
	if err != nil {
		return 0, err
	}

	sayi, err = strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return sayi, nil
}

func formStr(deger string) (string, error) {
	str := strings.TrimSpace(deger)

	if len(str) < 1 {
		return "", errors.New("Eksik değer!")
	}

	return str, nil
}
