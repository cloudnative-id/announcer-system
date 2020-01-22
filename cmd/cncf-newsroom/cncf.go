package main

import (
    "github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func GetContentCNCF(url, kind string, currentContent ContentCNCF) ContentCNCF {
	
	var newContent ContentCNCF
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".post-list-article").Each(func(i int, s0 *goquery.Selection) {
		s0.Find("h2").ChildrenFiltered("a").Each(func(i int, s1 *goquery.Selection) {

			var doAppend bool
			doAppend = true

			title := s1.Text()
			url, _ := s1.Attr("href")

			for _, v := range currentContent.Content {
				if v.Url == url{
					doAppend = false
				}
			}

			if doAppend {
				singleContent := ContentCNCFList{Title: title, Url: url, Kind: kind, IsDelivered: false}
				newContent.Content = append(newContent.Content,singleContent)
			}

		})
	})

	return newContent
}

