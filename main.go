package main

import (
	"Crawler-url-go/lib"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("[+] Crawler all links in web page.")
		fmt.Println("[+] Usage : ", os.Args[0], "<url>")
		fmt.Println("[+] Example : ", os.Args[0], "https://www.localhost.com")
		os.Exit(1)
	}
	target := os.Args[1]
	crawler := &lib.WCrawler{}
	c := crawler.Crawler(target)
	for _, link := range c.Links {
		fmt.Println("[*] Found : ", link)
	}
}
