package cpu

import (
	"fmt"
	"log"
	"time"

	sys_cpu "github.com/shirou/gopsutil/v3/cpu"
)

type CPUInfo struct {
	Cores        int     `json:"cores"`
	ModelName    string  `json:"model_name"`
	UsagePercent float64 `json:"usage_percent"`
}

func Fetch() (CPUInfo, error) {
	cores, error := sys_cpu.Counts(true)
	if error != nil {
		return CPUInfo{}, fmt.Errorf("error fetching CPU cores: %v", error)
	}
	info, error := sys_cpu.Info()
	modelName := "Unknown"
	if error != nil {
		log.Printf("Error fetching CPU model name: %v", error)
	} else if len(info) > 0 {
		modelName = info[0].ModelName
	}
	percent, error := sys_cpu.Percent(time.Second, false)
	if error != nil {
		return CPUInfo{}, fmt.Errorf("error fetching CPU usage percent: %v", error)
	}
	usagePercent := 0.0
	if len(percent) > 0 {
		usagePercent = percent[0]
	}
	return CPUInfo{
		Cores:        cores,
		ModelName:    modelName,
		UsagePercent: usagePercent,
	}, nil
}
