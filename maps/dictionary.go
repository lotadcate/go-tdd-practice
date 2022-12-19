package dictionary

type Dictionary map[string]string

func(dict Dictionary) Search(key string) string {
  return dict[key]
}
