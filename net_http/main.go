package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Anything goes here")
}

func main() {
	d := hotdog(1)
	http.ListenAndServe(":8080", d)

}
