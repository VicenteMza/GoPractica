package ejerintefaces

import (
	"cursoGo/GoPractica/interfaces"
	"fmt"
)

func HumanoRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}
