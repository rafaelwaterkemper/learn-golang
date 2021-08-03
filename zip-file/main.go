package main

import (
	"archive/zip"
	"fmt"
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

		// Create a buffer to write our archive to.
		fileW, _ := os.Create(file.Name + ".zip")

		// Create a new zip archive.
		w := zip.NewWriter(fileW)

		f, err := w.Create(file.Name + ".txt")
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
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
