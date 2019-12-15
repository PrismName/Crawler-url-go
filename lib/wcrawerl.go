package lib

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

// EXTENSION : white list
var EXTENSION = []string{".js", ".css", ".zip", ".jpg"}

// WCrawler : struct
type WCrawler struct {
	Links []string
}

// Crawler : extractor web page link
func (w *WCrawler) Crawler(targetURL string) {
	resp, err := WRequest(targetURL)
	if err != nil {
		fmt.Errorf("response content is None %s", err)
	}

	docs, _ := goquery.NewDocumentFromResponse(resp)
	w.ExtractorLink(docs)
}

// ExtractorLink : extractor all link from web page
func (w *WCrawler) ExtractorLink(docs *goquery.Document) {
	docs.Find("a").Each(func(i int, selection *goquery.Selection) {
		//fmt.Println(i)
		aLink, _ := selection.Attr("href")
		fmt.Println(aLink)
		w.Links = append(w.Links, aLink)
	})

	docs.Find("img").Each(func(i int, selection *goquery.Selection) {
		src, _ := selection.Attr("src")
		fmt.Println(src)
		w.Links = append(w.Links, src)
	})
}

func checkLinkExtension(link string) string {
	for _, v := range EXTENSION {
		if v == link {
			return link
		}
	}
	return ""
}

func checkLink(link string) string {
	return ""
}
