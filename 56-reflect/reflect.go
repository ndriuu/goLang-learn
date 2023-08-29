package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type OtherSample struct {
	Name        string `required:"true"`
	Description string `required:"true"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Ghufron"}
	var sampleType reflect.Type = reflect.TypeOf(sample)
	structField := sampleType.Field(0)
	required := structField.Tag.Get("required")
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(required)
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	sample.Name = ""
	fmt.Println(IsValid(sample))
	OtherSample := OtherSample{"Ghufron", "College"}
	fmt.Println(IsValid(OtherSample))

}
