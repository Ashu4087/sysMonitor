package main

import (
	"fmt"
	"log"
	"os/user"
	"sysMonitor/cpu"
	"sysMonitor/vmem"
)

func PrintMemoryInfo() {
	memInfo, err := vmem.GetInfo()
	if err != nil {
		log.Fatalf("Error fetching memory info: %v", err)
	}

	fmt.Printf("Total Memory: %.2f GB\n", memInfo.TotalGB)
	fmt.Printf("Free Memory: %.2f GB\n", memInfo.FreeGB)
	fmt.Printf("Used Percent: %.2f%%\n", memInfo.UsedPercent)
}

func PrintCPUInfo() {
	// Placeholder for CPU info fetching and printing logic
	cpuInfo, err := cpu.Fetch()
	if err != nil {
		log.Fatalf("Error fetching CPU info: %v", err)
	}

	fmt.Printf("CPU Cores: %d\n", cpuInfo.Cores)
	fmt.Printf("CPU Model: %s\n", cpuInfo.ModelName)
	fmt.Printf("CPU Usage: %.2f%%\n", cpuInfo.UsagePercent)
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

	// --- PART 2: Memory Fetch Logic ---
	PrintMemoryInfo()

	fmt.Println("===============================")
	// --- PART 3: CPU Fetch Logic ---
	PrintCPUInfo()
}
