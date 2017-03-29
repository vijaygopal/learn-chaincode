package main

import (
	"fmt"

	pointerFunction "github.com/vijaygopal/learn-chaincode/sample_programs/PointerFunctions"
)

func main() {

	nextint := intSeq()

	fmt.Println(nextint())
	fmt.Println(nextint())
	fmt.Println(nextint())
	fmt.Println(nextint())

	newInt := intSeq()
	fmt.Println(newInt())

	pointerFunction.InvokePointer()
	pointerFunction.TestStruct()
	builderPattern()
}

func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func builderPattern() {
	builder := pointerFunction.New()
	car := builder.TopSpeed(50).Paint(pointerFunction.BLUE).Build()

	fmt.Println(car.Drive())
	fmt.Println(car.Stop())
}
