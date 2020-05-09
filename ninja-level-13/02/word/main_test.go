package word

import (
	"testing"
	"udemy-hands-on/ninja-level-13/02/quote"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++{
		Count(quote.SunAlso)
	}
}


func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++{
		UseCount(quote.SunAlso)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three three")

	for k,v := range m{
		switch k{
		case "one":
			if v != 1 {
				t.Errorf("Expected 1")
			}
		case "two":
			if v != 1{
				t.Errorf("Expected 1")
			}
		case "three":
			if v != 3{
				t.Errorf("Expected 3")
			}
		}
	}
}

func TestCount(t *testing.T) {
	got := Count("thor is a god")

	if got != 4{
		t.Errorf("WROOONG")
	}
}

func ExampleCount() {
	Count("thor is a god")
	//Output
	//4
}