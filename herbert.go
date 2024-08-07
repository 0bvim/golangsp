package main

import "fmt"

// like a typedef in c/cpp
type Visitor struct {
	Name string
	Age  int
}

//func main() {

//	Visitors := []Visitor{
//		Visitor{
//			Name: "Herbert",
//			Age:  11,
//		},
//		Visitor{
//			Name: "Mario",
//			Age:  15,
//		},
//		Visitor{
//			Name: "Luiggi",
//			Age:  18,
//		},
//		Visitor{
//			Name: "Nivi",
//			Age:  31,
//		},
//	}

//	for _, vis := range Visitors {

//		if vis.Age >= 18 {
//			HelloVisitor(vis.Name)
//		} else {
//			fmt.Println("Age out of range")
//		}
//	}

//	//HelloVisitor("Lais")
//	////TODO: improve with function to get input from cli

//	//pessoa := Visitor{
//	//	Name: "Nivi",
//	//	Age:  31,
//	//}

//	// pessoa.Age = 10
//}

// function declaration are case sensitive. With capital letter are visible for all.
func HelloVisitor(name string) {
	fmt.Println("Hello, " + name + "!")
}
