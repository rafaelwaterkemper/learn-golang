package main

import (
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Tenis", "Adidas", 199.99, 2},
		{"Calça", "Nike", 99.99, 10},
		{"Pc Gamer", "Bom demais", 2599.99, 1},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
