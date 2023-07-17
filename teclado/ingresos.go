package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var leyenda string
var err error

func IngresoNumero() {
	fmt.Println("Ingrese número")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese número 2")
	scanner = bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}
	fmt.Println("Ingrese leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, num1*num2)
}
