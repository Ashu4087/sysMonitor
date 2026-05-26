package vmem

import (
	"fmt"

	system_memory "github.com/shirou/gopsutil/v3/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type MemoryInfo struct {
	TotalGB     float64 `json:"total_gb"`
	FreeGB      float64 `json:"free_gb"`
	UsedPercent float64 `json:"used_percent"`
}

func GetInfo() (MemoryInfo, error) {
	vm, err := system_memory.VirtualMemory()
	if err != nil {
		return MemoryInfo{}, fmt.Errorf("error fetching memory: %v", err)
	}
	return MemoryInfo{
		TotalGB:     float64(vm.Total) / float64(GB),
		FreeGB:      float64(vm.Free) / float64(GB),
		UsedPercent: vm.UsedPercent,
	}, nil
}
