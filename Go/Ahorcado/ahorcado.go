package main

import 	(
	"fmt"
	t "time"
	r "math/rand"
)

var palabras = [7] string {"avion", "barco", "arbol", "carro", "lamina", "luna", "internet"}
var n uint
var guiones, aciertos int
var palabra, letra string
var intentos int = 6


func main() {

	ran := r.NewSource(t.Now().UnixNano())
	palabra := palabras[r.New(ran).Intn(len(palabras))]
	guiones := len(palabra)
	npalabra := ""
	estado := true


	for i := 0 i < guiones; i ++ {
		npalabra += "_ "
	}

		fmt.Println("-------------")
		fmt.Println("|        (o_o)")
		fmt.Println("|          |")
		fmt.Println("|         /|\\")
		fmt.Println("|          |")
		fmt.Println("|          |")
		fmt.Println("|         / \\")
		fmt.Println("|        /   \\")
		fmt.Println(npalabra)

	for  intentos != 0 {
		fmt.Print("\nEscribe una letra: ")
		fmt.Scanln(&letra)
		npalabra, estado = comparar(palabra, npalabra, letra)
		if !estado { intentos = intentos -1 }
		if (npalabra==palabra){
			fmt.Println("Adivinaste la palabra")
			break;
		}
	
		switch intentos{
			case 6:
				fmt.Print("Intentos realizados: ")
				fmt.Println(intentos)
				fmt.Println("-------------")
				fmt.Println("|        (o_o)")
				fmt.Println("|          |")
				fmt.Println("|         /|\\")
				fmt.Println("|          |")
				fmt.Println("|          |")
				fmt.Println("|         / \\")
				fmt.Println("|        /   \\")
				fmt.Println(npalabra)
				break
			
			case 5:
		
				fmt.Print("Intentos realizados: ")
				fmt.Println(intentos)

				fmt.Println("-------------")
				fmt.Println("|        (o_o)")
				fmt.Println("|          |")
				fmt.Println("|         /|\\")
				fmt.Println("|        / | \\")
				fmt.Println("|          |")
				fmt.Println("|         /")
				fmt.Println("|        /")
				fmt.Println(npalabra)
				break
			
			case 4:
		
				fmt.Print("Intentos realizados: ")
				fmt.Println(intentos)

				fmt.Println("-------------")
				fmt.Println("|        (o_o)")
				fmt.Println("|          |")
				fmt.Println("|         /|\\")
				fmt.Println("|        / | \\")
				fmt.Println("|          |")
				fmt.Println(npalabra)
			case 3:
		
				fmt.Print("Intentos realizados: ")
				fmt.Println(intentos)

				fmt.Println("-------------")
				fmt.Println("|        (o_o)")
				fmt.Println("|          |")
				fmt.Println("|         /|")
				fmt.Println("|        / |")
				fmt.Println("|          |")
				fmt.Println(npalabra)
				break
			case 2:
		
				fmt.Print("Intentos realizados: ")
				fmt.Println(intentos)

				fmt.Println("-------------")
				fmt.Println("|        (o_o)")
				fmt.Println("|          |")
				fmt.Println("|          |")
				fmt.Println("|          |")
				fmt.Println("|          |")
				fmt.Println(npalabra)
				break
			case 1:
		
				fmt.Print("Intentos realizados: ")
				fmt.Println(intentos)

				fmt.Println("-------------")
				fmt.Println("|        (o_o)")
				fmt.Println("|")
				fmt.Println("|")
				fmt.Println("|")
				fmt.Println("|")
				fmt.Println(npalabra)
				break
			case 0:
		
				fmt.Print("Intentos realizados: ")
				fmt.Println(intentos)

				fmt.Println("-------------")
				fmt.Println("|           ")
				fmt.Println("|")
				fmt.Println("|")
				fmt.Println("|")
				fmt.Println("|")
				fmt.Println(npalabra)
				fmt.Println("La palabra era: ", palabra)
				fmt.Println("FIN DEL JUEGO")
				break
		}
	}
}

func comparar(palabra string, npalabra string, letra string) (let string, estado bool){
	let = ""
	estado = false
	tamaño := len(palabra)

	arregloguiones := make ([]string,tamaño)
	arreglonuevo := make([]string,tamaño)

	for i := 0; i < tamaño; i ++{
		arregloguiones[i] = string (palabra[i])
		arreglonuevo[i] = string(npalabra[i])
	}
	for i := 0; i < tamaño; i ++{
		if (letra == arregloguiones[i]){
			arreglonuevo[i] = letra
			estado = true
		}
		let += string (arreglonuevo[i])	
	}
	return
}