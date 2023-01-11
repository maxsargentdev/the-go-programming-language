package json

import (
	"fmt"
	"sync"
)

var xkcdJsonUrl = "https://xkcd.com/COMIC_NUMBER/info.0.json"
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

// use goroutines to download all files quickly and in parallel
// then we need to use some kind of index.map to improve the search
// how about, comic number -> string[] of each word in the
func RunXKCDIndexGen() {

	var wg sync.WaitGroup
	for i := xkcdStartIndex; i <= xkcdFinalIndex; i++ {

		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			getRequestWorker(i)
		}()

	}

	wg.Wait()

}

func getRequestWorker(comicNumber int) {
	//xkcdURL := strings.Replace(xkcdJsonUrl, "COMIC_NUMBER", fmt.Sprintf("%d", i), 1) // just replcae with sprintf
	//fmt.Println(xkcdURL)
	//_, err := http.Get(xkcdURL)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	fmt.Println("Executed a task")
}

func RunXKCDIndexSearch() {
	fmt.Println("Searching XKCD index")
}
