# Scraping stock prices using Go and Colly


## Overview
This Go project performs the following tasks:
- Takes in a list of company ticker symbols (e.g. AMZN for Amazon.com, Inc.)
- Sends HTTP requests to retrieve HTML from target web pages
- Extracts data on stock prices by specifying HTML elements
- Writes out the retrieved data to a .csv file

## Install
To run this project, you will need to install [Go](https://go.dev/doc/install)  
Once you have Go installed on your local machine, run the following commands in the terminal:
```
git clone https://github.com/shosukenoma/stock-price-scraper.git
go get -u github.com/gocolly/colly/...
```

## Usage

Run the following command in the terminal:
```
go run main.go
```
Once the program finishes running, you will see a `stocks.csv` file generated in the same directory.  
<img src="/assets/images/csv_file_created.png" alt="Screenshot of the new .csv being created" width="200" style={padding: 10px 0;}/>

## Acknowledgements

This project is inspired by [Web Scraping in Golang: Complete Guide 2023](https://www.zenrows.com/blog/web-scraping-golang#how-to-web-scrape-in-go)  
This tutorial article walks us through how to scrape a mock e-commerce website that consists of a well-structured HTML.

## Other Resources
  
[Web Scraping for Stock Prices in Python](https://www.geeksforgeeks.org/web-scraping-for-stock-prices-in-python/#)  
This article is helpful in getting a general sense of the project structure.

[Building a web scraper in Go with Colly](https://blog.logrocket.com/building-web-scraper-go-colly/)  
This article accurately shows that `c.Visit()` should come *after* all of the callback functions, not before.

[Programming in Go for Web Scraping](https://medium.com/@shahidahmed.org/programming-in-go-for-web-scraping-aedf937e769d#:~:text=The%20most%20important%20callback%20for%20our%20need%20is%20OnHTML().)  
This article is helpful in understanding how to retrieve data from HTML elements in varying cases not mentioned in the Colly documentation.  

[Introduction to bufio package in Golang](https://medium.com/golangspec/introduction-to-bufio-package-in-golang-ad7d1877f762#:~:text=bufio.Writer%20sends%20data%20only%20when%20buffer%20is%20either%20full%20or%20when%20explicitly%20requested%20with%20Flush%20method.)  
This article is helpful in understanding the bufio package and the role of `defer writer.Flush()`.
