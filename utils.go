package utils

import (
	"fmt"
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
	log.Printf("==[%s, %T, %s]\n", description, obj, obj)
}

// Color ...
type Color string

const (
	colorReset  Color = "\033[0m"
	colorRed    Color = "\033[31m"
	colorGreen  Color = "\033[32m"
	colorYellow Color = "\033[33m"
	colorBlue   Color = "\033[34m"
	colorPurple Color = "\033[35m"
	colorCyan   Color = "\033[36m"
	colorWhite  Color = "\033[37m"
)

// PrintColor ...
func PrintColor(color Color, str string) {
	fmt.Printf(string(color) + str + string(colorReset))
}
