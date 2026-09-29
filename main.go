package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/google/uuid"
)

func main() {
	pwd, err := os.Getwd()
	panicOnError(err)
	photosPath := path.Join(pwd, "hosted")

	http.HandleFunc("GET /file/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.PathValue("id"))
		filePath := path.Join(photosPath, r.PathValue("id"))
		fmt.Println(filePath)
		bytes, _ := os.ReadFile(filePath)
		fmt.Println(len(bytes))
		w.Write(bytes)
	})

	http.HandleFunc("POST /file", func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)

		id := uuid.NewString()
		p := path.Join(photosPath, id)
		os.WriteFile(p, b, 0644)

		io.WriteString(w, fmt.Sprintf("ID: http://localhost:8080/file/%s\n", id))
	})

	http.ListenAndServe(":8080", nil)
	fmt.Println("Host on localhost:8080")
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
