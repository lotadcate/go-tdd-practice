package main 

import "fmt"

const Defaultprefix = "Hello, "
const Spprefix = "Hola, "
const Frprefix = "Bonjure, "

func hello(name string, language string) string {
  if name == "" {
    name = "World"
  }
  return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
  switch language {
      case "French":
        prefix = Frprefix
      case "Spanish":
        prefix = Spprefix
      default:
      prefix = Defaultprefix                        }
  return
}

func main() {
  fmt.Println(hello("World", ""))
}
