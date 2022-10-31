package tutorial

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Fetch2() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = fmt.Sprintf("http://%s", url)
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
	}
}
