package main

import (
	"strings"
	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error){
	doc, err:= html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}
	var f func(*html.Node, *[]string)
	f = func(n *html.Node, links *[]string){
		if n.Type == html.ElementNode && n.Data == "a"{
			for _, a := range n.Attr{
				if a.Key == "href" {
					(*links) = append((*links), a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c!= nil; c = c.NextSibling {
			f(c, links)
		}
	}
	var links []string
	f(doc, &links)
	return links, nil
}
