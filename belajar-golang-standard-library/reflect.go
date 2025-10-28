package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var structField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		// structField.Tag.Get("required")
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("required"))
	}
}

func main() {
	readField(Sample{"Arya"})
	readField(Person{"Arya", "", ""})
}
