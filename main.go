package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type Info struct {
	Title         string
	Author        string
	JournalTitle  string
	SubmissionDay string
}

const url = "https://scholar.google.com.vn/citations?hl=en&user=75ul2BYAAAAJ"

func oneSR(pathUrl string) (*Info, error) {
	doc, err := goquery.NewDocument(pathUrl)
	if err != nil {
		return nil, err
	}
	link, ok := doc.Find("a.gsc_a_at").Attr("data-href")
	if ok {
		log.Println(link)
	}
	// infoDoc, err := goquery.NewDocument(link)
	return nil, err
}

// func onePage() ([]Info, error) {
// 	doc, err := goquery.NewDocument(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	infoList := make([]Info, 0)
// 	doc.Find("tr.gsc_a_tr").Each(func(index int, trHtml *goquery.Selection) {
// 		var info Info
// 		info.Title = trHtml.Find("a.gsc_a_at").Text()
// 		row := make([]string, 0)
// 		trHtml.Find("div").Each(func(index int, rowHtml *goquery.Selection) {
// 			row = append(row, rowHtml.Text())
// 		})
// 		info.Author = row[0]
// 		info.JournalTitle = row[1]
// 		infoList = append(infoList, info)
// 	})
// 	return infoList, nil
// }

func onePage() ([]Info, error) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	list := []Info{}

	c.OnHTML("td.gsc_a_t div", func(e *colly.HTMLElement) {
		info := Info{}
		info.Author = e.Text
		list = append(list, info)
	})

	c.Visit(url)
	return list, nil
}

func main() {
	list, err := onePage()
	if err != nil {
		log.Println(err)
	}
	b, _ := json.Marshal(list)
	// file, _ := os.Create("test.txt")
	// fmt.Fprintln(file, string(b))
	log.Println(string(b))
}
