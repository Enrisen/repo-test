// create a simple hellow world program in go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(greeting())
}

func greeting() string {
	return "Hellow World"
}
