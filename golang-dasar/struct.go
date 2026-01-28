package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello ", name, "My name is", customer.Name)
}

// func main() {
// 	var arya Customer
// 	arya.Name = "Arya Fadhil"
// 	arya.Address = "Indonesia"
// 	arya.Age = 23

// 	joko := Customer{
// 		Name:    "Joko",
// 		Address: "Bangladesh",
// 		Age:     30,
// 	}

// 	fmt.Println(arya)
// 	fmt.Println(arya.Name)

// 	fmt.Println(joko)
// 	fmt.Println(joko.Address)

// 	arya.sayHello(joko.Name)
// 	arya.sayHello("Budi")
// }
