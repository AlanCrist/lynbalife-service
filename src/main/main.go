package main

import (
	"fmt"

	"../router"
)

func main() {
	fmt.Println("welcome to the server")
	e := router.New()

	e.Start(":8000")
}
