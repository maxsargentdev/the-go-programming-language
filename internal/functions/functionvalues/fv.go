package functionvalues

import (
	"fmt"
	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"sort"
	"strings"
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

func forEachNode(n *html.Node, pre, post func(n *html.Node) (b bool)) {
	if pre != nil {
		stop := pre(n)
		if stop {
			return
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		stop := post(n)
		if stop {
			return
		}
	}
}

var depth int

func startElement(n *html.Node) (b bool) {

	// Element node with attributes
	if n.Type == html.ElementNode && n.Attr != nil {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, attr := range n.Attr {
			fmt.Printf(" %s=\"%s\"", attr.Key, attr.Val)
		}
		fmt.Printf(">\n")
		depth++
		return false
	}

	// Comment nodes, make sure they get indented correctly with sprintf hack
	if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s", depth*2, "", strings.ReplaceAll(n.Data, "\n", fmt.Sprintf("\n%*s", depth*2, "")))
		fmt.Printf("-->\n")
		return false
	}

	if n.Type == html.TextNode && (n.Parent.Data == "script" || n.Parent.Data == "style") {
		for _, line := range strings.Split(n.Data, "\n") {
			line = strings.TrimSpace(line)
			if line != "" && line != "\n" {
				fmt.Printf("%*s%s\n", depth*2, "", line)
			}
		}
		return false
	}

	// Standard element node
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}

	return false
}

func endElement(n *html.Node) (b bool) {

	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
	return false
}

// c5e7 Still todo:
//- Print text nodes
//- Remove trailing tag for elements with no children
//- Write a test

func RunGetElementByID(url string, id string) {
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

	ElementByID(doc, id)
}

func ElementByID(doc *html.Node, id string) *html.Node {

	pre := func(n *html.Node) (b bool) {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return true
			}
		}
		return false
	}

	forEachNode(doc, pre, nil)

	tmp := html.Node{}
	return &tmp
}

func RunExpand(s string) {
	f := func(s string) string {
		return strings.ToUpper(s)
	}
	expand(s, f)
}

func expand(s string, f func(string) string) (result string) {
	items := strings.Split(s, " ")
	for i, item := range items {
		if strings.HasPrefix(item, "$") {
			items[i] = f(item[1:])
		}
	}

	result = strings.Join(items, " ")
	fmt.Println(result)
	return
}

func RunTopoSort(prereqs map[string][]string) {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func RunTopoSortMap(m map[string]map[string]bool) {
	for i, course := range topoSortMap(m) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSortMap(m map[string]map[string]bool) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for item, _ := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys = make(map[string]bool)
	for key := range m {
		keys[key] = true
	}
	visitAll(keys)
	return order
}

func RunTopoSortWithCycleDetection(prereqs map[string][]string) {
	for i, course := range topoSortWithCycleDetection(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSortWithCycleDetection(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {

			for _, requirement := range m[item] {
				if slices.Contains(m[requirement], item) {
					fmt.Printf("Cycle detected: %s <-> %s\n", item, requirement)
					os.Exit(1)
				}
			}

			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
