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

// Sends an HTTP request to the specified URL and returns the response.
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

// Given an HTTP response, finds all HTML <p> tags and extracts their text content.
// Returns a plaintext string of all the extracted text.
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

// Retrieves the document at the URL specified by the '-url' flag, and prints a plaintext
// representation of its content to standard output. For example:
//
//  nopaywall -url=http://example.com
//
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
