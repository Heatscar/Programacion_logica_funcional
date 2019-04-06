package main

import (
	"fmt"
	"time"
)

func main() {
	var n, m uint64
	fmt.Print("n: ")
	fmt.Scanln(&n)
	fmt.Print("m: ")
	fmt.Scanln(&m)
	if n<m { goRoutines(n, m) } else { fmt.Print("'n' debe ser menor a 'm'") }
	fmt.Scanln()
}

func goRoutines(n uint64, m uint64) {
	auxn, auxm := n+((m-n)/2), n+((m-n)/2)+1
	go buscarPrimos(n, auxn)
	go buscarPrimos(auxm, m)
}

func buscarPrimos(n uint64, m uint64) {
	start := time.Now()
	var array []uint64
	cont := 0
	for x:=n; x<=m; x++ {
		if esPrimo(x) { 
			array = append(array, x)
			cont++
		}
	}
	elapsed := time.Since(start)
	fmt.Println("\n", array, "\nExecution time (", n, "-", m, "): " , elapsed)
}

func esPrimo(n uint64) bool {
	var cont uint64 = 0
	for i:=uint64(1); i<=n; i++{
		if n%i == 0 { cont++ }
	}
	if cont==2 { return true } else { return false }
}
