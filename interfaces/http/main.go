package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	///Code foe getting the response Body from a request
	// bs := make([]byte, 9999)

	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//OR

	///Code foe getting the response Body from a request
	io.Copy(os.Stdout, resp.Body)

}
