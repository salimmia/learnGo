package main

import "fmt"

type contactInfo struct{
	email string;
	zipCode int;
}

type person struct{
	firstName string;
	lastName string;
	contactInfo;
}

func main(){
	// var alex person

	// alex.firstName = "Alex";
	// alex.lastName = "Anderson";

	// fmt.Println(alex);
	// fmt.Printf("%+v", alex);

	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jin@gmail.com",
			zipCode: 1630,
		},
	}
	jimPointer := &jim;
	jimPointer.updateName("jimmy");
	// jim.updateName("jimmy");
	jim.print();

	fmt.Println();
}

func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName;
}

func (p person) print(){
	fmt.Printf("%+v", p);
}