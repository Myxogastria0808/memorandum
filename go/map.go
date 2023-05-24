package main

import "fmt"

func main() {
	//辞書式
	//m := map[キー]値{"キー値": 値, ...}
	m := map[string]int{"apple": 100, "banana": 200 }
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	//キーに値があるかどうかを2つめの引数で確認できる
	fmt.Println(m["nothing"])
	v, ok := m["apple"]
	fmt.Println(v, ok)
	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	//varでの宣言では。メモリが確保されないから問題
	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

}