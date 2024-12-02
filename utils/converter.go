package utils

import (
	"log"
	"strconv"
)

func StringAsInt(s string) int {
	asInt, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return asInt
}
