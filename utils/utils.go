package utils

import (
	"strings"
)

func GetKeyValue(str string) (string, string) {
	keyValue := strings.Split(str, "=")
	return keyValue[0], keyValue[1]
}

func IsMatchFormat(str string) bool {
	keyValue := strings.Split(str, "\n")
	formatted := strings.ToUpper(strings.ReplaceAll(keyValue[0], " ", ""))
	format := "FORMATLAPORANSERVIS"
	return formatted == format
}
