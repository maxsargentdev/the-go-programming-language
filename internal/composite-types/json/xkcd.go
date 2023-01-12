package json

import (
	"fmt"
	"strings"
	"sync"
)

var xkcdJsonUrl = "https://xkcd.com/%d/info.0.json"
var xkcdStartIndex = 1
var xkcdFinalIndex = 2723

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

type TitleWordList []string

// use goroutines to download all files quickly and in parallel
// then we need to use some kind of index.map to improve the search
// how about, comic number -> string[] of each word in the
func RunXKCDIndexGen() {

	var xkcdIndex = make(map[int]TitleWordList)

	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := xkcdStartIndex; i <= xkcdFinalIndex; i++ {

		wg.Add(1)

		i := i // variable is shadowing

		go func() {
			mu.Lock()
			defer wg.Done()
			defer mu.Unlock()
			xkcdIndex[i] = strings.Split(getXKCDWorker(i).Title, " ")
		}()

	}

	wg.Wait()
	fmt.Println(xkcdIndex[1][3])
}

func splitTitle(string title) string {
	return strings.Split(title, " ")
}

func getXKCDWorker(comicNumber int) Comic {
	xkcdURL := fmt.Sprintf(xkcdJsonUrl, comicNumber)
	fmt.Println(xkcdURL)

	//_, err := http.Get(xkcdURL)
	//if err != nil {
	//	log.Fatalln(err)
	//} Update to return JSON later and unmarshal, for now just return a string and focus on building the index

	return Comic{Title: "Test Title Has Has A Few Words"}
}

func RunXKCDIndexSearch() {
	fmt.Println("Searching XKCD index")
}

// Bonus - generate the index and save to disk as a cache, then when generating update cache
func RunXKCDMaterialize() {
	fmt.Println("Materializing XKCD index to disk")
}
