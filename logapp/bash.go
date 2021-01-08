package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	acao := exec.Command("bash", "-c", "kubectl describe svc api-folha | sed -nE 's/^[[:blank:]]*(version=)(.*)/\\2/p' | head -1")
	// acao.Stdout = os.Stdout
	// acao.Stderr = os.Stderr
	fmt.Println(acao)
	out, _ := acao.Output()
	println(string(out))
	convert := string(out)
	lines := strings.Split(convert, "\n")

	for i, val := range lines {
		if val == "" {
			return
		}
		fmt.Println("Linha", i, " possui:", val)
	}
}

// func main() {
// 	var acao *exec.Cmd = exec.Command("bash", "-c", "kubectl get pods | grep api-folha-calculo")
// 	acao.Stdout = os.Stdout
// 	acao.Stderr = os.Stderr

// 	acao.Start() //Utiliza uma nova "thread"

// 	fmt.Println("Using goroutines")
// 	acao.Wait()
// 	fmt.Println("Finished")
// }
