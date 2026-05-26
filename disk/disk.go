package disk

import (
	"fmt"

	sys_disk "github.com/shirou/gopsutil/v3/disk"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type DiskInfo struct {
	TotalGB        float64 `json:"total_gb"`
	FreeGB         float64 `json:"free_gb"`
	UsedPercentage float64 `json:"used_percentage"`
}

func Fetch() (DiskInfo, error) {
	d, err := sys_disk.Usage("/")
	if err != nil {
		return DiskInfo{}, fmt.Errorf("error fetching disk info: %v", err)
	}
	return DiskInfo{
		TotalGB:        float64(d.Total) / float64(GB),
		FreeGB:         float64(d.Free) / float64(GB),
		UsedPercentage: d.UsedPercent,
	}, nil

}
