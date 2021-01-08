package models

import (
	"database/sql"
	"learn-golang/go-webapi/repository"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func selectProdutos() *sql.Rows {
	db := repository.ConectaDb()
	defer db.Close()
	selectProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}
	return selectProdutos
}

func FindAllProdutos() []Produto {
	produto := Produto{}
	produtos := []Produto{}
	rows := selectProdutos()
	for rows.Next() {
		var nome, descricao string
		var id, quantidade int
		var preco float64

		err := rows.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		produtos = append(produtos, produto)
	}
	return produtos
}
