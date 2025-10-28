package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Arya Fadhil", "Arya"))
	fmt.Println(strings.Split("Arya Fadhil", " "))
	fmt.Println(strings.ToLower("Arya Fadhil"))
	fmt.Println(strings.ToUpper("Arya Fadhil"))
	fmt.Println(strings.Trim("    Arya Fadhil   ", " "))
	fmt.Println(strings.ReplaceAll("Arya Fadhil", "Budi", "Arya"))
}
