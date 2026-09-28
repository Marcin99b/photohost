package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	pwd, err := os.Getwd()
	panicOnError(err)

	photosPath := path.Join(pwd, "input")
	fmt.Println(photosPath)

	dir, err := os.ReadDir(photosPath)
	panicOnError(err)

	for _, file := range dir {
		fmt.Println(file.Name())
	}

	http.HandleFunc("GET /file/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.PathValue("id"))
		filePath := path.Join(photosPath, r.PathValue("id"))
		fmt.Println(filePath)
		bytes, _ := os.ReadFile(filePath)
		fmt.Println(len(bytes))
		w.Write(bytes)
	})

	http.HandleFunc("POST /file", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Test\n")
	})

	http.ListenAndServe(":8080", nil)
	fmt.Println("Host on localhost:8080")
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
