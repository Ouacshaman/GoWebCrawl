package main

import (
	"strings"
	"golang.org/x/net/html"
	"net/url"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error){
	if htmlBody == ``{
		return []string{}, nil
	}
	_, err := url.Parse(rawBaseURL)
	if err != nil{
		return []string{}, nil
	}
	doc, err:= html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}
	var f func(*html.Node, *[]string)
	f = func(n *html.Node, links *[]string){
		if n.Type == html.ElementNode && n.Data == "a"{
			for _, a := range n.Attr{
				if a.Key == "href" {
					u, err := url.Parse(a.Val)
					if err != nil{
						break
					}
					res := a.Val 
					if !u.IsAbs(){
							res = rawBaseURL + res
					}
					(*links) = append((*links), res)
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
