package internal

import (
	"fmt"
	"math"
)

// method 作法
func Method() {
	m := MyStr{3, 4}
	fmt.Println(m.abs())
}

type MyStr struct {
	X, Y float64
}

func (m MyStr) abs() float64 {
	return math.Sqrt(m.X*m.X + m.Y*m.Y)
}

// 一般 function
func MethodAreFunc() {
	m := MyStr{3, 4}
	fmt.Println(absFunc(m))
}

func absFunc(m MyStr) float64 {
	return math.Sqrt(m.X*m.X + m.Y*m.Y)
}

// 非 struct 也可使用
type MyFloat float64

func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func NonStr() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.abs())
}

// 答案為 50
// 移除*會變成5，無法改變原本的 m，改到複製出來的新實體，除非*指到原本的 m
func PointRec() {
	m := MyStr{3, 4}
	m.scale(10)
	fmt.Println(m.abs())
}

func (m *MyStr) scale(f float64) {
	m.X = m.X * f
	m.Y = m.Y * f
}

func scaleFunc(m *MyStr, f float64) {
	m.X = m.X * f
	m.Y = m.Y * f
}

func PointerAndFunc() {
	m := MyStr{3, 4}
	scaleFunc(&m, 10)
	fmt.Println(absFunc(m))
}

func MedPointerIndirect() {
	v := MyStr{3, 4}
	v.scale(2)        //value, pointer 都能接收
	scaleFunc(&v, 10) //只能接收 Pointer

	p := &MyStr{4, 3}
	p.scale(3)
	scaleFunc(p, 8)

	fmt.Println(v, p)

	v2 := MyStr{3, 4}
	fmt.Println(v2.abs())
	fmt.Println(absFunc(v2))

	p2 := &MyStr{4, 3}
	fmt.Println(p2.abs())
	fmt.Println(absFunc(*p2))
}

func Choosing() {
	v := &MyStr{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.abs())
	v.scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.abs())
}
