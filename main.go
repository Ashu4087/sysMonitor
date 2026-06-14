package main

import (
	"fmt"
	"log"
	"os/user"
	"sync"

	// "sysMonitor/cpu"
	// "sysMonitor/disk"
	// "sysMonitor/vmem"
	"github.com/Ashu4087/sysMonitor/cpu"
	"github.com/Ashu4087/sysMonitor/disk"
	"github.com/Ashu4087/sysMonitor/vmem"
)

func PrintMemoryInfo() {
	memInfo, err := vmem.GetInfo()
	if err != nil {
		log.Printf("Error fetching memory info: %v", err)
		return
	}

	fmt.Printf("Total Memory: %.2f GB\n", memInfo.TotalGB)
	fmt.Printf("Free Memory: %.2f GB\n", memInfo.FreeGB)
	fmt.Printf("Used Percent: %.2f%%\n", memInfo.UsedPercent)
}

func PrintCPUInfo() {
	// Placeholder for CPU info fetching and printing logic
	cpuInfo, err := cpu.Fetch()
	if err != nil {
		log.Printf("Error fetching CPU info: %v", err)
		return
	}

	fmt.Printf("CPU Cores: %d\n", cpuInfo.Cores)
	fmt.Printf("CPU Model: %s\n", cpuInfo.ModelName)
	fmt.Printf("CPU Usage: %.2f%%\n", cpuInfo.UsagePercent)
}
func PrintDiskInfo() {
	d, error := disk.Fetch()
	if error != nil {
		log.Printf("Error fetching disk info: %v", error)
		return
	}

	fmt.Printf("Total Disk Space: %.2f GB\n", d.TotalGB)
	fmt.Printf("Free Disk Space: %.2f GB\n", d.FreeGB)
	fmt.Printf("Used Percentage: %.2f%%\n", d.UsedPercentage)
}
func main() {

	// --- PART 1: User & Nickname Logic ---
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error fetching current user: %v", err) //log error and exit & kill the program with exit(1) system call
	}
	sysUsername := currentUser.Username

	fmt.Printf("Hi %s, please enter your nickName: ", sysUsername)

	var nickName string
	_, err = fmt.Scanln(&nickName)
	if err != nil {
		fmt.Println("\n[!] No input detected or error reading name. Continuing as 'Guest'...")
		nickName = "Guest"
	}
	fmt.Printf("\nWelcome %s! Fetching your system details...\n", nickName)

	fmt.Println("===============================")
	fmt.Println("Fetching system metrics concurrently...")
	fmt.Println("===============================")
	var wg sync.WaitGroup
	wg.Add(3) // Add 3 background tasks to the WaitGroup counter

	// --- PART 2: Memory Fetch Logic ---
	go func() {
		defer wg.Done() // Mark this task as done when the function exits and counter will be decremented by 1
		PrintMemoryInfo()
		fmt.Println("===============================")
	}()

	// --- PART 3: CPU Fetch Logic ---
	go func() {
		defer wg.Done() // Mark this task as done when the function exits and counter will be decremented by 1
		PrintCPUInfo()
		fmt.Println("===============================")
	}()

	// --- PART 4: Disk Fetch Logic ---
	go func() {
		defer wg.Done() // Mark this task as done when the function exits and counter will be decremented by 1
		PrintDiskInfo()
		fmt.Println("===============================")

	}()
	wg.Wait() // Wait for all tasks to complete before exiting the program
	fmt.Print(" All metrics fetched successfully!")
}
