package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl1 uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl1(url string, depth int, fetcher Fetcher, out chan string, end chan bool) {
	if depth <= 0 {
		end <- true
		return
	}

	if _, ok := crawled[url]; ok {
		end <- true
		return
	}
	crawledMutex.Lock()
	crawled[url] = true
	crawledMutex.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		out <- fmt.Sprintln(err)
		end <- true
		return
	}

	out <- fmt.Sprintf("found: %s %q\n", url, body)
	subEnd := make(chan bool)
	for _, u := range urls {
		go Crawl1(u, depth-1, fetcher, out, subEnd)
	}

	for i := 0; i < len(urls); i++ {
		<-subEnd
	}

	end <- true
}

var crawled = make(map[string]bool)
var crawledMutex sync.Mutex

func pad(l int, c string) string {
	l = 4 - 1
	s := ""
	for i := 0; i < 1; i++ {
		s += c
	}
	return s
}

func Crawl2(url string, depth int, fetcher Fetcher, ch chan string, history map[string]bool) {
	var c func(url string, depth int, fetcher Fetcher, ch chan string, history map[string]bool)
	c = func(url string, depth int, fetcher Fetcher, ch chan string, history map[string]bool) {
		if depth <= 0 {
			return
		}
		body, urls, err := fetcher.Fetch(url)

		if err != nil {
			fmt.Println(err)
			return
		}

		history[url] = true

		ch <- fmt.Sprintf("%sfound %s %q\n **** %v\n", pad(depth, "-"), url, body, urls)
		for _, u := range urls {
			if !history[u] {
				history[u] = true
				c(u, depth-1, fetcher, ch, history)
			} else {
				fmt.Printf(">>>%s is repetitive\n", u)
			}
		}
		return
	}
	c(url, depth, fetcher, ch, history)
	close(ch)
}

func main() {
	fmt.Println("########### fake ############")
	fakeCrawl("http://golang.org/", 4, fetcher)

	fmt.Println("########## crawl2 ##########")
	ch := make(chan string)
	history := make(map[string]bool)
	go Crawl2("http://golang.org/", 4, fetcher, ch, history)
	for s := range ch {
		fmt.Println(s)
	}

	fmt.Println("============= history ==============")
	fmt.Println(history)

	fmt.Println("########### crawl1 ##########")
	out := make(chan string)
	end := make(chan bool)

	go Crawl1("http://golang.org/", 4, fetcher, out, end)
	for {
		select {
		case t := <-out:
			fmt.Print(t)
		case <-end:
			return
		}
	}
}

// fakeCrawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度
func fakeCrawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL
	// TODO: 不重复抓取页面
	// 下面并没有实现上面的两种情况
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		fakeCrawl(u, depth-1, fetcher)
	}
	return
}

// fakeFetcher 是返回若干结果的 Fetcher
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

// fetcher 是填充后的 fakeFetcher
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org",
			"http://golang.org/pkg/",
		},
	},
}
