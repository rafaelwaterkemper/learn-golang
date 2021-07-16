package main

import (
	"net/http"
	"html/template"
	"time"
	"strings"
	"os/exec"
	"log"
)

type InfoService struct {
	Name, Version string 
}

func main() {
	var stable []InfoService = []InfoService{}
	loaded := make(chan bool)
	var alreadyClosed bool

	go func() {

		for {
			log.Println("Carregando infos...")

			services := exec.Command("bash", "-c", "kubectl get svc --selector 'tier=api' --template '{{range .items}}{{.metadata.name}}{{\"\\n\"}}{{end}}'  | grep -v \"hotfix\\|feature\\|bugfix\\|[0-9]\\|release\\|master\"")

			outServices, err := services.Output()
			if err != nil {
				log.Println("Falha ao capturar as versões")
			}

			returned := strings.Split(string(outServices), "\n")
			
			infosService := make([]InfoService, len(returned) -1)

			for i := 0; i < len(returned) -1; i++ {
				
				version := exec.Command("bash", "-c", "kubectl describe svc " + returned[i] + " | sed -nE 's/^[[:blank:]]*(version=)(.*)/\\2/p' | head -1")
				outVersion, err := version.Output()

				if err != nil {
					log.Println("Falha ao retornar versão para o svc ", returned[i])
					continue
				}

				svcVersion := InfoService{returned[i], string(outVersion)}
				infosService[i] = svcVersion
			}

			stable = infosService
			
			if alreadyClosed != true {
				loaded <- true
			}
			
			time.Sleep(3 * time.Second)
		}
	}()

	<- loaded
	close(loaded)
	alreadyClosed = true

	var temp = template.Must(template.ParseGlob("templates/*.html"))
	
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "Index", stable)
	})

	http.ListenAndServe(":8080", nil)
}