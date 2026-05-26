package cpu

import (
	"testing"
)

func TestFetch(t *testing.T) {
	cpuInfo, err := Fetch()
	if err != nil {
		t.Fatalf("Error fetching CPU info: %v", err)
	}
	if cpuInfo.Cores <= 0 {
		t.Errorf("Expected positive number of CPU cores, got %d", cpuInfo.Cores)
	}
	if cpuInfo.ModelName == "" {
		t.Error("Expected non-empty CPU model name")
	}
	if cpuInfo.UsagePercent < 0 || cpuInfo.UsagePercent > 100 {
		t.Errorf("Expected CPU usage percent between 0 and 100, got %.2f", cpuInfo.UsagePercent)
	}
}
