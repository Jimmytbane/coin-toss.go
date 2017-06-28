package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Println("I am: ", runtime.GOOS)
}
