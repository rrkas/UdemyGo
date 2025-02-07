package main

import (
	"fmt"
)

type person struct {
	fname, lname string
	age          int
}

func main() {
	//optional fields
	p1 := person{
		fname: "James",
		lname: "Bond",
	}
	//named fields
	p2 := person{
		fname: "Rohnak",
		lname: "Agarwal",
		age:   20,
	}
	//positional fields
	p3 := person{
		"R",
		"A",
		20,
	}
	// p4 := person{
	// 	"R",
	// 	"A",
	// 	age: 20,
	// } //error: mixture of field:value and value initializers
	fmt.Println(p1, p2, p3)
	fmt.Println(p1.fname, p2.lname, p2.age)

	//embeded struct
	type student struct {
		person
		roll int
	}

	s1 := student{
		person: person{
			fname: "Rohnak",
			lname: "Agarwal",
			age:   20,
		},
		roll: 30,
	}

	fmt.Println(s1.person.fname)
	fmt.Println(s1.fname)

	//same field name of inner and outer structs
	type student2 struct {
		person
		fname string
		roll  int
	}

	s2 := student2{
		person: person{
			fname: "Rohnak",
			lname: "Agarwal",
			age:   20,
		},
		fname: "abc",
		roll:  30,
	}

	fmt.Println(s2.person.fname)
	fmt.Println(s2.fname)

	//anonymous struct
	a1 := struct {
		a, b int
	}{
		a: 10,
		b: 20,
	}
	fmt.Println(a1, a1.a, a1.b)
}
