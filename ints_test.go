package relief

import (
	"fmt"
	"github.com/a2htray/relief/ints"
	"testing"
)

func TestRange(t *testing.T) {
	r := ints.Range(0, 10, 1)
	for i, v := range r {
		if i != v {
			t.Fatal(fmt.Sprintf("test range: %d != %d", i, v))
		}
	}
	fmt.Println(r)
	r = ints.Range(-1, 10, 2)
	fmt.Println(r)
}