package days

import "testing"

func TestDay2_calculate(t *testing.T) {
	data := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	
	result := day2_calculate(data)

	if result != 150 {
		t.Error("Expected 150, got", result)
	}

}

func TestDay2_second_calculate(t *testing.T) {
	data := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	
	result := day2_second_calculate(data)

	if result != 900 {
		t.Error("Expected 900, got", result)
	}
}
