package main

import "fmt"

type Ktp struct {
	Nik, Age      int
	Name, Address string
}

type SomeFunc interface {
	ShowNik() int
}

func SaySomething(theFunc SomeFunc) {
	fmt.Println("Nik is", theFunc.ShowNik())
}

func (theKtp Ktp) ShowNik() int {
	return theKtp.Nik
}

// func (person Ktp) checkNik() {
// 	if person.Nik > 0 {
// 		fmt.Println("Nik is true")
// 	} else {
// 		fmt.Println("Nik is false")
// 	}
// }
func main() {
	// people := map[string]bool{"Anna": true, "Bob": false, "Chris": true}
	// for key, value := range people {
	// 	if value {
	// 		fmt.Println(key, "is a true value")
	// 	} else if !value {
	// 		fmt.Println(key, "is a false value")
	// 	}
	// }
	dedi := Ktp{
		Nik:     22,
		Age:     20,
		Name:    "Dedi",
		Address: "Jakarta",
	}
	// SaySomething(dedi)
	// dodot := dedi
	dodot := &dedi
	*dodot = Ktp{Nik: 22, Age: 20, Name: "Dodot", Address: "Jakarta"}
	dodo := dedi

	fmt.Println(dodot)
	fmt.Println(dodo)
	// dedi.checkNik()

}
