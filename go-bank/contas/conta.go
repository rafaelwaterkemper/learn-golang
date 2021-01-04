package contas

import "errors"

type ContaCorrente struct {
	Titular string //iniciando com letra mínuscula retem acesso somente ao pacote
	Conta   int    //Letra maiúscula exporta para os demais pacotes
	Agencia int
	Saldo   float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (bool, error) {
	if c.Saldo <= valorDoSaque && valorDoSaque > 0 {
		return false, errors.New("Saldo insuficiente")
	}
	c.Saldo -= valorDoSaque
	return true, nil
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (bool, error) {
	if valorDeposito <= 0 {
		return false, errors.New("Valor do depósito deve ser maior que zero")
	}
	c.Saldo += valorDeposito
	return true, nil
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) (bool, error) {
	if valorTransferencia <= 0 {
		return false, errors.New("O valor da transferencia deve ser maior que zero")
	}
	c.Saldo -= valorTransferencia
	contaDestino.Depositar(valorTransferencia)
	return true, nil
}
