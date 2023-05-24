package main
import "fmt"


type Student struct {
	name string
	math, english float64
}
//メソッド
// func (メソッド変数 ストラクト) メソッド名(代入する引数) 戻り値 {
// 	メソッド内の処理
// }
func (s Student) avg(math, english float64) float64 {
	return (math + english) / 2
}
//valueの置き換え
func changeVertex(s *Student){
	s.name = "GoGoGo"
}

func main() {
	//初期化
	var s Student
	//パターン1
	s.name = "sato"
	s.math = 80
	s.english = 70
	fmt.Println(s)
	//パターン2
	a := Student{name:"hello", math:40, english:60}
    a_avg := a.avg(a.math, a.english)
	fmt.Println(a)
	fmt.Println(a_avg)


	//ポインターを用いて、valueの変更
	s1 := &Student{"ABABAB", 20, 41.23}
	fmt.Println(s1)
	changeVertex(s1)
	fmt.Println(*s1)
	fmt.Println(s1)

	
}