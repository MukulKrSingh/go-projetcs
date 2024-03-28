package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	jim := person{"H", "A"}

	fmt.Printf("%+v", jim)

}
