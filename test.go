package main

import "fmt"

func main() {
	var a int = 20
	var b *int

	b = &a
	fmt.Println(a)
	fmt.Println(*b)

	test(a)
}

func test(a int) {
	fmt.Println(a)
}
