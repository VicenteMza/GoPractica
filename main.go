package main

import (
	"cursoGo/GoPractica/files"
)

func main() {
	/*
						variables.MostrarEnteros()
						variables.RestoVariables()
						estado, texto := variables.ComviertoATexto(1253)
						fmt.Println(estado)
						fmt.Println(texto)
					/////////////////////
					if os := runtime.GOOS; os == "linux." || os == "os x." {
						fmt.Println("Esto no es windows, es ", os)
					} else {
						fmt.Println("Esto es windows")
					}

					switch os := runtime.GOOS; os {
					case "linux":
						fmt.Println("Esto es Linux")
					case "Darwin":
						fmt.Println("Esto es darwin")
					default:
						fmt.Println("\n", os)
					}

					entero1, texto := ejercicios.ConvNumerico("95")
					fmt.Print(entero1)
					fmt.Println(texto)

					entero2, texto := ejercicios.ConvNumerico("150")
					fmt.Print(entero2)
					fmt.Println(texto)
				//////////
				teclado.IngresoNumero()
			//////
			iteraciones.Iterar()
		/////
		fmt.Println(ejercicios.IngresarNumero())
	*/
	//files.GrabaTabla()
	//files.SumaTabla()
	files.LeoArchivo()
}
