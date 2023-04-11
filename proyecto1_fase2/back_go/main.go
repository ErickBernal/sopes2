package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	puerto := "8080"
	http.HandleFunc("/home", homeHandler)
	fmt.Println("Servidor escuchando en el puerto " + puerto)
	// http.ListenAndServe(":"+puerto, nil)
	http.HandleFunc("/ram", handleCors(getRamHandler))
	log.Fatal(http.ListenAndServe(":"+puerto, nil))
}

// ------------------------------ Server ----------------------------------
type Ram struct {
	Total      float32 `json:"total"`
	Disponible float32 `json:"disponible"`
	EnUso      float32 `json:"usado"`
}

func GetRam() Ram {
	m := ReadMemoryStats()
	ram := Ram{
		Total:      KbToGb(m.MemTotal),
		Disponible: KbToGb(m.MemAvailable),
		EnUso:      KbToGb(m.MemTotal - m.MemAvailable),
	}
	// fmt.Println("--> ", ram)
	return ram
}

// GET
func getRamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GetRam())
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Golang!!!")
	fmt.Println("homeHandler")
}

func handleCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

// ------------------------------  RAM ----------------------------------------
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
