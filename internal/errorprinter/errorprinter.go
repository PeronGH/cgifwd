package errorprinter

import (
	"fmt"
)

func PrintErrorInCGI(err error) {
	fmt.Println("Status: 500 Internal Server Error")
	fmt.Println("Content-Type: text/plain")
	fmt.Println()
	fmt.Println(err)
}
