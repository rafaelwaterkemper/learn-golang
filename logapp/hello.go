// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"reflect"
// 	"runtime"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func main() {
// 	exibeIntroducao()
// 	exibeOpcoes()
// 	opcao := coletaOpcaoEscolhida()
// 	fmt.Println("Opção escolhida", opcao)
// 	executaAcaoEscolhida(opcao)
// }

// func exibeIntroducao() {
// 	nome := "Rafael"
// 	fmt.Println("Variavel nome é do tipo", reflect.TypeOf(nome))
// 	fmt.Println("Olá Sr, Rafa")
// 	fmt.Println("Executando no OS", runtime.GOOS)

// }

// func exibeOpcoes() {
// 	fmt.Println("Escolha uma das opções abaixo")
// 	fmt.Println("1 - Iniciar Monitoramento")
// 	fmt.Println("2 - Exibir logs")
// 	fmt.Println("0 - Sair do Programa")
// }

// func coletaOpcaoEscolhida() int {
// 	var comando int
// 	fmt.Scan(&comando)

// 	return comando
// }

// func executaAcaoEscolhida(opcao int) {
// 	switch opcao {
// 	case 1:
// 		fmt.Println("Iniciando monitoramento")
// 		executaMontioramento()
// 	case 2:
// 		fmt.Println("Exibindo logs")
// 		exibeLogs()
// 	case 0:
// 		fmt.Println("Saindo do programa")
// 		os.Exit(0)
// 	default:
// 		fmt.Println("Opção desconhecida")
// 		os.Exit(-1)
// 	}
// }

// func executaMontioramento() {
// 	// var sites [2]string
// 	// sites[0] = "https://www.alura.com.br"
// 	// sites[1] = "https://www.google.com"
// 	// sites := []string{"https://www.alura.com.br"}
// 	// sites = append(sites, "https://www.google.com")
// 	// sites = append(sites, "https://www.terra.com.br/")
// 	// sites = append(sites[:1], sites[2:]...) //Remover itens do slice, uso o indice 0 e do 2 em diante, excluindo o indíce 1
// 	sites := leArquivoDeSites()

// 	for {
// 		for i, site := range sites {
// 			fmt.Println("Monitorando ", i, site)
// 			resp, err := http.Get(site)

// 			if err != nil {
// 				fmt.Println("Erro ao consultar site", site, ", error ", err)
// 				registraLogs(site, false)
// 				continue
// 			}

// 			if resp.StatusCode != 200 {
// 				registraLogs(site, false)
// 			} else {
// 				fmt.Println("Site", site, " retornou o status code ", resp.StatusCode)
// 				registraLogs(site, true)
// 			}
// 		}
// 		time.Sleep(10 * time.Second)
// 	}
// }

// func leArquivoDeSites() []string {
// 	sites := []string{}

// 	file, err := os.Open("sites.txt") //retorna o enredeço de memória do arquivo
// 	// file, err := ioutil.ReadFile("sitex.txt") Le o arquivo inteiro string(file) - converte para string
// 	if err != nil {
// 		fmt.Println("Falha ao abrir o arquivo", err)
// 	}
// 	reader := bufio.NewReader(file)

// 	for {
// 		linha, err := reader.ReadString('\n')
// 		linha = strings.TrimSpace(linha)
// 		sites = append(sites, linha)
// 		if err == io.EOF {
// 			break
// 		}
// 	}

// 	file.Close()
// 	return sites
// }

// func registraLogs(site string, status bool) {
// 	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

// 	if err != nil {
// 		fmt.Println("Ocorreu um erro", err)
// 	}

// 	pointer := &site

// 	file.WriteString("Data/Hora=" + time.Now().Format("02/01/2006 15:04:05") + ";Site=" + site + ";Status=" + strconv.FormatBool(status) + "\n")
// }

// func exibeLogs() {
// 	file, err := ioutil.ReadFile("logs.txt")

// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}

// 	fmt.Println(string(file))
// }
