package contas

import (
	"errors"
	"learn-golang/go-bank/clientes"
)

type ContaPoupanca struct {
	Titular clientes.Titular //iniciando com letra mínuscula retem acesso somente ao pacote
	Conta   int              //Letra maiúscula exporta para os demais pacotes
	Agencia int
	saldo   float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) (bool, error) {
	if c.saldo <= valorDoSaque && valorDoSaque > 0 {
		return false, errors.New("saldo insuficiente")
	}
	c.saldo -= valorDoSaque
	return true, nil
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (bool, error) {
	if valorDeposito <= 0 {
		return false, errors.New("Valor do depósito deve ser maior que zero")
	}
	c.saldo += valorDeposito
	return true, nil
}

func (c *ContaPoupanca) Transferir(valorTransferencia float64, contaDestino *ContaPoupanca) (bool, error) {
	if valorTransferencia <= 0 {
		return false, errors.New("O valor da transferencia deve ser maior que zero")
	}
	c.saldo -= valorTransferencia
	contaDestino.Depositar(valorTransferencia)
	return true, nil
}

//Encapsulamento, tornar variável package private, e criar "Getter"
func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
