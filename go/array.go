package main

import ("fmt")


func main() {
	a1 := [3]string{"sato", "suzuki", "takahashi"}
	//要素数の省略
	b1 := [...]string{"a", "b","c"}
	c1 := [2][3]string{{"a", "b", "c"}, {"1", "2", "3"}}
	fmt.Println(a1[0])
	fmt.Println(a1[1])
	fmt.Println(a1[2])
	fmt.Println(b1[0])
	fmt.Println(b1[1])
	fmt.Println(b1[2])
	fmt.Println(c1[0][0])
	fmt.Println(c1[0][1])
	fmt.Println(c1[0][2])

	//スライスは、appendeできる
	//以下、スライス
	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)
	var c []int = []int{1,2,3,4,5}
	fmt.Println(c)
	fmt.Println(c[2:4])
	fmt.Println(c[:3])
	fmt.Println(c[4:])
	fmt.Println(c[2:])
	fmt.Println(c[:])

	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)

	//makeでスライスの初期化
	n := make([]int, 3, 5)
	fmt.Printf("len=%d, cap=%d, value=%v\n", len(n), cap(n), n)
	n = append(n, 0, 0)
	fmt.Printf("len=%d, cap=%d, value=%v\n", len(n), cap(n), n)
	n = append(n, 0, 0, 1, 2, 34, 5)
	fmt.Printf("len=%d, cap=%d, value=%v\n", len(n), cap(n), n)
	//makeを1つだけ指定するとlen=capになる
	m := make([]int, 3)
	fmt.Printf("len=%d, cap=%d, value=%v\n", len(m), cap(m), m)
	//初期化の宣言
	o := make([]int, 0)
	fmt.Printf("len=%d, cap=%d, value=%v\n", len(o), cap(o), o)

	//初期化宣言の違いによる挙動の差異
	a := make([]int, 5)
	for i := 0; i < 5; i++ {
		a = append(a, i)
		fmt.Println(a)
	}
	fmt.Println(a)

	// a1 := make([]int, 0, 5)
	// for j := 0; j < 5; j++ {
	// 	a1 = append(a1, j)
	// 	fmt.Println(a1)
	// }
	// fmt.Println(a1)
}