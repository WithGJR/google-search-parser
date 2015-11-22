package google_search_parser

import (
	"github.com/PuerkitoBio/goquery"
	crawler "github.com/WithGJR/search-result-crawler"
	"strconv"
	"strings"
)

type GoogleSearchParser struct{}

func (p *GoogleSearchParser) GetSearchResultPageURL(keyword string, page int) string {
	keyword = strings.Join(strings.Split(keyword, " "), "+")
	url := "https://www.google.com.tw/search?q=" + keyword
	if page == 1 {
		return url
	} else {
		return url + "&start=" + strconv.Itoa((page-1)*10)
	}
}

func (p *GoogleSearchParser) Parse(doc *goquery.Document, keyword string, page int, channel chan crawler.IntermediatePair) {
	doc.Find(".rc").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".r > a")
		url, _ := title.Attr("href")
		description := s.Find(".s > div > .st").Text()
		result := crawler.Result{Title: title.Text(), URL: url, Description: description}
		channel <- crawler.IntermediatePair{
			Keyword: keyword,
			Index:   i,
			Page:    page - 1,
			Result:  result,
		}
	})
}
