package main

import "fmt"

type Human struct {
	Name string
}

func (h Human) Walk() {
	fmt.Println(h.Name, "is walking")
}

func main() {
	human := Human{
		Name: "Jacob",
	}

	fmt.Printf("%+v \n", human)

	human.Walk()
}
