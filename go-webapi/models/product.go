package models

import (
	"database/sql"
	"fmt"
	"learn-golang/go-webapi/repository"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func query() *sql.Rows {
	db := repository.GetConnection()
	defer db.Close()
	produtos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	return produtos
}

func FindAllProdutos() []Product {
	product := Product{}
	produtos := []Product{}
	rows := query()
	for rows.Next() {
		var nome, descricao string
		var id, quantidade int
		var preco float64

		err := rows.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = nome
		product.Description = descricao
		product.Price = preco
		product.Quantity = quantidade

		produtos = append(produtos, product)
	}
	return produtos
}

func Save(p Product) (bool, error) {
	db := repository.GetConnection()
	defer db.Close()

	result, err := db.Exec("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)",
		p.Name, p.Description, p.Price, p.Quantity)

	if err == nil {
		if rows, _ := result.RowsAffected(); rows > 0 {
			fmt.Println("Save product")
		}
		return true, nil
	}

	return false, err
}

func Update(p Product) (bool, error) {
	db := repository.GetConnection()
	defer db.Close()

	fmt.Println("Id updated", p.Id)
	result, err := db.Exec("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id = $5",
		p.Name, p.Description, p.Price, p.Quantity, p.Id)

	if err == nil {
		if rows, _ := result.RowsAffected(); rows > 0 {
			fmt.Println("Product Updated")
		}
		return true, nil
	}

	return false, err
}

func Find(id int) Product {
	db := repository.GetConnection()
	defer db.Close()

	rows, err := db.Query("select * from produtos where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	produto := Product{}
	for rows.Next() {
		var nome, descricao string
		var id, quantidade int
		var preco float64

		rows.Scan(&id, &nome, &descricao, &preco, &quantidade)
		produto.Id = id
		produto.Name = nome
		produto.Description = descricao
		produto.Price = preco
		produto.Quantity = quantidade
	}
	return produto
}

func Remove(id int) (bool, error) {
	db := repository.GetConnection()

	result, err := db.Exec("DELETE FROM produtos where id = $1", id)

	row, _ := result.RowsAffected()
	defer db.Close()
	return row > 0, err
}
