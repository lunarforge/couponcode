package voucher

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestExampleGenerate(t *testing.T) {
	rand.Seed(42) // to force consistent values for example
	code := Generate()
	fmt.Println(code)
	// Output: RCMD-CRVF-FK36
}

func TestExampleGenerateCustom(t *testing.T) {
	b := make(map[string]int)
	rand.Seed(1234742) // to force consistent values for example
	for i := 0; i <= 1000000; i++ {
		cc := New(3, 6)
		code := cc.Generate()
		if _, ok := b[code]; ok {
			fmt.Printf("Duplicate Code %v\n", code)
			break
		} else {
			validated, err := Validate(code)
			if err != nil {
				t.Fatalf("%v %v %v %v\n", i, code, validated, err)
			}

			b[validated] = i
		}
	}
}
