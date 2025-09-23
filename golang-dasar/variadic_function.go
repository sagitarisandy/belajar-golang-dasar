package main

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// func main() {
// 	fmt.Println(sumAll(10, 10))
// }
