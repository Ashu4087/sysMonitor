package vmem

import (
	"testing"
)

func TestGetInfo(t *testing.T) {
	memInfo, err := GetInfo()
	if err != nil {
		t.Fatalf("Error fetching memory info: %v", err)
	}
	if memInfo.TotalGB <= 0 {
		t.Errorf("Expected positive total memory, got %.2f GB", memInfo.TotalGB)
	}
	if memInfo.FreeGB < 0 {
		t.Errorf("Expected non-negative free memory, got %.2f GB", memInfo.FreeGB)
	}
	if memInfo.UsedPercent < 0 || memInfo.UsedPercent > 100 {
		t.Errorf("Expected used percent between 0 and 100, got %.2f%%", memInfo.UsedPercent)
	}
}
