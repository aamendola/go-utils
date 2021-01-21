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
	ColorRed    Color = "\033[31m"
	ColorGreen  Color = "\033[32m"
	ColorYellow Color = "\033[33m"
	ColorBlue   Color = "\033[34m"
	ColorPurple Color = "\033[35m"
	ColorCyan   Color = "\033[36m"
	ColorWhite  Color = "\033[37m"
)

// PrintColor ...
func PrintColor(color Color, str string) {
	fmt.Printf(string(color) + "" + str + "" + string(colorReset))
}
