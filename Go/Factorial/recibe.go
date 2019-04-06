package main

import ("fmt"
       "html/template"
    "net/http"
    "strconv"
)

func recibe(w http.ResponseWriter, r *http.Request) {
    fmt.Println("metodo:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("form.html")
        t.Execute(w, nil)
    } else {
        t, _ := template.ParseFiles("form.html")
        t.Execute(w, nil)
        r.ParseForm()
        num, err := strconv.Atoi(r.FormValue("num"))
        if err != nil {}
        fmt.Fprint(w, "<center><h1>", num, "! = ", factorial(num) ,"</h1></center>")
    }
}

func factorial (num int) int {
    cont := 1
    for i:=1; i<=num; i++ {
        cont *= i
    }
    return cont
}

func main(){
    PORT := ":8080"
    fmt.Printf("http://localhost%s/\n", PORT)
    http.HandleFunc("/", recibe)
    http.ListenAndServe(PORT, nil)
}

