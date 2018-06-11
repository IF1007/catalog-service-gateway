package log

import (
	"fmt"
)

func LogNow(message string) {
	Log(message)
}

func Log(message string) {
	fmt.Println(message)
}
