package days

import "testing"

func TestDay1_count(t *testing.T) {
	data := []string{"1", "2", "0", "1"}
	
	count := day1_count(data)

	if count != 2 {
		t.Error("Expected 2, got", count)
	}
}
