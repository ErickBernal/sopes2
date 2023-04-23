package main

import (
	"fmt"
	"log"
	"os/exec"
)

func execComand(cmd string) {
	out, err := exec.Command("/bin/sh", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func main() {
	var x int
	for {
		fmt.Println("Seleccione un opcion")
		fmt.Println("	1. Montar usb")
		fmt.Println("	2. Desmontar usb")
		fmt.Println("	3. Mostrar Log")
		fmt.Println("	9. salir")
		fmt.Scanf("%d", &x)
		if x == 1 {
			execComand("./mount.sh")
			fmt.Println("<Mount usb>")
		} else if x == 2 {
			execComand("./umount.sh")
			fmt.Println("(Umount usb)")
		} else if x == 3 {

		} else if x == 9 {
			break
		}
	}

}
