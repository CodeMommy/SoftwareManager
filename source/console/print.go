package console

import (
	"fmt"
)

const line = "--------------------------------------------------"

func PrintLine() {
	fmt.Println(line)
}

func PrintError(message string) {
	printColor(message, BackgroundColorNormal, ForegroundColorRed)
}
