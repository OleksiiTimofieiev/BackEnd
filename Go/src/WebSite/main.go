package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s.", u.name)
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{
		name:       "Bob",
		age:        25,
		money:      -50,
		avg_grades: 4.2,
		happiness:  0.8,
	}
	fmt.Fprintf(w, bob.getAllInfo())
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")

}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {

	handleRequest()
}
