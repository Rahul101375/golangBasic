package main

import (
	"fmt"
    "reflect"
)

const (first = iota
       second
	third)
const (a1 = iota
b1
c='d'
d
e=2
f)

// const(a := iota) The error in the code a:=iota is that iota is a keyword in Golang that can
//  only be used within constant declarations. It cannot be used to declare variables.

func main(){
	a:=10;
	b:=20;
	fmt.Println(first,second,third);
	fmt.Println(a1,b1,c,d,e,f)
	// not if you change a1 into a & b1 into b its give 10 & 20 not printing const value its print main values
	fmt.Println(test(a,b));
	switchFun(first);
	switchFun(third);
	switchFun(c)
}
func test(number1 int,number2 int)(int,error){
	return number1+number2,nil
}
// nil :-> nil is a frequently used and important predeclared identifier in Go. It is the literal representation of zero values of many kinds of types.


func switchFun(a int){
	fmt.Println("a is type %v",reflect.TypeOf(a))
	// if rune('a') == a{
	// 	fmt.Println("A is true")
	// }
	// if rune('d') == a{
	// 	fmt.Println("A is true")
	// }else{
	// 	fmt.Println("A is false")
	// }
	// if rune('c') == a{
	// 	fmt.Println("c is true")
	// }else if 0==a{
	// 	fmt.Println("first is true")
	// }else{
	// 	fmt.Println("else case")
	// }
}