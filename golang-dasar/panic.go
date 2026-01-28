package main

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Opps.. Error")
	}

	// endApp()
}

// func main() {
// 	runApp(true)
// 	fmt.Println("arya fadhil sagitarisandy")
// }
