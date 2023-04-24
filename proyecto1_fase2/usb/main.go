package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func execComand(cmd string) {
	_, err := exec.Command("/bin/sh", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(out))
}
func leerLogs() {
	// Abre el archivo para lectura
	archivo, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer archivo.Close()

	// Crea un scanner para leer el archivo línea por línea
	scanner := bufio.NewScanner(archivo)

	// Lee el archivo línea por línea e imprime cada línea
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}
}

func main() {
	var x int
	var usu string
	var pwd string

	fmt.Println("Ingrese usuario:")
	fmt.Scanf("%s", &usu)
	fmt.Scanf("%s", &pwd)
	if usu == "erick" && pwd == "123" {
		for {
			fmt.Println("Seleccione un opcion")
			fmt.Println("	1. Montar usb")
			fmt.Println("	2. Desmontar usb")
			fmt.Println("	3. Copiar archivos")
			fmt.Println("	4. Mostrar archivo Logs")
			fmt.Println("	9. salir")
			fmt.Scanf("%d", &x)
			if x == 1 {
				execComand("./mount.sh")
				fmt.Println("<Mount usb>")
			} else if x == 2 {
				execComand("./umount.sh")
				fmt.Println("(Umount usb)")
			} else if x == 3 {
				execComand("./log.sh")
				fmt.Println("**LOG**")
			} else if x == 4 {
				leerLogs()
			} else if x == 9 {
				break
			}
		}
	} else {
		fmt.Println("**Error al ingresar usuario**")
	}

}
