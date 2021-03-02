package main

import "fmt"

func main() {
	fmt.Println("Hello everyone!")

	//foo()

	for i := 0; i <= 100; i++ {
		if i % 5 == 0 {
			fmt.Println(i)
		}
	}

	//bar()
}

func foo() {
	fmt.Println("I am in foo")
}

func bar() {
	fmt.Println("And now we exited the code")
}