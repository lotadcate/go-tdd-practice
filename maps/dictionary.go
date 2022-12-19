package dictionary

type Dictionary map[string]string

type DictionaryErr string
const (
  ErrNotFound   = DictionaryErr("could not find the word you were looking for")
  ErrWordExists = DictionaryErr("cannot add word because it already exists")
  ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
  return string(e)
}

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

func(dict Dictionary) Update(word, definition string) error {
  _, err := dict.Search(word)

  switch err {
  case ErrNotFound:
      return ErrWordDoesNotExist
  case nil:
      dict[word] = definition
  default:
      return err
  }

  return nil
}

func (d Dictionary) Delete(word string) {
  delete(d, word)
}
