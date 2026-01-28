package main

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr " + man.Name
}

// func main() {
// 	arya := Man{"Arya"}
// 	arya.Married()

// 	fmt.Println(arya.Name)
// }
