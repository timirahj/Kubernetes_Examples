package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil) // set listen port
}

func HelloServer(rw http.ResponseWriter, req *http.Request) {

	fmt.Fprint(rw, "Hello World!!")
}
