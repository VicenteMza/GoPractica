package funciones

import "fmt"

func Exponecia(valor int) {
	if valor > 10000000 {
		return
	}
	fmt.Println(valor)
	Exponecia(valor * 2)
}
