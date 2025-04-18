package main

import (
  "fmt"
  "log"

  "example.com/greetings"
)

func main() {
  // adjust the loggers output
  log.SetPrefix("greetings: ")
  // Controlling standard log output like date, line number, source file
  log.SetFlags(0)

  message, err := greetings.Hello("")
  if(err != nil) {
    // Fatal causes process exit!
    log.Fatal(err)
  }

  fmt.Println(message)
}
