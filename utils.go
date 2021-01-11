package utils

import (
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

