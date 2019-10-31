package newmath

import (
	"testing"
	"fmt"
)

func TestSqrt(t *testing.T) {
    const in, out = 4, 2
    if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
		// Output: TestSqrt
    }
}

func TestT1(t *testing.T) {
	const in, out = 4, 104
    if x := T1(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
		// Output: TestT1:104
    }
}

func ExampleT1() {
    fmt.Println("hello")
    //Output: hello1
}
