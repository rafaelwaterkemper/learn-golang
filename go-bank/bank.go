package main

import (
	"errors"
	"fmt"
)

type ContaCorrente struct {
	titular string
	conta   int
	agencia int
	saldo   float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (bool, error) {
	if c.saldo <= valorDoSaque && valorDoSaque > 0 {
		return false, errors.New("Saldo insuficiente")
	}
	c.saldo -= valorDoSaque
	return true, nil
}

func main() {
	rafael := ContaCorrente{titular: "Rafael", saldo: 1500.}

	_, err := rafael.Sacar(300.)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Saque realizado com sucesso")
		fmt.Println("Valor atualizado", rafael.saldo)

	}
}
