//go:build aspect

package main

import (
	"fmt"
)

func sayHello(s string) {
	fmt.Println("hello " + s)
}
func sayHello2(s string) {
	fmt.Println("hello2 " + s)
}

func main() {
	sayHello("world")
	sayHello2("world")
}
