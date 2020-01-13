package lib

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
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
	w.ExtractorLink(docs, resp)
	return w
}

// ExtractorLink : extractor all link from web page
func (w *WCrawler) ExtractorLink(docs *goquery.Document, resp *http.Response) {
	docs.Find("a").Each(func(i int, selection *goquery.Selection) {
		aLink, _ := selection.Attr("href")
		ajUrl := joinUrl(aLink, resp)
		w.Links = append(w.Links, ajUrl)
	})

	docs.Find("img").Each(func(i int, selection *goquery.Selection) {
		src, _ := selection.Attr("src")

		sjUrl := joinUrl(src, resp)
		w.Links = append(w.Links, sjUrl)
	})

	docs.Find("script").Each(func(i int, selection *goquery.Selection) {
		link, _ := selection.Attr("src")
		jUrl := joinUrl(link, resp)
		w.Links = append(w.Links, jUrl)
	})
}

func joinUrl(targetURL string, response *http.Response) string {
	switch {
	case strings.HasPrefix(targetURL, "//"):
		return "http:" + targetURL
	case strings.HasPrefix(targetURL, "#") ||
		strings.HasPrefix(targetURL, "javascript:;") ||
		strings.HasPrefix(targetURL, "/"):
		return response.Request.URL.String()
	case strings.HasPrefix(targetURL, "http://") || strings.HasPrefix(targetURL, "https://"):
		return targetURL
	default:
		return targetURL
	}
}

func checkLinkExtension(link string) string {
	for _, v := range EXTENSION {
		if v == link {
			return link
		}
	}
	return ""
}
