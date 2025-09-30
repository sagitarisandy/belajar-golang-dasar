package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation error")
	NoFoundError    = errors.New("not found error")
)

func GetById(id string) error {

	if id == "" {
		return ValidationError
	}

	if id != "arya" {
		return NoFoundError
	}

	return nil
}

func main() {
	err := GetById("arya")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NoFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	}
}
