package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"time"
)

func main() {
	exibeIntroducao()
	exibeOpcoes()
	opcao := coletaOpcaoEscolhida()
	fmt.Println("Opção escolhida", opcao)
	executaAcaoEscolhida(opcao)
}

func exibeIntroducao() {
	nome := "Rafael"
	fmt.Println("Variavel nome é do tipo", reflect.TypeOf(nome))
	fmt.Println("Olá Sr, Rafa")
	fmt.Println("Executando no OS", runtime.GOOS)

}

func exibeOpcoes() {
	fmt.Println("Escolha uma das opções abaixo")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")
}

func coletaOpcaoEscolhida() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func executaAcaoEscolhida(opcao int) {
	switch opcao {
	case 1:
		fmt.Println("Iniciando monitoramento")
		executaMontioramento()
	case 2:
		fmt.Println("Exibindo logs")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Opção desconhecida")
		os.Exit(-1)
	}
}

func executaMontioramento() {
	var sites [2]string
	sites[0] = "https://www.alura.com.br"
	sites[1] = "https://www.google.com"
	// var sites []string = []string{"https://www.alura.com.br"}
	for {
		for _, site := range sites {
			fmt.Println("Monitorando ", site)
			resp, err := http.Get(site)

			if err == nil {
				fmt.Println("Site", site, " retornou o status code ", resp.StatusCode)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
