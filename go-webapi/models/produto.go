package models

import (
	"database/sql"
	"fmt"
	"learn-golang/go-webapi/repository"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func queryProdutos() *sql.Rows {
	db := repository.ConectaDb()
	defer db.Close()
	selectedProducts, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}
	return selectedProducts
}

func FindAllProdutos() []Produto {
	produto := Produto{}
	produtos := []Produto{}
	rows := queryProdutos()
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

func SaveProduto(p Produto) (bool, error) {
	db := repository.ConectaDb()
	defer db.Close()

	result, err := db.Exec("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)",
		p.Nome, p.Descricao, p.Preco, p.Quantidade)

	if err == nil {
		if rows, _ := result.RowsAffected(); rows > 0 {
			fmt.Println("Produto salvo")
		}
		return true, nil
	}

	return false, err
}
