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

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		produtos = append(produtos, produto)
	}
	return produtos
}

func Save(p Produto) (bool, error) {
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

func Update(p Produto) (bool, error) {
	db := repository.ConectaDb()
	defer db.Close()

	fmt.Println("Id atualizado", p.Id)
	result, err := db.Exec("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id = $5",
		p.Nome, p.Descricao, p.Preco, p.Quantidade, p.Id)

	if err == nil {
		if rows, _ := result.RowsAffected(); rows > 0 {
			fmt.Println("Produto Atualizado")
		}
		return true, nil
	}

	return false, err
}

func Find(id int) Produto {
	db := repository.ConectaDb()
	defer db.Close()

	rows, err := db.Query("select * from produtos where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	for rows.Next() {
		var nome, descricao string
		var id, quantidade int
		var preco float64

		rows.Scan(&id, &nome, &descricao, &preco, &quantidade)
		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}
	return produto
}

func RemoveProduct(id int) (bool, error) {
	db := repository.ConectaDb()

	result, err := db.Exec("DELETE FROM produtos where id = $1", id)

	row, _ := result.RowsAffected()
	defer db.Close()
	return row > 0, err
}
