package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Stock struct{
	company, price, change, volume string
}

func main() {

	// urls := []string{
	// 	"https://www.investing.com/equities/nike",
  //   "https://www.investing.com/equities/coca-cola-co",
  //   "https://www.investing.com/equities/microsoft-corp",
  //   "https://www.investing.com/equities/3m-co",
  //   "https://www.investing.com/equities/american-express",
  //   "https://www.investing.com/equities/amgen-inc",
  //   "https://www.investing.com/equities/apple-computer-inc",
  //   "https://www.investing.com/equities/boeing-co",
  //   "https://www.investing.com/equities/cisco-sys-inc",
  //   "https://www.investing.com/equities/goldman-sachs-group",
  //   "https://www.investing.com/equities/ibm",
  //   "https://www.investing.com/equities/intel-corp",
  //   "https://www.investing.com/equities/jp-morgan-chase",
  //   "https://www.investing.com/equities/mcdonalds",
  //   "https://www.investing.com/equities/salesforce-com",
  //   "https://www.investing.com/equities/verizon-communications",
  //   "https://www.investing.com/equities/visa-inc",
  //   "https://www.investing.com/equities/wal-mart-stores",
  //   "https://www.investing.com/equities/disney",
	// }

	// Initialize a slice to store structs for each Stock
	stocks := []Stock{}

	// A Collector allows you to perform HTTP requests.
	// It also gives you access to the web scraping callbacks offered by the Colly interface.
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	// c.OnError(func(_ *colly.Response, err error) { 
	// 	log.Println("Something went wrong: ", err) 
	// }) 
	 
	// c.OnResponse(func(r *colly.Response) { 
	// 	fmt.Println("Page visited: ", r.Request.URL) 
	// }) 

	// iterating over the list of HTML product elements
	// e represents a single li.product element
	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {
		// initializing a new PokemonProduct instance
		stock := Stock{}

		// scraping data of interest
		stock.company = e.ChildText("h1")
		fmt.Println("Company:", stock.company)
		stock.price = e.ChildText("fin-streamer[data-field='regularMarketPrice']")
		fmt.Println("Price:", stock.price)
		stock.change = e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")
		fmt.Println("Change:", stock.change)

		// https://medium.com/@shahidahmed.org/programming-in-go-for-web-scraping-aedf937e769d#:~:text=Here%20is%20the%20second%20part%20of%20our%20complete%20program.

		// add product instance to the list of products
		stocks = append(stocks, stock)
	})

	// for _, stock := range urls {
		// c.Visit() has to come after all the callback functions
		c.Visit("https://finance.yahoo.com/quote/NKE/")
	// }

	fmt.Println(stocks)
	

	// // opening the CSV file 
	// file, err := os.Create("products.csv") 
	// if err != nil { 
	// 	log.Fatalln("Failed to create output CSV file", err) 
	// } 
	// defer file.Close() 

	// // initializing a file writer 
	// writer := csv.NewWriter(file) 
	
	// // defining the CSV headers 
	// headers := []string{ 
	// 	"url", 
	// 	"image", 
	// 	"name", 
	// 	"price", 
	// } 
	// // writing the column headers 
	// writer.Write(headers) 

	// // adding each Pokemon product to the CSV output file 
	// for _, pokemonProduct := range pokemonProducts { 
	// 	// converting a PokemonProduct to an array of strings 
	// 	record := []string{ 
	// 		pokemonProduct.url, 
	// 		pokemonProduct.image, 
	// 		pokemonProduct.name, 
	// 		pokemonProduct.price, 
	// 	} 
	
	// 	// writing a new CSV record 
	// 	writer.Write(record) 
	// } 
	// // writer.Flush() ensures that any remaining data is written.
	// // bufio.Writer sends data only when buffer is either full or when explicitly requested with Flush method.
	// // https://medium.com/golangspec/introduction-to-bufio-package-in-golang-ad7d1877f762#:~:text=bufio.Writer%20sends%20data%20only%20when%20buffer%20is%20either%20full%20or%20when%20explicitly%20requested%20with%20Flush%20method.
	// defer writer.Flush()
}

