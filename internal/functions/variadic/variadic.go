package variadic

import (
	"fmt"
	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
	"net/http"
)

func Max(integers ...int) {

	if len(integers) == 0 {
		fmt.Println("no input provided")
		return
	}

	max := integers[0]
	for _, v := range integers {
		if v > max {
			max = v
		}
	}
	fmt.Printf("%d\n", max)
}

func Min(integers ...int) {

	if len(integers) == 0 {
		fmt.Println("no input provided")
		return
	}

	min := integers[0]
	for _, v := range integers {
		if v < min {
			min = v
		}
	}
	fmt.Printf("%d\n", min)
}

func Join(sep string, strings ...string) string {

	if sep == "" {
		fmt.Println("no input provided")
		return ""
	}
	if len(strings) == 0 {
		fmt.Println("no input provided")
		return ""
	}

	var returnMe string
	for i, v := range strings {
		if i > 0 {
			returnMe = fmt.Sprintf("%s%s%s", returnMe, sep, v)
		} else {
			returnMe = fmt.Sprintf("%s%s", returnMe, v)
		}
	}
	fmt.Println(returnMe)
	return returnMe
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func RunGetElementsByTagName(url string, names ...string) {
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
	elementsByTagName(doc, names...)
}

func elementsByTagName(node *html.Node, names ...string) []*html.Node {

	returnMe := []*html.Node{}

	pre := func(n *html.Node) {
		if n.Type == html.ElementNode && slices.Contains(names, n.Data) {
			returnMe = append(returnMe, n)
		}
	}

	forEachNode(node, pre, nil)

	return returnMe
}
