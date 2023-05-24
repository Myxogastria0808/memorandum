package main
import (
	"fmt"
	"os/user"
	"time"
	"strconv"
)


//init関数
func init() {
	fmt.Println("Init！")
}
//関数外で動く変数宣言
//パターン1
var i int = 12
//パターン2
var (
	u8 uint8 = 255
	f32 float32 = 0.53
	c64 complex64 = -5 + 12i
)


//不変変数
const (
	Username = "test"
	Password = "Good job"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(time.Now())
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Println(user.Current())
	fmt.Println(u8, f32, c64)
	fmt.Printf("type=%T, value=%v\n", c64, c64)
	fmt.Printf(string("Hello World"[0]))
	fmt.Printf(`"`)

	var s string = "14"
	//_はエラーハンドリングの変数を使わないという宣言
	iii, _ := strconv.Atoi(s)
	fmt.Printf("type=%T, value=%v\n", iii, iii)
}
