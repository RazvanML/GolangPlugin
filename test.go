package main

import (
	"fmt"
	"reflect"
)

func test(x int, y string, z *string) {
	print(fmt.Printf("int is: %d, string is: %s, string pointer is : %s\n", x, y, *z))
}

func testMe(s string, data ...any) {
	y := test
	v := reflect.ValueOf(y)
	var vals []reflect.Value
	for _, vv := range data {
		vvv := reflect.ValueOf(vv)
		vals = append(vals, vvv)
	}
	fmt.Printf("Calling %s\n", s)
	reflect.Value.Call(v, vals)
	// test(data[0], data[1], data[2])
}

func main() {
	x := "zz"
	testMe("abc", 1, "x", &x)
	testMe("abc", "1", "x", &x)
}
