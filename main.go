package main

import (
	"fmt"

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
	
}

// https://www.zenrows.com/blog/web-scraping-golang#set-up-go-project