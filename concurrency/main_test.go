package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "test://test.com" {
		return false
	}

	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{"https://baidu.com", "https://google.com", "test://test.com"}

	want := map[string]bool{
		"https://baidu.com":  true,
		"https://google.com": true,
		"test://test.com":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Microsecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		b.ResetTimer()
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
