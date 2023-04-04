package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	for {
		m := ReadMemoryStats()
		gb := "Gb"
		kb := "Kb"
		fmt.Println("Total memory: ", m.MemTotal, kb)
		fmt.Println("Memory Available: ", m.MemAvailable, kb)
		fmt.Println("Memory used: ", m.MemTotal-m.MemAvailable, kb)

		fmt.Println("-----")
		fmt.Println("Total memory: ", KbToGb(m.MemTotal), gb)
		fmt.Println("Memory Available: ", KbToGb(m.MemAvailable), gb)
		fmt.Println("Memory used: ", KbToGb(m.MemTotal-m.MemAvailable), gb)
		fmt.Println("*********************************************************")
		time.Sleep(time.Second)
	}
}

type Memory struct {
	MemTotal     int
	MemFree      int
	MemAvailable int
}

func ReadMemoryStats() Memory {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bufio.NewScanner(file)
	scanner := bufio.NewScanner(file)
	res := Memory{}
	for scanner.Scan() {
		key, value := parseLine(scanner.Text())
		switch key {
		case "MemTotal":
			res.MemTotal = value
		case "MemFree":
			res.MemFree = value
		case "MemAvailable":
			res.MemAvailable = value
		}
	}
	return res
}

func parseLine(raw string) (key string, value int) {
	//fmt.Println(raw)
	text := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
	keyValue := strings.Split(text, ":")
	return keyValue[0], toInt(keyValue[1])
}

func toInt(raw string) int {
	if raw == "" {
		return 0
	}
	res, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}
	return res
}

func KbToGb(dato int) float32 {
	return float32(float32(dato) / 1024 / 1024)
}
