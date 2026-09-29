package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/google/uuid"
)

type FileCreatedResponse struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

func newFileCreatedResponse(id string, url string) *FileCreatedResponse {
	r := FileCreatedResponse{Id: id, Url: url}
	return &r
}

func main() {
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "2137"
	}

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

		resp := newFileCreatedResponse(id, fmt.Sprintf("http://localhost:%s/file/%s", httpPort, id))
		bytes, _ := json.Marshal(resp)

		w.Write(bytes)
	})

	fmt.Printf("Host on localhost:%s\n\n", httpPort)
	http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
