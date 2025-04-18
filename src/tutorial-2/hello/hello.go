package main

import (
  "fmt"

  "example.com/greetings"
)

func main() {
  message := greetings.Hello("Paul")
  fmt.Println(message)
}
