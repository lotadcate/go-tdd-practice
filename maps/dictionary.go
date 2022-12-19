package dictionary
import (
  "errors"
  // "fmt"
)

type Dictionary map[string]string

var (
  ErrNotFound   = errors.New("could not find the word you were looking for")
  ErrWordExists = errors.New("cannot add word because it already exists")
)

func(dict Dictionary) Search(key string) (string, error) {
  // Map can return two values. The first value is value, and the second value is a boolean value indicating whether the key was successfully detected.
  definition, ok := dict[key]
  // fmt.Printf("def: 「%q」 ok : 「%t」\n", definition, ok)
  if !ok {
    return "", ErrNotFound
  }

  return definition, nil
}

func(dict Dictionary) Add(word, definition string) error{
  _, err := dict.Search(word)

  switch err {
    case ErrNotFound:
      dict[word] = definition
    case nil:
      return ErrWordExists
    default:
      return err
  }    

  return nil
}