package console

import (
	"fmt"
)

const BackgroundColorNormal = 0
const BackgroundColorWhite = 47

const ForegroundColorNormal = 0
const ForegroundColorRed = 31

func printColor(message string, backgroundColor int, foregroundColor int) {
	fmt.Printf("\x1b[%d;%dm", backgroundColor, foregroundColor)
	fmt.Print(message)
	fmt.Println("\x1b[0m")
}
