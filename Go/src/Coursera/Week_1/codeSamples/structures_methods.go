package main

import "fmt"

type Detailed struct {
	age    int
	weight int
}

type Person struct {
	Id      int
	Cleaner func(string)
	Detailed
}

func test(in string) {
	fmt.Println(in)
}

func (p Person) changeId_copy(id int) {
	p.Id = id
}

func (p *Person) changeId(id int) {
	p.Id = id
}

func main() {
	var person Person = Person{
		Id:      5,
		Cleaner: test,
		Detailed: Detailed{
			age:    15,
			weight: 45,
		},
	}

	person.Cleaner("test")
	person.changeId(89)

	fmt.Println(person)

}
