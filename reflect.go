package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

type SampleWithTag struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	UserName, Address, Email string
	Age                      int
}

type PersonWithTag struct {
	UserName string `required:"true" max:"10"`
	Address  string `required:"true" max:"10"`
	Email    string `required:"true" max:"10"`
	Age      int    `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("type name =>", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
	}
}

func readFieldWithTag(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("type name =>", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

func IsValid(data any) (result bool) {
	result = true
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			d := reflect.ValueOf(data).Field(i).Interface()
			result = d != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	sample := Sample{"san"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)

	fmt.Println(sampleType.Name())
	fmt.Println(structField)
	fmt.Println(structField.Name)
	fmt.Println(structField.Type)

	fmt.Println("================")

	readField(Person{"ryuu", "medan", "san@ryuu.id", 25})

	fmt.Println("================")
	fmt.Println("struct tag")
	fmt.Println("================")

	readFieldWithTag(SampleWithTag{"ryuu"})

	fmt.Println("================")
	fmt.Println("struct tag with validation")
	fmt.Println("================")

	personWithTag := PersonWithTag{"ryuu", "medan", "san@ryuu.id", 25}
	fmt.Println(IsValid(personWithTag))

	personWithTag = PersonWithTag{"ryuu", "medan", "", 25}
	fmt.Println(IsValid(personWithTag))

	personWithTag = PersonWithTag{"ryuu", "", "", 25}
	fmt.Println(IsValid(personWithTag))
}
