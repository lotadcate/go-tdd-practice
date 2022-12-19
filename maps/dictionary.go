package dictionary
import (
  "errors"
  // "fmt"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
func(dict Dictionary) Search(key string) (string, error) {
  // Map can return two values. The first value is value, and the second value is a boolean value indicating whether the key was successfully detected.
  definition, ok := dict[key]
  // fmt.Printf("def: 「%q」 ok : 「%t」\n", definition, ok)
  if !ok {
    return "", ErrNotFound
  }

  return definition, nil
}
