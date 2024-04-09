package main

import (
	"fmt"
	"net/http"
)

func d(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Anything goes here :  Dog")
}
func c(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Anything goes here : Cat")
}
func cat(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Anything goes here : Cat Meows")
}

func main() {

	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)
	http.HandleFunc("/cat/mewo", cat)
	http.ListenAndServe(":8080", nil) //http already has Dafault Mux in place and we can use it by passing nil handler of ListenAndServe

}
