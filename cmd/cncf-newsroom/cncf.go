package main

import (
	"fmt"
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

			title := s1.Text()

			for _,v := range currentContent.content {
				if v.title == title {
					url, _ := s1.Attr("href")

					singleContent := ContentCNCFList{title: title, url: url, kind: kind, isDelivered: false}
					newContent.content = append(newContent.content,singleContent)
				}
			}
		})
	})

	fmt.Println(newContent)
	return newContent
}

