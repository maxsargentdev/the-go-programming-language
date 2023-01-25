package recursion

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func FindLinks() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {

	// If the node is a HTML element and is of the a tag
	if n.Type == html.ElementNode && n.Data == "a" {
		// for each attribute in the a tag
		for _, a := range n.Attr {
			// if tghe attribute is a href (link) then append
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	//for c := n.FirstChild; c != nil; c = c.NextSibling {
	//	links = visit(links, c)
	//}

	return links
}

func Outline() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func GenerateElementMap() {
	fmt.Println("Generate element map...")
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elementmap: %v\n", err)
		os.Exit(1)
	}

	elementMap := make(map[string]int)
	fillElementMap(&elementMap, doc)
	fmt.Println(elementMap)

}

func fillElementMap(elementMap *map[string]int, n *html.Node) {

	if n.Type == html.ElementNode {
		mapRef := *elementMap
		mapRef[n.Data]++
	}

	if n.FirstChild != nil {
		fillElementMap(elementMap, n.FirstChild)
	}

	if n.NextSibling != nil {
		fillElementMap(elementMap, n.NextSibling)
	}
}

func GenerateTextNodes() {
	fmt.Println("Generating text nodes..........")
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "textnodes: %v\n", err)
		os.Exit(1)
	}

	descendAndPrint(doc)
}

func descendAndPrint(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
		fmt.Println(n.)
	}

	if n.FirstChild != nil {
		descendAndPrint(n.FirstChild)
	}

	if n.NextSibling != nil {
		descendAndPrint(n.NextSibling)
	}
}
