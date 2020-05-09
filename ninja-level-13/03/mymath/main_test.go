package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {

	type centeredAvgValues struct {
		expectedResult float64
		arguments []int
	}


	testEx := []centeredAvgValues{
		centeredAvgValues{40, []int{10, 20, 40, 60, 80}},
		centeredAvgValues{7, []int{2, 4, 6, 8, 10, 12}},
		centeredAvgValues{5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, v := range testEx {
		f := CenteredAvg(v.arguments)
		if f != v.expectedResult {
			t.Error("got", f, "want", v.arguments)
		}
	}
}




func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 4, 6, 9}))
	// Output:
	// 4
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10000})
	}
}