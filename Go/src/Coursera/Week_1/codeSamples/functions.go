package math

func Sum(in ...int) (result int) {
	for _, value := range in {
		result += value
	}
	return
}

// func main() {
// 	fmt.Println(sum(1, 2, 3, 4, 5))

// 	// anonymous func
// 	func(in string) {
// 		fmt.Println(in)
// 	}("test")

// 	// anon in variable
// 	printer := func(in string) {
// 		fmt.Println(in)
// 	}

// 	printer("as variable")

// 	type strFuncType func(string)

// 	// callback function

// 	worker := func(callback strFuncType, test string) {
// 		callback(test)
// 	}

// 	worker(printer, "as callback -> test")

// 	// closure
// 	prefixer := func(prefix string) strFuncType {
// 		return func(in string) {
// 			fmt.Printf("[%s] %s\n", prefix, in)
// 		}
// 	}

// 	logger := prefixer("SUCCESS")
// 	logger("expected behaviour")
// }
