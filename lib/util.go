package lib

import "fmt"

func Assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func Assertf(condition bool, message string, args ...interface{}) {
	if !condition {
		panic(fmt.Sprintf(message, args...))
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
