package main

import (
	"fmt"
	"log"
	"os/user"

	system_memory "github.com/shirou/gopsutil/v3/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func main() {

	// --- PART 1: User & Nickname Logic ---
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error fetching current user: %v", err) //log error and exit & kill the program with exit(1) system call
	}
	sysUsername := currentUser.Username

	fmt.Printf("Hi %s, please enter your nickName: ", sysUsername)
	fmt.Println("--------------------------------------------------")

	var nickName string
	_, err = fmt.Scanln(&nickName)
	if err != nil {
		fmt.Println("\n[!] No input detected or error reading name. Continuing as 'Guest'...")
		nickName = "Guest"
	}
	fmt.Printf("\nWelcome %s! Fetching your system details...\n", nickName)
	fmt.Println("--------------------------------------------------")

	// --- PART 2: Memory Fetch Logic ---
	vm, error := system_memory.VirtualMemory()
	if error != nil {
		log.Fatalf("Error fetching memory: %v", error) //log error and exit & kill the program with exit(1) system call
	}
	totalGB := float64(vm.Total) / float64(GB)
	freeGB := float64(vm.Free) / float64(GB)
	usedPercent := vm.UsedPercent

	fmt.Println("==== System Memory ====")
	fmt.Printf("Total Memory: %.2f GB\n", totalGB)
	fmt.Printf("Free Memory: %.2f GB\n", freeGB)
	fmt.Printf("Used Percent: %.2f%%\n", usedPercent)

}
