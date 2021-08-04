package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme", "This archive contains some text files."},
		{"gopher", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {

		fmt.Println(file)
		// Create a buffer to write our archive to.
		fileW, _ := os.Create("0187cd9c-0da1-4f69-9478-a04dd4f66913" + ".zip")
		// fileW, _ := os.Open("0187cd9c-0da1-4f69-9478-a04dd4f66913")

		// Create a new zip archive.
		w := zip.NewWriter(fileW)

		// f, err := w.Create(file.Name + ".txt")
		f, err := w.Create("0187cd9c-0da1-4f69-9478-a04dd4f66913.jpeg")
		if err != nil {
			log.Fatal(err)
		}
		// _, err = f.Write([]byte(file.Body))

		fileRead, _ := ioutil.ReadFile("0187cd9c-0da1-4f69-9478-a04dd4f66913")
		_, err = f.Write(fileRead)
		fmt.Println("Escreveu os files")

		// Make sure to check the error on Close.
		errClose := w.Close()
		if errClose != nil {
			log.Fatal(errClose)
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Gerado files")

}
