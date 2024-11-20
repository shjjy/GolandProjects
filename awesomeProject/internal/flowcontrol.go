package internal

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func ForLoop() {
	sum := 0
	for i := 0; i < 11; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func WhileLoop() {
	sum := 1
	i := 0
	for sum < 1000 {
		sum += sum
		i++
	}
	fmt.Printf("sum is %d , Iteration is %d\n", sum, i)
}

func InfiniteLoop() {
	for {
		fmt.Println("Infinite loop")
		break
	}
}

func IfExample() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(math.Sqrt(2), math.Sqrt(-4))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func IfCondExample() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	}
	return limit
}

func IfElse() {
	fmt.Println(
		powElse(3, 2, 10),
		powElse(3, 3, 20),
	)
}

func powElse(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}
	return limit
}

func LoopPractice() {
	mySqrt(2)
	mySqrtC(2)
	fmt.Println(GOSqrt(2))
}

func mySqrt(x float64) float64 {
	if x < 0 {
		return math.NaN()
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %d: z = %g\n", i+1, z)
	}
	return z
}

func mySqrtC(x float64) float64 {
	if x < 0 {
		return math.NaN()
	}
	z := x / 2
	delta := 1e-10
	i := 0
	for {
		nextZ := z - (z*z-x)/(2*z)
		if math.Abs(nextZ-z) < delta {
			break
		}
		z = nextZ
		i++
		fmt.Printf("Iteration %d: z = %g\n", i, z)
	}
	return z
}

const (
	mask  = 0x7FF
	shift = 64 - 11 - 1
	bias  = 1023
)

func GOSqrt(x float64) float64 {
	// special cases
	switch {
	case x == 0 || math.IsNaN(x) || math.IsInf(x, 1):
		return x
	case x < 0:
		return math.NaN()
	}
	ix := math.Float64bits(x)
	fmt.Println(ix)
	// normalize x
	exp := int((ix >> shift) & mask)
	fmt.Println(ix)
	fmt.Println(exp)
	if exp == 0 { // subnormal x
		for ix&(1<<shift) == 0 {
			ix <<= 1
			exp--
		}
		exp++
	}
	exp -= bias // unbias exponent
	ix &^= mask << shift
	ix |= 1 << shift
	if exp&1 == 1 { // odd exp, double x to make it even
		ix <<= 1
	}
	exp >>= 1 // exp = exp/2, exponent of square root
	// generate sqrt(x) bit by bit
	ix <<= 1
	var q, s uint64               // q = sqrt(x)
	r := uint64(1 << (shift + 1)) // r = moving bit from MSB to LSB
	for r != 0 {
		t := s + r
		if t <= ix {
			s = t + r
			ix -= t
			q += r
		}
		ix <<= 1
		r >>= 1
	}
	// final rounding
	if ix != 0 { // remainder, result not exact
		q += q & 1 // round according to extra bit
	}
	ix = q>>1 + uint64(exp-1+bias)<<shift // significand + biased exponent
	return math.Float64frombits(ix)
}

func Switch() {
	fmt.Print("Go OS：")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func SwitchOrder() {
	fmt.Println("周六是哪天？")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("今天。")
	case today + 1:
		fmt.Println("明天。")
	case today + 2:
		fmt.Println("後天。")
	default:
		fmt.Println("很多天後。")
	}
}

// 另一種 if else
func SwitchWithoutCond() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("早上")
	case t.Hour() < 17:
		fmt.Println("下午")
	default:
		fmt.Println("晚上")
	}
}

// 用在關閉資源
func Defer() {
	defer fmt.Println("First")
	fmt.Println("Second")
}

func DeferFor() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
