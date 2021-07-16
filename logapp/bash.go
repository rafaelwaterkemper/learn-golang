// package main

// import (
// 	"fmt"
// )

// func main() {
// 	nomes := []string{"Rafa", "Gustavo"}

// 	var nomeSelecionado string

// 	for _, nome := range nomes {
// 		fmt.Printf("%p", &nome)
// 		fmt.Println(nome)
// 		if nome == "Rafa" {
// 			fmt.Println("[E o Rafa")
// 			nomeSelecionado = nome
// 		}
// 	}
// 	fmt.Printf("%p", nomeSelecionado)
// 	fmt.Println("\n", nomeSelecionado)
// }

// // func main() {
// // 	acao := exec.Command("bash", "-c", "kubectl describe svc api-folha | sed -nE 's/^[[:blank:]]*(version=)(.*)/\\2/p' | head -1")
// // 	// acao.Stdout = os.Stdout
// // 	// acao.Stderr = os.Stderr
// // 	fmt.Println(acao)
// // 	out, _ := acao.Output()
// // 	println(string(out))
// // 	convert := string(out)
// // 	lines := strings.Split(convert, "\n")

// // 	for i, val := range lines {
// // 		if val == "" {
// // 			return
// // 		}
// // 		fmt.Println("Linha", i, " possui:", val)
// // 	}
// // }

// // func main() {
// // 	var acao *exec.Cmd = exec.Command("bash", "-c", "kubectl get pods | grep api-folha-calculo")
// // 	acao.Stdout = os.Stdout
// // 	acao.Stderr = os.Stderr

// // 	acao.Start() //Utiliza uma nova "thread"

// // 	fmt.Println("Using goroutines")
// // 	acao.Wait()
// // 	fmt.Println("Finished")
// // }
