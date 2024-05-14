package main

import (
    "fmt"
    "reflect"
)

func main(){
	var a int;
	var b="10"; // this is type decide  on run time
	c := false  // this is short-hand method variable declaration
	// printf is used to formate type of printing
	fmt.Printf("this is a type of %v & b is type %v & c is type %v",reflect.TypeOf(a),reflect.TypeOf(b),reflect.ValueOf(c));
	test()
}
func test(){
	fmt.Printf("call another function /n")
}
