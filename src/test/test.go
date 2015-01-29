package main

import (
	"fmt"
	"reflect"
)

func main() {
	type testStruct struct {
		a string `xml:"Psersion"`
		b int
	}
	t := testStruct{a: "test"}
	fmt.Println(t)

	ref := reflect.ValueOf(t)
	fmt.Println(ref)
	fmt.Println(ref.Kind())
	fmt.Println(ref.Interface())
	fmt.Println(ref.Type())
	fmt.Println(ref.Elem())
	ty := reflect.TypeOf(t)
	fmt.Println(ty)
	f := ty.Field(0)
	fmt.Println(f.Tag)
}
