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
	if n.Type == html.ElementNode &&
		(n.Data == "style" || n.Data == "script") {
		return
	}

	// the above works because if the first element in a subtree is script (or style), its entire contents can be omitted
	// because we return before any recursive call takes place we do not traverse that part of the HTML doc

	if n.Type == html.TextNode {
		fmt.Printf("%s\n", n.Data)
	}

	if n.FirstChild != nil {
		descendAndPrint(n.FirstChild)
	}

	if n.NextSibling != nil {
		descendAndPrint(n.NextSibling)
	}
}

func GenerateExtendedVisit() {
	fmt.Println("Generate extended visit...")
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "extendedvisit: %v\n", err)
		os.Exit(1)
	}

	hrefs, imgs, styles, scripts := extendedVisit(doc)

	fmt.Println(hrefs)
	fmt.Println(imgs)
	fmt.Println(styles)
	fmt.Println(scripts)

}

// extendedVisit appends hrefs, imgs, style & script elements
func extendedVisit(n *html.Node) (hrefs, imgs, styles, scripts []string) {

	isHref := n.Type == html.ElementNode && n.Data == "a"
	//isScript := n.Type == html.ElementNode && n.Data == "script"
	isImage := n.Type == html.ElementNode && n.Data == "img"
	//isStyle := n.Type == html.ElementNode && n.Data == "link"

	switch {
	case isHref:
		attr, _ := getAttribute("href", *n)
		hrefs = append(hrefs, attr)

	//case isScript:
	//	attr, _ := getAttribute("src", *n)
	//	hrefs = append(hrefs, attr)

	case isImage:
		attr, _ := getAttribute("src", *n)
		hrefs = append(hrefs, attr)

		//case isStyle:
		//	fmt.Println("style found")
		//
	
	}

	if n.FirstChild != nil {
		hrefs, imgs, styles, scripts = extendedVisit(n.FirstChild)
	}
	if n.NextSibling != nil {
		hrefs, imgs, styles, scripts = extendedVisit(n.NextSibling)
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
