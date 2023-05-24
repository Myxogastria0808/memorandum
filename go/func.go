package main
import "fmt"

func sayHello(greeting string) {
	fmt.Println(greeting)
}
func cal(x, y int) int {
	return(x / y)
}

func cal2(x, y int) (int, int) {
	a := x + y
	b := x * y 
	return a, b
}

func cal3(price, item int) (result int) {
	result = price * item
	return
}
//クロージャー関数
func incrementGenerator() (func() int) {
	x := 0
	return func() int{
		x++
		return x
	}
}

func circleArea(pi float64) (func(radius float64) float64) {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func main() {
	sayHello("Hello")
	result := cal(12,3)
	fmt.Println(result)
	result1, result2 := cal2(12,70)
	fmt.Println(result1)
	fmt.Println(result2)

	cal3 := cal3(10, 2)
	fmt.Println(cal3)

	//無名関数
	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("inner func2", x)
	}(1)


	//クロージャー関数
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	pi1 := circleArea(3)
	fmt.Println(pi1(5))

	pi2 := circleArea(3.14)
	fmt.Println(pi2(5))
}