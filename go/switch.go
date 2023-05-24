package main
import "fmt"


func GG() string {
	return "GG"
}

func main() {
	a := "a"
	switch a {
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
	case "c":
		fmt.Println("c")
	}

	os := GG()
	switch os {
	case "GG":
		fmt.Println("GG")
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
	}
}