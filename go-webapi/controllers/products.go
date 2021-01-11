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
	name := r.FormValue("name")
	description := r.FormValue("description")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))

	result, err := models.Save(models.Product{Name: name, Description: description, Price: price, Quantity: quantity})

	if err != nil {
		temp.ExecuteTemplate(w, "New", err.Error())
	}

	if result {
		http.Redirect(w, r, "/", 301)
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	models.Remove(id)
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	description := r.FormValue("description")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))

	result, _ := models.Update(models.Product{Id: id, Name: name, Description: description, Price: price, Quantity: quantity})

	if result {
		http.Redirect(w, r, "/", 301)
	}

}
