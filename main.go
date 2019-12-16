package main

import (
	"fmt"
	"os"
	"Crawler-url-go/lib"
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

/*
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("[+] Crawler all links in web page.")
		fmt.Println("[+] Usage : ", os.Args[0], "<url>")
		fmt.Println("[+] Example : ", os.Args[0], "https://www.localhost.com")
		os.Exit(1)
	}

	targetUrl := os.Args[1]

	client := &http.Client{}

	request, err := http.NewRequest("GET", targetUrl, nil)

	if err != nil {
		log.Fatal("[-] Created request object error : ", err)
	}

	request.Header.Add("User-Agent", "__spider_man__")

	response, err := client.Do(request)

	urlsArr := []string{}

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

			if strings.HasPrefix(href, "/") || strings.HasPrefix(href, "//") || strings.HasPrefix(href, "javascript:;") {
				fmt.Println("[*] the url is invalid.", href)
			} else {
				//fmt.Println("[+] Found url : ", href)
				urlsArr = append(urlsArr, href)
			}
		}
	})

	for k, v := range urlsArr {
		fmt.Println(k, v)
	}
}
*/
