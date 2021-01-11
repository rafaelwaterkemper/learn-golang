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
	products, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}
	return products
}

func FindAllProdutos() []Product {
	product := Product{}
	products := []Product{}
	rows := query()
	for rows.Next() {
		var name, description string
		var id, quantity int
		var price float64

		err := rows.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)
	}
	return products
}

func Save(p Product) (bool, error) {
	db := repository.GetConnection()
	defer db.Close()

	result, err := db.Exec("INSERT INTO products (name, description, price, quantity) VALUES($1, $2, $3, $4)",
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
	result, err := db.Exec("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id = $5",
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

	rows, err := db.Query("select * from products where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	produto := Product{}
	for rows.Next() {
		var name, description string
		var id, quantity int
		var price float64

		rows.Scan(&id, &name, &description, &price, &quantity)
		produto.Id = id
		produto.Name = name
		produto.Description = description
		produto.Price = price
		produto.Quantity = quantity
	}
	return produto
}

func Remove(id int) (bool, error) {
	db := repository.GetConnection()

	result, err := db.Exec("DELETE FROM products where id = $1", id)

	row, _ := result.RowsAffected()
	defer db.Close()
	return row > 0, err
}
