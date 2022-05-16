package env

import (
	"os"
	"strconv"
)

func GetIntVal(key string) (res int, err error) {
	iv := os.Getenv(key)
	res, err = strconv.Atoi(iv)
	return
}

func GetStringVal(key string) (res string) {
	res = os.Getenv(key)
	return
}
