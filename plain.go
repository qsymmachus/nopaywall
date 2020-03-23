package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/prometheus/common/log"
)

const (
	DefaultURL   = "https://en.wikipedia.org/wiki/%22Hello,_World!%22_program"
	TextSelector = "p, h1, h2, h3, h4, h5, h6"
)

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

// Given an HTTP response, finds all HTML "text" tags and extracts their text content.
// What we consider a "text tag" is defined in the `TextSelector` constant. Returns a
// plaintext string of all the extracted text.
func ExtractText(response *http.Response) (string, error) {
	defer response.Body.Close()
	var textContents []string

	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return "", err
	}

	doc.Find(TextSelector).Each(func(i int, s *goquery.Selection) {
		textContents = append(textContents, FormatText(s))
	})

	return strings.Join(textContents, "\n\n"), nil
}

// Extracts and formats the text from a selected HTML tag. We capitalize headers, and
// remove extra newlines that may be in paragraphs.
func FormatText(s *goquery.Selection) string {
	var text string

	switch s.Nodes[0].Data {
	case "p":
		text = strings.ReplaceAll(s.Text(), "\n", " ")
	case "h1", "h2", "h3", "h4", "h5", "h6":
		text = strings.ToUpper(s.Text())
	}

	return text
}

// Retrieves the document at the URL specified by the '-url' flag, and prints a
// plaintext representation of its content to standard output. For example:
//
//  nopaywall -url=http://example.com
//
func main() {
	url := flag.String("url", DefaultURL, "URL of the page you'd like to read")
	flag.Parse()

	response, err := LoadPage(*url)
	if err != nil {
		log.Error(err)
	}

	text, err := ExtractText(response)
	if err != nil {
		log.Error(err)
	}

	fmt.Println(text)
}
