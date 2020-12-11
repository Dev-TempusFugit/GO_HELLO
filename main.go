package main

import (
	"fmt"

	"./greet"
)

func main() {
	saludo := greet.English()
	fmt.Println(saludo)
}
