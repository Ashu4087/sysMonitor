package disk

import (
	"fmt"
	"sysMonitor/units"

	sys_disk "github.com/shirou/gopsutil/v3/disk"
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
		TotalGB:        float64(d.Total) / float64(units.GB),
		FreeGB:         float64(d.Free) / float64(units.GB),
		UsedPercentage: d.UsedPercent,
	}, nil

}
