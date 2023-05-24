package main
import "fmt"

func main() {
	for i:=0; i<=4; i++ {
		fmt.Println(i)
	}
	for i:=0; i<=4; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i * 10)
	}
	for i:=0; i<=4; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i * 100)
	}
	for i:=0; i<=4; i++ {
		for j:=0; j<=4; j++ {
			fmt.Println(i, "-", j)
		}
	}

	arr := [...]int{2, 4, 6, 8, 10}
	sum := 0
	for i:=0; i<=4; i++ {
		sum += arr[i]
		fmt.Println(sum)
	}

	//range
	l := []string{"python", "go", "java"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
}