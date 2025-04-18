package greetings

import (
  "fmt"
  "errors"
  "math/rand"
)

func Hello(name string) (string, error) {
  // Guarding
  if(name == "") {
    return "", errors.New("empty name")
  }

  message := fmt.Sprintf(randomFormat(), name)
  return message, nil
}

func randomFormat() string {
  formats := []string {
    "Hi, %v, Welcome!",
    "Great to see you %v!",
    "By the power of 42! - welcome %v!",
    "May Go be with you %v!",
  }

  return  formats[rand.Intn(len(formats))]
}
