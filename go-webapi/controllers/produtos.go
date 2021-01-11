package controllers

import (
	"html/template"
	"learn-golang/go-webapi/models"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.FindAllProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	product := models.Find(id)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
	quantidade, _ := strconv.Atoi(r.FormValue("quantidade"))

	result, err := models.Save(models.Produto{Nome: nome, Descricao: descricao, Preco: preco, Quantidade: quantidade})

	if err != nil {
		temp.ExecuteTemplate(w, "New", err.Error())
	}

	if result {
		http.Redirect(w, r, "/", 301)
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	models.RemoveProduct(id)
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
	quantidade, _ := strconv.Atoi(r.FormValue("quantidade"))

	result, _ := models.Update(models.Produto{Id: id, Nome: nome, Descricao: descricao, Preco: preco, Quantidade: quantidade})

	if result {
		http.Redirect(w, r, "/", 301)
	}

}
