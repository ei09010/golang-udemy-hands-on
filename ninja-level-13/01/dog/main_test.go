package dog

import "testing"

func TestYears_PositiveArgument(t *testing.T) {
	got := Years(5)
	expected := 35

	if got != expected {
		t.Errorf("Expected %v,  but got %v", expected, got)
	}
}


func TestYears_NegativeArgument(t *testing.T) {
	got := Years(-1)
	expected := 0

	if got != expected {
		t.Errorf("Expected %v,  but got %v", expected, got)
	}
}

func TestYears_DocWay(t *testing.T) {
	Years(-1)

	// Output
	// 0
}


func TestYearsTwo_PositiveArgument(t *testing.T) {
	got := YearsTwo(5)
	expected := 35

	if got != expected {
		t.Errorf("Expected %v,  but got %v", expected, got)
	}
}


func TestYearsTwo_NegativeArgument(t *testing.T) {
	got := YearsTwo(-1)
	expected := 0

	if got != expected {
		t.Errorf("Expected %v,  but got %v", expected, got)
	}
}

func TestYearsTwo_DocWay(t *testing.T) {
	YearsTwo(-1)

	// Output
	// 0
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(2)
	}
}


func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(2)
	}
}
