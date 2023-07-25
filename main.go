package main

import (
	"cursoGo/GoPractica/webserver"
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
	//files.LeoArchivo()
	//funciones.Calculos()
	//funciones.LlamarClosure()
	//funciones.Exponecia(2)
	//arreglosslices.MuestroArreglos()
	//arreglosslices.Capacidad()
	//mapas.MostrarMapas()
	//users.AltaUsuario()
	/*
		Pedro := new(modelos.Hombre)
		ejerintefaces.HumanoRespirando(Pedro)

		Maria := new(modelos.Mujer)
		ejerintefaces.HumanoRespirando(Maria)
	*/
	//defer_panic.EjemploPanic()
	/*
		canal1 := make(chan bool)
		defer func() {
			<-canal1
		}()
		go goroutines.MiNombreLento("vicente", canal1)
	*/
	//Servidor web en go
	webserver.MiWebServer()
}
