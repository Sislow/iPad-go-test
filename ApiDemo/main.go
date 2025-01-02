package main

import (
	"fmt"
	"net/http"
)

type person struct {
	sso  int
	name string
}

var group []person

func main() {
	// Title screen
	fmt.Println("Ldap Api")

	// Add some basic stuff into the slice
	addPerson(330404332, "Rezzie")
	addPerson(502722040, "Oliver")

	// print those to console
	for _, p := range group {
		fmt.Println(p)
	}

	http.HandleFunc("/", test)
	http.HandleFunc("/names", getNames)
	http.HandleFunc("/sso", getSSOS)
	http.HandleFunc("/groups", getGroups)

	http.ListenAndServe(":8080", nil)
}

func addPerson(sso int, name string) {
	p := person{sso, name}
	group = append(group, p)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "works")
}

func getGroups(w http.ResponseWriter, r *http.Request) {
	for _, g := range group {
		formattedString := fmt.Sprintf("Name: %s, SSO: %d \n", g.name, g.sso)
		fmt.Fprintf(w, formattedString)
	}
}

func getSSOS(w http.ResponseWriter, r *http.Request) {
	for _, p := range group {
		fmt.Fprintf(w, string(p.sso))
	}
}

func getNames(w http.ResponseWriter, r *http.Request) {
	for _, p := range group {
		fmt.Fprintf(w, p.name)
	}
}
