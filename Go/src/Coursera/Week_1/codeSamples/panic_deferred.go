package main

import "fmt"

func getSomeVars() string {
	fmt.Println("execution")
	return "result"
}

func main() {
	defer func() {
		fmt.Println("in panic")
		if err := recover(); err != nil {
			fmt.Println("panic happend:", err)
		}
	}()

	defer fmt.Println("after")
	defer func() {
		fmt.Println(getSomeVars())
	}()

	panic("test")
	fmt.Println("some work")

}
