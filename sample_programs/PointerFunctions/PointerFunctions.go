package PointerFunctions

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

// InvokePointer to test the functionlity of pointers
func InvokePointer() {
	i := 1
	fmt.Println("initial", i)

	zeroval(i)
	fmt.Println("zeroval", i)

	zeroptr(&i)
	fmt.Println("zeroptr", i)

	fmt.Println("pointer", &i)
}

func TestStruct() {
	personObj := CreatePerson("Vijay", 100)

	personObj.displayPersonInfo("Newly created person object .... ")

	SetNameUsingValue(personObj)

	personObj.displayPersonInfo("Values post calling change by value method .. ")

	fmt.Println("Address of person object before changing value ... ", &personObj.age, &personObj.name, &personObj)
	SetNameUsingPointer(&personObj)

	personObj.displayPersonInfo("Values post calling change by pointer method .. ")

	fmt.Println("Address of person object after changing value ... ", &personObj.age, &personObj.name, &personObj)

}
