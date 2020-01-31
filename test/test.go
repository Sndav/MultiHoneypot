package main

import (
	"fmt"
	"reflect"
)

type Map struct {
	Keys  map[string]interface{}
	Types map[string]reflect.Type
}

func main() {
	a := "123456"
	fmt.Print(a[4:])
}
