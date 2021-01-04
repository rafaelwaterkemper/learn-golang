package main

import "fmt"

type ContaCorrente struct {
	titular string
	conta   int
	agencia int
	saldo   float64
}

func main() {
	buildObjects()
	comparation()
}

func buildObjects() {
	var contaRafa *ContaCorrente
	contaRafa = new(ContaCorrente) //Retorna um ponteiro do objeto
	contaRafa.titular = "Rafael"
	fmt.Println(*contaRafa) //Acessa o conteúdo do ponteiro, o objeto literal

	contaLu := new(ContaCorrente) //Retorna o ponteiro com o short assignment
	contaLu.titular = "Lu"
}

func comparation() {
	contaGuga1 := ContaCorrente{titular: "Guga", saldo: 150.55} //Retorna o objeto em si
	contaGuga2 := ContaCorrente{titular: "Guga", saldo: 150.55} //Retorna o objeto em si

	fmt.Println("Comparing by value is ", contaGuga1 == contaGuga2)

	var contaRafa1 *ContaCorrente
	contaRafa1 = new(ContaCorrente) //Retorna um ponteiro do objeto
	contaRafa1.titular = "Rafael"

	var contaRafa2 *ContaCorrente
	contaRafa2 = new(ContaCorrente) //Retorna um ponteiro do objeto
	contaRafa2.titular = "Rafael"

	fmt.Println("Comparing by ref is ", contaRafa1 == contaRafa2)
	fmt.Println("Comparing by value is ", *contaRafa1 == *contaRafa2)

	sameRafa1 := contaRafa1
	fmt.Println("Comparing by same ref is ", contaRafa1 == sameRafa1)

	fmt.Printf("sameRafa1=%p ownAdress=%p \n", sameRafa1, &sameRafa1) //Print value of pointer, and, your own address
	fmt.Printf("contaRafa1=%p ownAdress=%p \n", contaRafa1, &contaRafa1)
}
