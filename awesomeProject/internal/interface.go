package internal

import (
	"fmt"
	"math"
)

func MyInterface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := VertexI{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	fmt.Println(a.Abs())
}

type Abser interface {
	Abs() float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type VertexI struct {
	X, Y float64
}

func (v *VertexI) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Implement() {
	var i I = T{"Test"}
	i.M()
}

type I interface {
	M()
}
type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func InterfaceValue() {
	var i I
	i = &T{"Test"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

type T2 struct {
	S string
}

func (t *T2) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func InterfaceNilValue() {
	var i I

	var t *T3
	i = t
	describe(i)
	i.M()

	i = &T{"Test"}
	describe(i)
	i.M()
}

type T3 struct {
	S string
}

func (t *T3) M() {
	if t == nil {
		fmt.Println("<nil la>")
		return
	}
	fmt.Println(t.S)
}

func InterfaceNil() {
	var i I
	describe(i)
	i.M()
}

func InterfaceEmpty() {
	var i interface{}
	describeEmpty(i)

	i = 42
	describeEmpty(i)

	i = "hello"
	describeEmpty(i)
}

func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func Assert() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok) // ok 為 false,f 為 i 的零值

	//f = i.(float64) // panic
	//fmt.Println(f)
}

func TypeSwitch() {
	do(21)
	do("hello")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func Stringer() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func StringerPrac() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}
