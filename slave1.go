package main

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/fasta", handleFasta)
	http.HandleFunc("/", index)
	fmt.Println("starting server")
	http.ListenAndServe(":8091", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	file_bytes, err := ioutil.ReadFile("slave1.fasta")
	fmt.Println("Handling / req")
	fmt.Fprintf(w, string(file_bytes))
	if err != nil {
		errorHandler(w, req, http.StatusInternalServerError, err)
		return
	}
}

func handleFasta(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		// fmt.Fprintf(w,"Nothing to Get yet")
		get(w, req)
	} else if req.Method == "POST" {
		post(w, req)
	} else {
		fmt.Println("Handling invalid /users req")
		errorHandler(w, req, http.StatusMethodNotAllowed, fmt.Errorf("Invalid Method"))
	}
}

func get(w http.ResponseWriter, req *http.Request) {
	file_bytes, err := ioutil.ReadFile("slave1.fasta")
	fmt.Println("Handling / req")
	fmt.Fprintf(w, string(file_bytes))
	if err != nil {
		errorHandler(w, req, http.StatusInternalServerError, err)
		return
	}

}

func post(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling POST req")
	defer req.Body.Close()

	// read req body
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errorHandler(w, req, http.StatusInternalServerError, err)
		return
	}

	err = ioutil.WriteFile("slave1.fasta", b, 0644)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}

func errorHandler(w http.ResponseWriter, req *http.Request, status int, err error) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `{error:%v}`, err.Error())
}
