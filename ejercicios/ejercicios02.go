package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var err error
var texto string

func IngresarNumero() string {
	var texto1 string
	for {
		fmt.Println("Ingrese número")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			num1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Numero ingresado incorrecto!")
				continue
			}
			texto1 = imprimirMultiplicacion(num1)
			break
		}
	}
	return texto1
}

func imprimirMultiplicacion(int) string {
	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", num1, i, i*num1)
	}
	return texto
}
