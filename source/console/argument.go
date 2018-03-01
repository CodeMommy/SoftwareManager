package console

import (
	"os"
	"errors"
)

func GetArgument(index int) (string, error){
	argument := os.Args
	argumentLength := len(argument)
	if argumentLength > index {
		return argument[index], nil
	}
	return "", errors.New("argument not exist")
}
