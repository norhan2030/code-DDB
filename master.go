package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type SlaveDevice struct {
	id     string
	ipAddr string
	data   string
}

var slave1 = SlaveDevice{id: "1", ipAddr: "http://192.168.43.254:8091/fasta"}
var slave2 = SlaveDevice{id: "2", ipAddr: "http://192.168.43.198:8092/fasta"}
var slave3 = SlaveDevice{id: "3", ipAddr: "http://192.168.43.32:8093/fasta"}

func make_chuncks(fileName string) {
	// read file in chunks
	f, err := os.Open(fileName)
	panicOnErrorM(err)
	defer f.Close()
	size_of_chunck, err := f.Seek(0, 2)
	size_of_chunck = int64(math.Ceil(float64(float64(size_of_chunck) / 3.0)))

	// fmt.Println(size_of_chunck)
	panicOnErrorM(err)

	f.Seek(0, 0)
	reader := bufio.NewReader(f)
	for i := 1; i < 4; i++ {

		b := make([]byte, size_of_chunck)
		numberOfBytesRead, err := reader.Read(b)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		// fmt.Println(string(b[0:numberOfBytesRead]))
		if i == 1 {
			slave1.data = string(b[0:numberOfBytesRead])
		} else if i == 2 {
			slave2.data = string(b[0:numberOfBytesRead])
		} else {
			slave3.data = string(b[0:numberOfBytesRead])
		}

	}
}
func master_as_client() {

	make_chuncks("Big_Data.txt")

	{
		b := strings.NewReader(slave1.data)
		resp, err := http.Post(slave1.ipAddr, "text/plain", b)
		panicOnErrorM(err)
		defer resp.Body.Close()
		// get stautus code
		fmt.Println("Status code:", resp.StatusCode)
	}

	{
		b := strings.NewReader(slave2.data)
		resp, err := http.Post(slave2.ipAddr, "text/plain", b)
		panicOnErrorM(err)
		defer resp.Body.Close()

		// get stautus code
		fmt.Println("Status code:", resp.StatusCode)
	}

	{
		b := strings.NewReader(slave3.data)
		resp, err := http.Post(slave3.ipAddr, "text/plain", b)
		panicOnErrorM(err)
		defer resp.Body.Close()

		// get stautus code
		fmt.Println("Status code:", resp.StatusCode)
	}

}

func main() {
    master_as_client()
	master_as_server()
}

func master_as_server() {
	http.HandleFunc("/", indexM)
	http.HandleFunc("/fasta", get_Slave_ip)
	fmt.Println("starting server")
	http.ListenAndServe(":8090", nil)

}

func indexM(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling / req")
	fmt.Fprintf(w, "Hello from Master")
}
func get_Slave_ip(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling GET req")
	// http://localhost:8090/fasta?id=0
	query := req.URL.Query()
	id := query.Get("id")
	var result string
	var err error
	var idInt int

	if id != "" {
		idInt, err = strconv.Atoi(id)
		if idInt == 0 {
			result = "0\n" + slave1.ipAddr + "\n" + slave2.ipAddr + "\n" + slave3.ipAddr
		} else if idInt == 1 {
			result = slave1.id + "\n" + slave1.ipAddr
		} else if idInt == 2 {
			result = slave2.id + "\n" + slave2.ipAddr
		} else if idInt == 3 {
			result = slave3.id + "\n" + slave3.ipAddr
		} else {
			result = "Write id value from 0 to 3"
		}
	} else {
		result = "Write id value from 0 to 3 ,example:(http://localhost:8091/fasta?id=1)"
	}
	// if we had any error return status 500 and error
	if err != nil {
		errorHandlerM(w, req, http.StatusInternalServerError, err)
		return
	}
	// set header return data
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, string(result))
}

func panicOnErrorM(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func errorHandlerM(w http.ResponseWriter, req *http.Request, status int, err error) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `{error:%v}`, err.Error())
}
