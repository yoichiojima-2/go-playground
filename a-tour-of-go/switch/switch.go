package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("go run on ")
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("macos")
		case "linux":
			fmt.Println("linux")
		default:
			fmt.Printf("%s. \n", os)
	}
}

