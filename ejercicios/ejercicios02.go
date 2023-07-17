package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var err error

func IngresarNumero() {
	for {
		fmt.Println("Ingrese número")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			num1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				//panic("El dato ingresado es incorrecto " + err.Error())
				return
			}
		}
		imprimirMultiplicacion(num1)
		break
	}
}

func imprimirMultiplicacion(int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, " * ", num1, " = ", (num1 * i))
	}
}
