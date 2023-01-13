package json

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

var xkcdJsonUrl = "https://xkcd.com/%d/info.0.json"
var xkcdStartIndex = 1
var xkcdFinalIndex = 2723
var indexFileLocation = "/tmp/xkcd-index.json"

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

type IndexedComic struct {
	Comic      Comic
	TitleIndex TitleWordList
}

var xkcdIndex = make(map[int]IndexedComic)

// use goroutines to download all files quickly and in parallel
// then we need to use some kind of index.map to improve the search
// how about, comic number -> string[] of each word in the
func RunXKCDIndexGen() {

	var mu sync.Mutex
	var wg sync.WaitGroup

	// Add step here to check existence of index file, unmarshal and load into index first
	// Then adjust start index for worker group so we just cache new stuff....
	// We will also need a REST request to fetch the latest ID from xkcd.com

	for i := xkcdStartIndex; i <= xkcdFinalIndex; i++ {

		wg.Add(1)

		i := i // variable is shadowing

		go func() {
			mu.Lock()
			defer wg.Done()
			defer mu.Unlock()
			comic := getXKCDWorker(i)
			titleList := splitTitle(comic.Title)
			sort.Strings(titleList)
			xkcdIndex[i] = IndexedComic{Comic: comic, TitleIndex: titleList}
		}()

	}

	wg.Wait()
}

func splitTitle(title string) []string {
	return strings.Split(title, " ")
}

func getXKCDWorker(comicNumber int) Comic {
	//xkcdURL := fmt.Sprintf(xkcdJsonUrl, comicNumber)
	//_, err := http.Get(xkcdURL)
	//if err != nil {
	//	log.Fatalln(err)
	//} Update to return JSON later and unmarshal, for now just return a string and focus on building the index

	return Comic{Title: "Test Title Has Has A Few Words"}
}

func RunXKCDIndexSearch() {
	fmt.Println("Searching XKCD index")
}

// Bonus - generate the index and save to disk as a cache, for working offline on the train
func RunXKCDMaterialize() {
	serializedXkcdIndex, _ := json.Marshal(xkcdIndex)

	xkcdIndexFile := createFile(indexFileLocation)
	defer closeFile(xkcdIndexFile)

	_, err := fmt.Fprintln(xkcdIndexFile, string(serializedXkcdIndex))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error updating issue: %v\n", err)
		os.Exit(1)
	}
}
