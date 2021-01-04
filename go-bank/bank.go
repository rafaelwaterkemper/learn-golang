package main

import (
	"fmt"
	"learn-golang/go-bank/contas"
)

func main() {
	rafael := new(contas.ContaCorrente)
	rafael.Titular = "Rafael"
	rafael.Saldo = 1500

	gustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 2300}
	fmt.Println(gustavo)
	fmt.Println(rafael)
	executeSaque(rafael)
	executeDeposito(&gustavo)
	transferir(&gustavo, rafael)
}

func executeSaque(conta *contas.ContaCorrente) {
	_, err := conta.Sacar(300.)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Saque realizado com sucesso")
		fmt.Println("Valor atualizado", conta.Saldo)
	}
}

func executeDeposito(conta *contas.ContaCorrente) {
	_, err := conta.Depositar(350)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Depósito realizado com sucesso")
		fmt.Println("Valor atualizado", conta.Saldo)
	}
}

func transferir(contaOrigem *contas.ContaCorrente, contaDestino *contas.ContaCorrente) {
	fmt.Println("Conta Origem - Saldo atual", contaOrigem.Saldo)
	fmt.Println("Conta Destino - Saldo atual", contaDestino.Saldo)

	_, err := contaOrigem.Transferir(250, contaDestino)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Transferência realizado com sucesso")
		fmt.Println("Conta Origem - Valor atualizado", contaOrigem.Saldo)
		fmt.Println("Conta Destino - Valor atualizado", contaDestino.Saldo)
	}
}
