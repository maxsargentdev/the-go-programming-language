package deferred

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func RunFetch(url string) {
	filename, bytesWritten, err := fetch(url)

	fmt.Println(filename)
	fmt.Println(bytesWritten)
	fmt.Println(err)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)

	defer func() {
		// Close file, but prefer error from Copy, if any.
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	return local, n, err
}
