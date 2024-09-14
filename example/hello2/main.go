//go:build aspect

package main

import (
	"fmt"
)

func sayHello(s string) {
	fmt.Println("hello " + s)
}
func sayHello2(s string) {
	fmt.Println("hello " + s)
}
func sayHello3(s string) {
	fmt.Println("hello " + s)
}

func main() {
	sayHello("world")
	sayHello("jim")
	sayHello("tom")
}
