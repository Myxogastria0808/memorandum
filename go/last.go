package main

import "fmt"

type User struct {
	name string
	weight, height float64
}
func (u User) cal(weight, height float64) float64 {
	return u.weight / u.height /u.height * 10000
}
func main() {
	user001 := User{name: "YukiOsada", weight: 40, height: 60.5}
	user001_BMI := user001.cal(user001.weight, user001.height)
	fmt.Println(user001)
	fmt.Println(user001_BMI)
}