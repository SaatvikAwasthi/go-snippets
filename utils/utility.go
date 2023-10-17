package utils

import (
	"errors"
	"log"
	"runtime/debug"
)

func RecoverPanic() {
	if r := recover(); r != nil {
		log.Printf("Recovered in f %v", r)
		var err error
		switch x := r.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = errors.New("unknown panic")
		}
		if err != nil {
			stackTrace := string(debug.Stack())
			log.Printf("wru service down %v and stacktrace from panic: \n [%s]", err, stackTrace)
		}
	}
}
