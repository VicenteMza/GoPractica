package ejercicios

import "strconv"

func ConvNumerico(valor string) (int, string) {
	entero, _ := strconv.Atoi(valor)
	var texto string

	if entero > 100 {
		texto = " Es mayor a 100"
	} else {
		texto = " Es menor a 100"
	}

	return entero, texto
}
