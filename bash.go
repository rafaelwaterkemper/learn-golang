package main

import (
	"fmt"
	"os"
	"os/exec"
)

// func main() {
// 	acao := exec.Command("ls", "-la")
// 	// acao.Stdout = os.Stdout
// 	// acao.Stderr = os.Stderr

// 	out, _ := acao.Output()
// 	convert := string(out)
// 	lines := strings.Split(convert, "\n")

// 	for i, val := range lines {
// 		if val == "" {
// 			return
// 		}
// 		fmt.Println("Linha", i, " possui:", val)
// 	}
// }

func main() {
	var acao *exec.Cmd = exec.Command("ls", "-la")
	acao.Stdout = os.Stdout
	acao.Stderr = os.Stderr

	acao.Start() //Utiliza uma nova "thread"

	fmt.Println("Using goroutines")
	acao.Wait()
	fmt.Println("Finished")
}
