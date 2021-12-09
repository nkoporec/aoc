package days

import "testing"

func TestDay1_count(t *testing.T) {
	data := []string{"1", "2", "0", "1"}
	
	count := day1_count(data)

	if count != 2 {
		t.Error("Expected 2, got", count)
	}
}


func TestDay1_second_count(t *testing.T) {
	data := []string{"199", "200", "208", "210", "200","207", "19"}
	
	count := day1_second_count(data)

	if count != 1 {
		t.Error("Expected 1, got", count)
	}
}
