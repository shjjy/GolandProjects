package internal

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func Hello() {
	fmt.Println("Hello World")
}

func MathRand() {
	fmt.Println("Math Rand is  %d  haha", rand.Intn(10))
	fmt.Printf("Math Rand is  %d  haha\n", rand.Intn(10)) //format
	fmt.Printf("7的平方根是 %g 嗎?\n", math.Sqrt(7))
	fmt.Printf("Math Rand is  %v  haha\n", rand.Intn(10))
	fmt.Printf("7的平方根是 %v 嗎?\n", math.Sqrt(7))
	fmt.Println(math.Pi)
}

func Add() {
	fmt.Println("3 + 5 =", add(3, 5))
}

func add(i, j int) int {
	return i + j
}

func Swap() {
	//多回傳值
	a, b := swap("GO", "Lang")
	fmt.Println(a, b)
}

func swap(i, j string) (string, string) {
	return j, i
}

func Split() {
	var sq int
	var sqr float64
	sq, sqr = split(4)
	fmt.Printf("Square of 4 is: %d\n", sq)
	fmt.Printf("Square root of 4 is: %g\n", sqr)
}

func split(x int) (sq int, sqr float64) {
	sq = x * x
	sqr = math.Sqrt(float64(x))
	//直接 return，使用端造成閱讀困難，需仰賴 IDE
	return
}

// 全局與區域變數
var c, java bool

func Var() {
	//先給型態
	var i int
	fmt.Println(i, c, java)
}

var i, j int = 1, 2

func VarInit() {
	//從值判斷型態
	var c, java = true, "no!"
	fmt.Println(i, j, c, java)
}

func ShortVar() {
	var i, j int = 1, 2
	//短變量，只能在 func 內
	k := 3
	c, java := true, "no!"

	fmt.Println(i, j, k, c, java)
}

/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8

rune // int32 一個 Unicode

float32 float64

complex64 complex128
TODO: 看 typecheck.go
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i) //複數
)

func BType() {
	fmt.Printf("Type：%T value：%v\n", ToBe, ToBe)
	fmt.Printf("Type：%T value：%v\n", MaxInt, MaxInt)
	fmt.Printf("Type：%T value：%v\n", z, z)
}

func Zero() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("Type：%T value：%v\n", i, i)
	fmt.Printf("Type：%T value：%v\n", f, f)
	fmt.Printf("Type：%T value：%v\n", b, b)
	fmt.Printf("Type：%T value：%v\n", s, s)
}

func Conversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func AutoVar() {
	var i int
	j := i
	v := 42           //int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("i Type：%T value：%v\n", i, i)
	fmt.Printf("i Type：%T value：%v\n", j, j)
	fmt.Printf("i Type：%T value：%v\n", v, v)
	fmt.Printf("i Type：%T value：%v\n", f, f)
	fmt.Printf("i Type：%T value：%v\n", g, g)
}

// const 不能用 :=
const (
	Pi    = 3.14
	Big   = 1 << 100
	Small = Big >> 99
)

func ConstExample() {
	const World = "World"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
