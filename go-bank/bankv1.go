package main

import (
	"fmt"
	"learn-golang/go-bank/clientes"
	"learn-golang/go-bank/contas"
)

func main() {
	rafael := new(contas.ContaCorrente)
	rafael.Titular = clientes.Titular{Nome: "Rafa"}
	rafael.Depositar(1500)

	gustavo := contas.ContaCorrente{Titular: rafael.Titular}
	gustavo.Depositar(200)

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
		fmt.Println("Valor atualizado", conta.ObterSaldo())
	}
}

func executeDeposito(conta *contas.ContaCorrente) {
	_, err := conta.Depositar(350)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Depósito realizado com sucesso")
		fmt.Println("Valor atualizado", conta.ObterSaldo())
	}
}

func transferir(contaOrigem *contas.ContaCorrente, contaDestino *contas.ContaCorrente) {
	fmt.Println("Conta Origem - Saldo atual", contaOrigem.ObterSaldo())
	fmt.Println("Conta Destino - Saldo atual", contaDestino.ObterSaldo())

	_, err := contaOrigem.Transferir(250, contaDestino)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Transferência realizado com sucesso")
		fmt.Println("Conta Origem - Valor atualizado", contaOrigem.ObterSaldo())
		fmt.Println("Conta Destino - Valor atualizado", contaDestino.ObterSaldo())
	}
}
