package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("[+] Crawler all links in web page.")
		fmt.Println("[+] Usage : ", os.Args[0], "<url>")
		fmt.Println("[+] Example : ", os.Args[0], "https://www.localhost.com")
		os.Exit(1)
	}

	url := os.Args[1]

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal("[-] Created request object error : ", err)
	}

	request.Header.Add("User-Agent", "__spider_man__")

	response, err := client.Do(request)

	if err != nil {
		log.Fatal("[-] load response content error : ", err)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Fatal("[-] Can't parser response document : ", err)
	}

	doc.Find("a").Each(func(i int, selector *goquery.Selection) {
		href, err := selector.Attr("href")
		if err {
			fmt.Println("[*] Found url : ", href)
		}
	})
}
