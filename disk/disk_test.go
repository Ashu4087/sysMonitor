package disk

import (
	"testing"
)

func TestFetch(t *testing.T) {
	d, error := Fetch()
	if error != nil {
		t.Fatalf("Error fetching disk info: %v", error)
	}
	if d.TotalGB <= 0 {
		t.Errorf("Expected total disk space to be greater than 0, got %.2f GB", d.TotalGB)
	}
	if d.FreeGB < 0 {
		t.Errorf("Expected free disk space to be non-negative, got %.2f GB", d.FreeGB)
	}
	if d.UsedPercentage < 0 || d.UsedPercentage > 100 {
		t.Errorf("Expected used percentage to be between 0 and 100, got %.2f%%", d.UsedPercentage)
	}
}
