//write a function walk(x interface{}, fn func(string)) which takes a struct x and calls fn for all strings fields found inside
package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fn(field.String())
	}
}
