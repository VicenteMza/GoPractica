package main

import (
	"cursoGo/GoPractica/variables"
	"fmt"
)

func main() {
	variables.MostrarEnteros()
	variables.RestoVariables()
	estado, texto := variables.ComviertoATexto(1253)
	fmt.Println(estado)
	fmt.Println(texto)
}
