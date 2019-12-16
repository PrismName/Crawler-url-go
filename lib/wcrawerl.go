package lib

import (
	"fmt"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

// EXTENSION : white list
var EXTENSION = []string{".js", ".css", ".zip", ".jpg"}

// WCrawler : struct
type WCrawler struct {
	Links []string
}

// Crawler : extractor web page link
func (w *WCrawler) Crawler(targetURL string) *WCrawler {
	resp, err := WRequest(targetURL)
	if err != nil {
		fmt.Errorf("response content is None %s", err)
	}

	docs, _ := goquery.NewDocumentFromResponse(resp)
	w.ExtractorLink(docs)
	return w
}

// ExtractorLink : extractor all link from web page
func (w *WCrawler) ExtractorLink(docs *goquery.Document) {
	docs.Find("a").Each(func (i int, selection *goquery.Selection) {
		aLink, _ := selection.Attr("href")
		if isentire(aLink) {
			w.Links = append(w.Links, aLink)
		}
	})

	docs.Find("img").Each(func (i int, selection *goquery.Selection) {
		src, _ := selection.Attr("src")
		if isentire(src) {
			w.Links = append(w.Links, src)
		}
	})

	docs.Find("script").Each(func (i int, selection *goquery.Selection)  {
		link, _ := selection.Attr("src")
		if isentire(link) {
			w.Links = append(w.Links, link)
		}
	})
}

func isentire(targetURL string) bool {
	if strings.HasPrefix(targetURL, "http://") || strings.HasPrefix(targetURL, "https://") {
		return true
	} else if strings.HasPrefix(targetURL, "/") || strings.HasPrefix(targetURL, "//") || strings.HasPrefix(targetURL, "javascript;:") {
		return false
	}
	return false
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
