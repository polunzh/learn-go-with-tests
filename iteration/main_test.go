package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got ''%s'", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 5)
	fmt.Println(repeated)
	// Output: bbbbb
}

func BenchmarkRepeat(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Repeat("a", 6)
	}
}
