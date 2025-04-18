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

  names := []string{"Robert Griesemer", " Rob Pike", "Ken Thompson"}

  messages, err := greetings.Hellos(names)
  if(err != nil) {
    // Fatal causes process exit!
    log.Fatal(err)
  }

  fmt.Println(messages)
}
