package utils

import (
	"strconv"
	"fmt"
	"errors"
)

func isValidPortString(input string) (string,error) {
	portNum, err := strconv.Atoi(input)
	if err != nil {
		return "",err
	}
	if portNum < 1 || portNum > 65535 {
		return "", errors.New("invalid port number")
	}
	return input, nil
}

func GetInputPort(args []string) (string,error) {
	if (len(args[1:]) > 0) {
		s := args[1]
		fmt.Println("Custom port entered.")
		_, err := isValidPortString(s)
		if err != nil {
			return "",err
		}
		return s,nil
	}
	return "8080",nil
}