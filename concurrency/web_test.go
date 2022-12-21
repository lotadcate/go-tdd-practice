package concurrency
import (
  "reflect"
  "testing"
)

// Satisfy the type "WebsiteChecker".
func mockWebsiteChecker(url string) bool {
  if url == "waat://furhurterwe.geds" {
    return false
  }
  return true
}

func TestCheckWebsites(t *testing.T) {
  websites := []string {
    "https://google.com",
    "https://blog.gypsydave5.com",
    "waat://furhurterwe.geds",
  }

  want := map[string]bool {
    "https://google.com" : true,
    "https://blog.gypsydave5.com" : true,
    "waat://furhurterwe.geds" : false,
  }

  got := CheckWebsites(mockWebsiteChecker, websites)

  if !reflect.DeepEqual(want, got) {
    t.Fatalf("Wanted %v, got %v", want, got)
  }
}
