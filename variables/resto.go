package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Vicente"
	Estado = true
	Sueldo = 1987.32
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ComviertoATexto(num int) (bool, string) {
	texto := strconv.Itoa(num)
	return true, texto
}
