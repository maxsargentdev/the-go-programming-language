package tutorial

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func FetchAll2() {
	start := time.Now()
	ch := make(chan string)

	for _, urlstring := range os.Args[1:] {
		go fetch2(urlstring, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch2(urlstring string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(urlstring)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	url, _ := url.Parse(urlstring)
	hostname := strings.TrimPrefix(url.Hostname(), "www.")

	file, _ := os.Create(hostname)
	nbytes, err := io.Copy(file, resp.Body)

	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", urlstring, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, urlstring)
}
