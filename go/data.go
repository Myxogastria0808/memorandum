package main

import ("fmt")
import ("reflect")

func main(){
	var num1 int = 123
	var num2 float64 = 1.23
	var num3 = 1
	var string1 = "Hello, World!"
	var bool1 bool = num1 > num3
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(string1)
	fmt.Println(bool1)
	fmt.Println(reflect.TypeOf(num1))
	fmt.Println(reflect.TypeOf(num2))
	fmt.Println(reflect.TypeOf(string1))
	fmt.Println(reflect.TypeOf(bool1))
}