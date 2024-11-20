package internal

import (
	"fmt"
	"sync"
	"time"
)

func ConcurrentTest() {
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Channels() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func BuffChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func ChannelRC() {
	c := make(chan int, 10)
	go fibonacci2(cap(c), c) //cap取得 c 的容量 => 10
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) //一定要有，否則 for i := range c 會無窮迴圈 deadlock
}

func ChannelSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci3(c, quit)
}

func fibonacci3(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func DefaultSelect() {
	tick := time.Tick(100 * time.Millisecond)  //每 0.1 秒
	boom := time.After(500 * time.Millisecond) //0.5秒後
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func BinaryTree() {
	ch := make(chan int)
	go Walk(New(1), ch)

	// 印從 channel 中接收的值
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("Testing Walk function:")

	ch2 := make(chan int)
	go Walk(New(1), ch2)
	for i := range ch2 {
		fmt.Println(i)
	}

	// 測試 Same 函數
	fmt.Println("Testing Same function:")
	fmt.Println("Same(tree.New(1), tree.New(1)):", Same(New(1), New(1))) // 應返回 true
	fmt.Println("Same(tree.New(1), tree.New(2)):", Same(New(1), New(2))) // 應返回 false
}

// Walk 歷遍樹 t，樹中所有值送到 ch。
func Walk(t *Tree, ch chan int) {
	// 使用遞迴的方式中序遍歷二叉查找樹
	var walkHelper func(t *Tree)
	walkHelper = func(t *Tree) {
		if t == nil {
			return
		}
		walkHelper(t.Left)  // 歷左樹
		ch <- t.Value       // 發送當前節點的值
		walkHelper(t.Right) // 歷右樹
	}
	walkHelper(t)
	close(ch) // 關閉 channel
}

// Same 判断 t1 和 t2 是否包含相同的值。
func Same(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	// 同時歷兩樹
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// 比較兩個 channel 中接收到的值
	for v1 := range ch1 {
		v2, get := <-ch2
		if !get || v1 != v2 { // 如果值不同或 ch2 提前關閉，返回 false
			return false
		}
	}

	// 檢查 ch2 是否還有剩餘元素
	_, get := <-ch2
	return !get // 兩個 channel 都空了才返回 true
}

func Mutex() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock 一次一個 goroutine +1
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock 一次一個 goroutine 讀值
	defer c.mu.Unlock()
	return c.v[key]
}

func WebCrawler() {
	Crawl("https://golang.org/", 4, fetcher)
}

type Fetcher interface {
	// Fetch 回傳 URL 頁面的 body 內容
	// 將該頁面上找到的所有 URL 放到切片
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 用 fetcher 從 URL 遞迴爬頁面，直到達最大深度
func Crawl(url string, depth int, fetcher Fetcher) {
	// 確保並行抓取的 goroutines 完成
	var wg sync.WaitGroup
	// 記錄已經爬過的 URL
	visited := make(map[string]bool)

	// 用一個函數來實現並行抓取
	var crawlHelper func(url string, depth int)
	crawlHelper = func(url string, depth int) {
		// 保證每次遞迴結束後減少 WaitGroup 計數
		defer wg.Done()

		// 如果深度小於等於 0 或已經爬過該 URL，結束
		if depth <= 0 || visited[url] {
			return
		}

		// 標記為已訪問
		visited[url] = true

		// 使用 fetcher 抓取內容
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)

		// 對所有找到的 URL 開始新的 Crawl，並行進行
		for _, u := range urls {
			wg.Add(1)
			go crawlHelper(u, depth-1)
		}
	}

	// 開始爬
	wg.Add(1)
	go crawlHelper(url, depth)

	// 等所有 goroutines 完成
	wg.Wait()
}

// 假數據結構和抓取，官網提供
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
