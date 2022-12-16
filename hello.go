package main 

import "fmt"

const Defaultprefix = "Hello, "
const Spprefix = "Hola, "
const Frprefix = "Bonjure, "

func hello(name string, language string) string {
  if name == "" {
    name = "World"
  }
  prefix := Defaultprefix

  switch language {
    case "French":
      prefix = Frprefix
    case "Spanish":
      prefix = Spprefix
  }
  return prefix + name
}

func main() {
  fmt.Println(hello("World", ""))
}
