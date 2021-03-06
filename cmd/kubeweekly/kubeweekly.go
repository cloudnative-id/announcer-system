package main

import (
	"time"
	"regexp"
	"strings"
    "github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"github.com/cloudnative-id/announcer-system/models"
)

func GetDate() string {
	dt := time.Now()
	Date := dt.Format("01-02-2006")
	Date = strings.ReplaceAll(Date, "-", "/")

	return Date
}

func GetNewKubeweeklyTitle() string {
	Res, err := http.Get("https://kubeweekly.io/")
	if err != nil {
		log.Fatal(err)
	}
	defer Res.Body.Close()

	Doc, err := goquery.NewDocumentFromReader(Res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var Title string

	Doc.Find("td#templateHeader").ChildrenFiltered(".mcnTextBlock").Each(func(i int, s0 *goquery.Selection) {
		Text := strings.ToLower(s0.Find("strong").Text())
		Reg, _ := regexp.Compile(`kubeweekly(.*)#\d+`)
		Title = Reg.FindString(Text)
	})

	return Title

}
func GetContentListKubeweekly() models.KubeweeklyData {
	var Data models.KubeweeklyData

	Data.Title = strings.ReplaceAll(GetNewKubeweeklyTitle(), " #", "")
	Data.Date = GetDate()
	Data.ContentFile = "contents/"+Data.Title+".yaml"
	Data.Status.IsDelivered = false
	Data.Tags = append(Data.Tags,"#kubereads")

	return Data
}

func GetContentKubeweekly() models.KubeweeklyContent{

	var Content models.KubeweeklyContent

	Content.Title = GetNewKubeweeklyTitle()
	Content.Source = "kubeweekly"
	Content.Date = GetDate()

	Res, err := http.Get("https://kubeweekly.io/")
	if err != nil {
		log.Fatal(err)
	}
	defer Res.Body.Close()

	Doc, err := goquery.NewDocumentFromReader(Res.Body)
	if err != nil {
		log.Fatal(err)
	}

	Doc.Find("td#templateBody").ChildrenFiltered(".mcnTextBlock").Each(func(i int, s0 *goquery.Selection) {
		var Type string
		var Title string
		var Link string

		Type = s0.Find("h1").Text()

		s0.Find("a").Each(func(i int, s1 *goquery.Selection) {
			Title = s1.Text()
			Link, _ = s1.Attr("href")

			ContentData := models.KubeweeklyContentData{Title: Title, Type: Type, Link: Link}
			Content.Data = append(Content.Data,ContentData)
		})
	})

	return Content
}

