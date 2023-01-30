package multireturn

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func RunCountWordsAndImages(url string) {

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)

	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images := countWordsAndImages(doc)

	fmt.Println(words)
	fmt.Println(images)
}

func countWordsAndImages(n *html.Node) (words, images int) {
	// content in <style> or <script> are ignored
	if n.Type == html.ElementNode {
		if n.Data == "style" || n.Data == "script" {
			return
		} else if n.Data == "img" {
			images++
		}
	} else if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		for _, line := range strings.Split(text, "\n") {
			if line != "" {
				words += len(strings.Split(line, " "))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return
}

func getAttribute(attributeKey string, n html.Node) (string, error) {
	for _, a := range n.Attr {
		if a.Key == attributeKey {
			return a.Val, nil
		}
	}
	return "", fmt.Errorf("no attribute found")
}
