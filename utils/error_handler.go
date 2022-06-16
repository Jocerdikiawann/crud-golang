package utils

import "log"

func IfErrorHandler(err error) {
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
}
