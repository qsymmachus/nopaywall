package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/prometheus/common/log"
	"golang.org/x/net/html"
)

const DefaultURL = "https://en.wikipedia.org/wiki/%22Hello,_World!%22_program"

func LoadPage(url string) (*html.Tokenizer, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	tokenizer := html.NewTokenizer(response.Body)

	return tokenizer, nil
}

func ParseParagraphs(tokenizer *html.Tokenizer) string {
	var paragraphs []string

	for {
		nextToken := tokenizer.Next()

		switch nextToken {
		case html.StartTagToken:
			token := tokenizer.Token()

			isParagraph := token.Data == "p"
			if isParagraph {
				tokenizer.Next()
				text := tokenizer.Token().Data
				paragraphs = append(paragraphs, text)
			}
		case html.ErrorToken:
			// This usually means we hit the end of the document.
			return strings.Join(paragraphs, "\n\n")
		}
	}
}

func main() {
	url := flag.String("url", DefaultURL, "URL of the page you'd like to read")
	flag.Parse()

	tokenizer, err := LoadPage(*url)
	if err != nil {
		log.Error(err)
	}

	paragraphs := ParseParagraphs(tokenizer)

	fmt.Println(paragraphs)
}
