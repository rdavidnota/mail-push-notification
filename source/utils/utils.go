package utils

import (
	"fmt"
	"log"
	"math/rand"
)

func FailOnErrorFatal(err error, msg string) {
	if err != nil {
		FailOnErrorNormal(err, msg)
		log.Fatalf("%s: %s", msg, err)
	}
}

func FailOnErrorNormal(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
		log.Printf("%s: %s", msg, err)
	}
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}
