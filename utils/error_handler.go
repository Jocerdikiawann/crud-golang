package utils

import "fmt"

func IfErrorHandler(err error) {
	if err != nil {
		fmt.Println("Error : ", err)
	}
}
