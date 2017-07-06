package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("I'm", runtime.GOOS, "and my arch is:", runtime.GOARCH)
}
