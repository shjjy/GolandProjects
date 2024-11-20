package internal

import (
	"fmt"
	"math"
	"strings"
)

func Pointers() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	fmt.Println(p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func Struct() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
	p := &v
	p.Y = 20
	fmt.Println(v)
}

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func Literals() {
	fmt.Println(v1, p, v2, v3)
}

func Array() {
	var a [2]string
	a[0] = "Yang"
	a[1] = "Chang"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

}

func Slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	c := names[0:2]
	d := names[1:3]
	fmt.Println(c, d)

	d[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names) //slice 會修改原 Array

	q := []int{2, 3, 5, 7, 11, 13} //沒宣告長度就是 Slice
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	ss := []struct {
		ii int
		bb bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(ss)

	st := []int{2, 3, 5, 7, 11, 13}

	st = st[1:4]
	fmt.Println(st)

	st = st[:2]
	fmt.Println(st)

	st = st[1:]
	fmt.Println(st)

	//cap(acity):容量，可以裝幾個、len(gth):長度，實際裝多少
	si := []int{2, 3, 5, 7, 11, 13}
	printSlice(si)

	// 截取切片使長度 0
	si = si[:0]
	printSlice(si)

	// 加長
	si = si[:4]
	printSlice(si)

	// 捨棄前二
	si = si[2:]
	printSlice(si)

	si = si[:4]
	printSlice(si)

	var sNil []int
	fmt.Println(sNil, len(sNil), cap(sNil))
	if sNil == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func MakeSlice() {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
	printSlice2("b", b)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func SliceOfSlice() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func AppendSlice() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only).
// range 拿來 loop slice, map
func Range() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	//index, value
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i)
	}
	//不需要 index 時，硬寫編譯會錯
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}

func SlicePractice() {
	//pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	// 創建外層切片，大小為 dy（高度）
	image := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		// 創建內層切片，大小為 dx（寬度）
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// 使用 x * y 計算像素值，並轉換為 uint8
			row[x] = uint8(x * y)
		}
		// 將計算完的 row 添加到外層切片中
		image[y] = row
	}

	return image
}

type myStruct struct {
	x, y float64
}

var m map[string]myStruct

func MapParc() {
	m = make(map[string]myStruct)
	m["str A"] = myStruct{1.0, 2.0}
	fmt.Println(m["str A"])
	MapParc2()
}

var m2 = map[string]myStruct{
	"str A": myStruct{1.0, 2.0},
	"str B": myStruct{2.0, -4.0},
}

var m3 = map[string]myStruct{
	"str A": {1.0, 2.0},
	"str B": {2.0, -4.0},
}

func MapParc2() {
	fmt.Println(m2)
	fmt.Println(m3)
}

func MutatingMap() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	m["Answer"] = 48
	v2, ok2 := m["Answer"]
	fmt.Println("The value:", v2, "Present?", ok2)
}

func MapExercise() {
	//wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	// 初始化一個字串到整數的映射
	wordCount := make(map[string]int)

	// 分割字串 s，並取得所有單詞
	words := strings.Fields(s)

	// 計算每個單詞的出現次數
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func test(s string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}

func FunValue() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func Closure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func Fibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		result := x
		x, y = y, x+y
		return result
	}

}
