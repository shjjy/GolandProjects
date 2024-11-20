package internal

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"os"
	"strings"
	"time"
)

func ErrorFunc() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func ErrorPrac() {
	fmt.Println(errSqrt(2))
	fmt.Println(errSqrt(-2))
}

func errSqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}

type ErrNegativeSqrt float64

// fmt.Sprint(e) 會無窮迴圈，因為 fmt.Sprint(e) 處理 e 時，e 實現了 error ，Go 會自動調用 Error() 來獲取錯誤訊息
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// 第二次迴圈，slice 長度為 8 但因為 Return 6，所以可以正確回傳 eader!
func ReadPrac() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type myReader struct {
	data string
}

func ReadPrac2() {
	r := myReader{}
	b := make([]byte, 8)
	count := 0 //避免無窮迴圈
	for {
		n, err := r.read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		count++
		if count == 3 || err == io.EOF {
			break
		}
	}
}

func (r myReader) read(b []byte) (int, error) {
	// 將 b 中的每個字節設置為 'A'
	for i := range b {
		b[i] = 'A'
	}
	// 返回填滿的字節數和 nil 表示無錯誤
	return len(b), nil
}

func Rot13Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println()
}

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot13.r.Read(b)
	for i := 0; i < n; i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = (b[i]-'A'+13)%26 + 'A'
		} else if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = (b[i]-'a'+13)%26 + 'a'
		}
	}
	return n, err
}

func ImageT() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

type Image struct{}

func ImagePrac() {
	//m := Image{}
	//pic.ShowImage(m)
}
func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x ^ y), uint8(x ^ y), 255, 255}
}
