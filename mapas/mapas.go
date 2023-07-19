package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	//fmt.Println(paises)

	paises["Mexico"] = "D:F"
	paises["Argentina"] = "Buenos Aires"
	//fmt.Println(paises)
	//fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"River Plate": 40,
		"Barcelona":   39,
		"Real Madrid": 38,
		"Chivas":      37,
	}

	fmt.Println(campeonato)

	/*for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}*/
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["River Plate"]
	fmt.Printf("El mejor puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
}
