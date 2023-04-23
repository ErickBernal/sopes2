package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"syscall"
	"time"
)

//---------------------------------- DISK ----------------------------------------

// Creating structure for DiskStatus
type DiskStatus struct {
	All  uint64 `json:"All"`
	Used uint64 `json:"Used"`
	Free uint64 `json:"Free"`
}

// Function to get
// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

// Defining constants to convert size units
const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// ----------------------------------- CPU ---------------------------------------
func getCPUSample() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}

func main() {
	for {

		idle0, total0 := getCPUSample()
		time.Sleep(1 * time.Second)
		idle1, total1 := getCPUSample()

		idleTicks := float64(idle1 - idle0)
		totalTicks := float64(total1 - total0)
		cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

		// fmt.Printf("CPU usage is %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)
		fmt.Printf("CPU usage is %f%% \n", cpuUsage)

		fmt.Println("------------------------")

		// Getting filesystem statistics
		disk := DiskUsage("/")
		fmt.Println("--- Disk ---")
		fmt.Printf("	All: %.2f GB\n", float64(disk.All)/float64(GB))
		fmt.Printf("	Used: %.2f GB\n", float64(disk.Used)/float64(GB))
		fmt.Printf("	Free: %.2f GB\n", float64(disk.Free)/float64(GB))
		fmt.Println("---------------------------")
		fmt.Println("")
	}
}
