package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

func main() {
	PORT := ":8080"
	fmt.Printf("Servidor corriendo en http://localhost%s/\n", PORT)
	fmt.Printf("Ejemplo a = 1, b = -8, c = 16 Resultado X = 4")
	http.HandleFunc("/", render)
	http.ListenAndServe(PORT, nil)
}

func render(w http.ResponseWriter, r *http.Request) {
	fmt.Println("metodo:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
		r.ParseForm()
		//Ejemplo a = 1, b = -8, c = 16 Resultado X = 4
		a, err := strconv.Atoi(r.FormValue("a"))
		b, err := strconv.Atoi(r.FormValue("b"))
		c, err := strconv.Atoi(r.FormValue("c"))
		if err != nil {
		}
		fmt.Fprint(w, raices(float64(a), float64(b), float64(c)))
	}
}

func raices(a float64, b float64, c float64) string {
	var res string
	discriminante := math.Pow(b, 2) - (4 * a * c)
	var resultado float64
	if discriminante > 0 {
		resultado = (-b + math.Sqrt(discriminante)) / (2 * a)
		res = ("<h1>X<sub>1&nbsp;</sub> = " + fmt.Sprintf("%.2f", resultado) + "</h1>")
		resultado = (-b - math.Sqrt(discriminante)) / (2 * a)
		res += ("<h1>X<sub>2&nbsp;</sub> = " + fmt.Sprintf("%.2f", resultado) + "</h1>")
	} else if discriminante == 0 {
		resultado = (-b + math.Sqrt(discriminante)) / (2 * a)
		res = ("<h1>X<sub>1&nbsp;</sub> = " + fmt.Sprintf("%.2f", resultado) + "</h1>")
	} else {
		res = "<h1>No hay numeros imaginarios</h1>"
	}
	return res
}
