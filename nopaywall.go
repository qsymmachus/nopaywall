package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/prometheus/common/log"
)

const DefaultURL = "https://en.wikipedia.org/wiki/%22Hello,_World!%22_program"

func LoadPage(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected status code: %s", response.Status)
	}

	return response, nil
}

func ParseParagraphs(response *http.Response) (string, error) {
	defer response.Body.Close()
	var paragraphs []string

	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return "", err
	}

	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		paragraphs = append(paragraphs, s.Text())
	})

	return strings.Join(paragraphs, "\n\n"), nil
}

func main() {
	url := flag.String("url", DefaultURL, "URL of the page you'd like to read")
	flag.Parse()

	tokenizer, err := LoadPage(*url)
	if err != nil {
		log.Error(err)
	}

	paragraphs, err := ParseParagraphs(tokenizer)
	if err != nil {
		log.Error(err)
	}

	fmt.Println(paragraphs)
}
