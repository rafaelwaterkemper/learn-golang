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

	opened, errOpen := os.Open("streaming.txt")

	if errOpen != nil {
		panic(errOpen)
	}

	go func() {
		if _, err := io.Copy(rw, opened); err != nil {
			panic(err)
		}

		rw.Close()
		fmt.Println("Escreveu no pipe")
	}()

	w := zip.NewWriter(fileW)

	fInfo, err := opened.Stat()

	fih, _ := zip.FileInfoHeader(fInfo)
	fih.Name = "rafa.txt"
	fih.SetModTime(time.Now())
	fih.Method = zip.Deflate
	f, err := w.CreateHeader(fih)
	//To create without specify some aspects
	// f, err := w.Create("0187cd9c-0da1-4f69-9478-a04dd4f66913.txt")

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		// for {
		io.Copy(f, pr)

		errClose := w.Close()

		fileW.Close()

		if errClose != nil {
			log.Fatal(errClose)
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Final")
	}()

	// _, err = f.Write([]byte(file.Body))

	// b, err := ioutil.ReadAll(pr)
	// fmt.Println(string(b))
	// _, err = f.Write(b)

	// fileRead, _ := ioutil.ReadFile("0187cd9c-0da1-4f69-9478-a04dd4f66913")
	// _, err = f.Write(fileRead)
	fmt.Println("Escreveu os files")
	for {
	}
	// Make sure to check the error on Close.

}

// fmt.Println("Gerado files")

// }
