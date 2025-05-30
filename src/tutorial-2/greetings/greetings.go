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

  // message := fmt.Sprintf(randomFormat(), name)

  message := fmt.Sprint(randomFormat())

  return message, nil
}

func Hellos(names []string) (map[string]string, error) {
  messages := make(map[string]string)

  for _, name := range names {
    message, err := Hello(name)

    if err != nil {
      return nil, err
    }

    messages[name] = message
  }

  return messages, nil
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
