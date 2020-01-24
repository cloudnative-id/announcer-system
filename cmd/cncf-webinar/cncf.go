package main

import (
	"log"
	"fmt"
	"strings"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"github.com/cloudnative-id/announcer-system/models"
)

func GetContentCNCF(currentContent models.WebinarCNCFList) models.WebinarCNCFList {
	var newContent models.WebinarCNCFList
	res, err := http.Get("https://www.cncf.io/webinars/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".upcoming-webinars").Each(func(i int, s0 *goquery.Selection) {
		s0.Find("article").Not("a.button-like").Each(func(i int, s1 *goquery.Selection) {
			url, _ := s1.Find("a").Attr("href")
			title := strings.Replace(s1.Find("a").Text(),"Find Out More","",-1)
			fmt.Println(title)
			doAppend := true

			for _, v := range currentContent.Content {
				if (v.Url == url){
					doAppend = false
				}
			}

			if doAppend {
				month := s1.Find(".upcoming-date.upcoming-date-mobile").ChildrenFiltered(".month").Text()
				day := s1.Find(".upcoming-date.upcoming-date-mobile").ChildrenFiltered(".day").Text()
				year := s1.Find(".upcoming-date.upcoming-date-mobile").ChildrenFiltered(".year").Text()
				
				date := month+" "+day+" "+year
				time := s1.Find(".details").ChildrenFiltered(".time").Text()
				
				singleContent := models.WebinarCNCFContent{Title: title, Url: url, Date: date, Time: time, IsDelivered: false}
				newContent.Content = append(newContent.Content,singleContent)
			}
		})
	})

	return newContent
}

