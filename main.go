package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type PokemonProduct struct{
	url, image, name, price string
}

func main() {
	// Initialize a slice to store structs for each PokemonProduct
	pokemonProducts := []PokemonProduct{}

	// A Collector allows you to perform HTTP requests.
	// It also gives you access to the web scraping callbacks offered by the Colly interface.
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
})

	// iterating over the list of HTML product elements
	// e represents a single li.product element
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		// initializing a new PokemonProduct instance
		pokemonProduct := PokemonProduct{}

		// scraping data of interest
		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2") 
		pokemonProduct.price = e.ChildText(".price") 

		// add product instance to the list of products
		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})

	// c.Visit() has to come after all the callback functions
	c.Visit("https://scrapeme.live/shop/")

	fmt.Println(pokemonProducts)
	

	// opening the CSV file 
	file, err := os.Create("products.csv") 
	if err != nil { 
		log.Fatalln("Failed to create output CSV file", err) 
	} 
	defer file.Close() 

	// initializing a file writer 
	writer := csv.NewWriter(file) 
	
	// defining the CSV headers 
	headers := []string{ 
		"url", 
		"image", 
		"name", 
		"price", 
	} 
	// writing the column headers 
	writer.Write(headers) 

	// adding each Pokemon product to the CSV output file 
	for _, pokemonProduct := range pokemonProducts { 
		// converting a PokemonProduct to an array of strings 
		record := []string{ 
			pokemonProduct.url, 
			pokemonProduct.image, 
			pokemonProduct.name, 
			pokemonProduct.price, 
		} 
	
		// writing a new CSV record 
		writer.Write(record) 
	} 
	// writer.Flush() ensures that any remaining data is written.
	// bufio.Writer sends data only when buffer is either full or when explicitly requested with Flush method.
	// https://medium.com/golangspec/introduction-to-bufio-package-in-golang-ad7d1877f762#:~:text=bufio.Writer%20sends%20data%20only%20when%20buffer%20is%20either%20full%20or%20when%20explicitly%20requested%20with%20Flush%20method.
	defer writer.Flush()
}

