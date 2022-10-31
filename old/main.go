package main

import ( //tutorial "github.com/maxsargentdev/the-go-programming-language/1-tutorial"
	//"fmt"
	//mandelbrot "github.com/maxsargentdev/the-go-programming-language/3-basic-data-types/mandlebrot"
	// comma "github.com/maxsargentdev/the-go-programming-language/3-basic-data-types/comma"
	// anagram "github.com/maxsargentdev/the-go-programming-language/3-basic-data-types/anagram"
	//iota "github.com/maxsargentdev/the-go-programming-language/3-basic-data-types/iota"
	// arrays "github.com/maxsargentdev/the-go-programming-language/4-composite-datatypes/arrays"
	"flag"

	tutorial "github.com/maxsargentdev/the-go-programming-language/1-tutorial"
	"github.com/maxsargentdev/the-go-programming-language/4-composite-types/arrays"
)

var exerciseSelector string

func init() {
	flag.StringVar(&exerciseSelector, "ex", "1.1", "Which exercise to demonstrate, --ex")
}

func main() {
	flag.Parse()

	switch exerciseSelector {
	case "1.1":
		tutorial.Echo1()
	case "4.2":
		arrays.SHAzam()
	}
	//tutorial.Echo1()
	//tutorial.Echo2()
	//tutorial.Echo3()
	//tutorial.Dup2()
	//tutorial.Lissajous1(os.Stdout)
	//tutorial.Lissajous2(os.Stdout)
	//tutorial.Fetch1()
	//tutorial.Fetch2()
	//tutorial.Fetch3()
	//tutorial.FetchAll1()
	//tutorial.FetchAll2()
	//tutorial.Serve()
	//programstructure.RunMe()
	// fmt.Println(popcount.PopCount(1000100001))
	// fmt.Println(popcount.PopCount2(1000100001))
	// fmt.Println(popcount.PopCount3(1000100001))
	// fmt.Println(popcount.PopCount4(1000100001))
	//basicdatatypes.Surface()
	//abasicdatatypes.Surface2()
	//basicdatatypes.Surface3()
	//basicdatatypes.Serve()
	//mandelbrot.RenderAll()
	// comma.CommaFloatingP(("12.12345"))
	// comma.CommaFloatingP(("-12"))
	// comma.CommaFloatingP("-123.1")
	// comma.CommaFloatingP("-12345.123")
	// anagram.Anagram("maxrocksa", "rocksmaxa")
}
