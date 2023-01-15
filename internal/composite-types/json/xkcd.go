package json

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
)

var XKCDComicJsonURL = "https://xkcd.com/%d/info.0.json"
var latestXKCDComicJsonURL = "https://xkcd.com/info.0.json"
var indexFileLocation = "/tmp/xkcd-index.json"
var xkcdIndex = make([]IndexedComic, 0)

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

func RunXKCDIndexGen() {
	resp, err := http.Get(latestXKCDComicJsonURL)
	if err != nil {
		log.Fatal("error getting latest comic: ", err)
	}
	comic := Comic{}
	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &comic)
	if err != nil {
		log.Fatal("error during unmarshal of latest comic: ", err)
	}

	XKCDStartIndex := 1
	XKCDFinalIndex := comic.Num

	data, err := os.ReadFile(indexFileLocation)
	if os.IsNotExist(err) {
		log.Println("no existing index file detected")
	} else if err != nil {
		log.Fatal("fatal error during file read: ", err)
	} else {
		err = json.Unmarshal(data, &xkcdIndex)
		if err != nil {
			log.Fatal("error during unmarshal of existing index file: ", err)
		}
		XKCDStartIndex = xkcdIndex[len(xkcdIndex)-1].Comic.Num + 1
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := XKCDStartIndex; i <= XKCDFinalIndex; i++ {

		wg.Add(1)

		i := i // variable is shadowing

		go func() {
			mu.Lock()

			defer wg.Done()
			defer mu.Unlock()

			comic := getXKCDWorker(i)

			titleList := splitTitle(comic.Title)
			sort.Strings(titleList)

			xkcdIndex = append(xkcdIndex, IndexedComic{Comic: comic, TitleIndex: titleList})

			fmt.Printf("downloaded: %s (%d) \n", comic.Title, i)
		}()

	}

	wg.Wait()

	// Sort the slice so when we materialize it is easy to search and to determine our current position
	sort.Slice(xkcdIndex, func(i, j int) bool {
		return xkcdIndex[i].Comic.Num < xkcdIndex[j].Comic.Num
	})
}

func splitTitle(title string) []string {
	return strings.Split(title, " ")
}

// getXKCDWorker is a function designed to be called as a go routine, it fetches and unmarshalls comics.
func getXKCDWorker(comicNumber int) Comic {
	XKCDURL := fmt.Sprintf(XKCDComicJsonURL, comicNumber)

	resp, err := http.Get(XKCDURL)
	if err != nil {
		log.Fatalln(err)
	}

	comic := Comic{}

	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &comic)
	if err != nil {
		log.Println("error thrown in a worker process")
	}
	return comic
}

// RunXKCDIndexSearch generates, materializes and searches the the XKCD web-comic archives.
func RunXKCDIndexSearch() {
	fmt.Println("Searching XKCD index")
}

// RunXKCDMaterialize takes the current global index state and writes it to list
func RunXKCDMaterialize() {
	serializedXkcdIndex, _ := json.Marshal(xkcdIndex)
	xkcdIndexFile := createFile(indexFileLocation)
	defer closeFile(xkcdIndexFile)
	_, err := fmt.Fprintln(xkcdIndexFile, string(serializedXkcdIndex))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error materializing index: %v\n", err)
		os.Exit(1)
	}

}
