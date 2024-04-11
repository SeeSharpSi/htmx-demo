package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	add_routes(mux)
	server := http.Server{
		Addr:    "127.0.0.1:9779",
		Handler: mux,
	}
	fmt.Println("server is running")
	err := server.ListenAndServe()
	defer server.Close()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func add_routes(mux *http.ServeMux) {
	mux.HandleFunc("/", GetIndex)
	mux.HandleFunc("/button", GetButton)
	mux.HandleFunc("/button2", GetButton2)
	mux.HandleFunc("/htmx.min.js", ServerHtmx)
}

func ServerHtmx(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /htmx.min.js request\n")
	http.ServeFile(w, r, "./htmx.min.js")
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	http.ServeFile(w, r, "./index.html")
}

func GetButton(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /button request\n")
	http.ServeFile(w, r, "./button_response.html")
}

func GetButton2(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /button2 request\n")
	http.ServeFile(w, r, "./button_response2.html")
}
