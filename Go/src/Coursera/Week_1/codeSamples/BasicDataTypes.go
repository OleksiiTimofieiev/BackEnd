package main

import (
	"fmt"
)

const (
	helloConst = "Yo"
	e          = 2.718
)

const (
	zero = iota
	_    // ignoring value
	three
)

const (
	// const without a type
	year = 2017

	yearTyped int = 2017
)

type UserID int

func main() {
	{
		var num0 int // default value is 0

		var num1 int = 1

		var num2 = 2

		num3 := 3

		var weight, height int = 19, 20

		weight, height = 4, 5

		fmt.Println(num0, num1, num2, num3, weight, height)

		// rune == utf-8

		helloWorld := "HelloWorld"

		hello := helloWorld[:5]
		byteString := []byte(helloWorld)
		hello2 := string(byteString)
		fmt.Println(hello, byteString, hello2, zero)
	}
	/////
	{
		idx := 1
		var uid UserID = 42

		myID := UserID(idx)

		println(uid, myID)
	}
	///// no address arithmetic, pointer is a different data type
	///// arrays/slices
	{
		var a1 [3]int // [0,0,0]

		const size = 2 // size of array has to be of constant value

		var a2 [2 * size]bool
		fmt.Println(a1, a2)

		a3 := [...]int{1, 2, 3}

		fmt.Println(a3)

		// slice == array + capacity (x2) , kinda vector
		buf5 := make([]int, 5, 10)

		fmt.Println(buf5)

		var buf []int

		buf = append(buf, 9, 10)

		fmt.Println(buf)

		var user map[string]string = map[string]string{
			"name":      "a",
			"last name": "b",
		}

		fmt.Println(user)
	}
	/////
	{
		/*
			for {

			}

			for isRun{

			}

			for i :=0; i<2; i++

			for idx := range sl{

			}

			for key/_,val/_ := range map {

			}

			for pos,char := range str {

			}
		*/
	}
	/////
}
