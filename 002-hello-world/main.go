package main

import "fmt"

func main() {
	fmt.Println("Hello")
	foo()
	fmt.Println("Goodbye")
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo")
}
