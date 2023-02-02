package functionvalues

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func RunHTMLPrettyPrint(url string) {

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

	forEachNode(doc, startElement, endElement)

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

var depth int

func startElement(n *html.Node) {

	if n.Type == html.ElementNode && n.Attr != nil {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, attr := range n.Attr {
			fmt.Printf(" %s=\"%s\"", attr.Key, attr.Val)
		}
		fmt.Printf(">\n")
		depth++
		return
	}

	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}

	//if n.Type == html.TextNode {
	//	fmt.Printf("%s", strings.ReplaceAll(n.Data, "\n", ""))
	//}
}

func endElement(n *html.Node) {

	//if n.Type == html.ElementNode && n.PrevSibling.Type == html.TextNode {
	//	fmt.Printf("%*s</%s>", depth*2, "", n.Data)
	//	return
	//}

	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

// Still todo:
//
//- Print comments
//- Print text nodes
//- Remove trailing tag for elements with no children
//- Write a test
