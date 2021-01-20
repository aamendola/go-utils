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

func PrintTypeAndValue(description string, obj interface{}) {
	log.Printf("--> %s %T %s", description, obj, obj)
}
