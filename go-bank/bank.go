package main

import (
	"fmt"
	"learn-golang/go-bank/clientes"
	"learn-golang/go-bank/contas"
)

type Banking interface {
	Sacar(valor float64) (bool, error)
}

func PagarBoleto(conta Banking, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

func main() {
	clienteRafa := clientes.Titular{
		Nome:      "Rafael",
		Cpf:       "1235754678",
		Profissao: "Programmer"}

	contaRafa := contas.ContaCorrente{Titular: clienteRafa, Conta: 5435, Agencia: 987}
	contaRafa.Depositar(320)
	PagarBoleto(&contaRafa, 150)
	fmt.Println(contaRafa.ObterSaldo())

	poupancaRafa := contas.ContaPoupanca{Titular: clienteRafa, Conta: 5435, Agencia: 987}
	poupancaRafa.Depositar(1200)
	PagarBoleto(&poupancaRafa, 700)
	fmt.Println(poupancaRafa.ObterSaldo())
}
