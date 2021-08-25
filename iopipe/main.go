package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {

	fileW, _ := os.Create("0187cd9c-0da1-4f69-9478-a04dd4f66913" + ".zip")

	pr, rw := io.Pipe()

	go func() {
		rw.Write([]byte("Rafael Marangoni"))
		time.Sleep(time.Second * 1)
		rw.Write([]byte("Waterkemper"))
		rw.Close()
	}()

	w := zip.NewWriter(fileW)

	f, err := w.Create("rafa.txt")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		buffer := make([]byte, 8)

		for {
			size, err := pr.Read(buffer)

			if err == io.EOF {
				fmt.Print("Nao há mais dados para ler")
				break
			}

			_, err = f.Write(buffer[:size])

		}

		errClose := w.Close()

		fileW.Close()

		if errClose != nil {
			log.Fatal(errClose)
		}

		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Escreveu os files")
	for {
	}

}
