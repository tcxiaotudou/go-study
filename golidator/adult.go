package main

import (
	"fmt"
	"github.com/golodash/galidator"
)

type Adult struct {
	FirstName string `galidator:"not_john" not_john:"$value is not allowed"`
	LastName  string
	Age       int `galidator:"min=18" json:"age" min:"$value years old users are not allowed. minimum required is $min"`
}

func notJohn(input interface{}) bool {
	return input.(string) != "John"
}

func main() {

	invalidData := []Adult{
		{
			FirstName: "Ali",
			LastName:  "Ahmadi",
			Age:       20,
		},
		{
			FirstName: "John",
			LastName:  "Brown",
			Age:       15,
		},
		{
			FirstName: "Mahmood",
			LastName:  "Heidari",
			Age:       25,
		},
	}

	var g = galidator.New()

	g.CustomValidators(galidator.Validators{"not_john": notJohn})

	var adultListValidator = g.Validator([]Adult{})

	fmt.Println(adultListValidator.Validate(invalidData))
}
