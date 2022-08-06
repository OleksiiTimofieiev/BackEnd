package main

import (
	"fmt"
	"net/http"
)

func main() {
	request1()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcome to Go web API")
	fmt.Printf("Endpoint: Homepage")

}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	who := "otimofie"

	fmt.Fprintf(w, "test")
	fmt.Println("Endpoint: ", who)
}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutMe", aboutMe)
	http.ListenAndServe(":8080", nil)

}