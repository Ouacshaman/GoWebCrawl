package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) (map[string]int, error) {
	bURL, err := url.Parse(rawBaseURL)
	if err != nil{
		return make(map[string]int), nil
	}
	cURL, err := url.Parse(rawCurrentURL)
	if err != nil{
		return pages, err
	}
	if bURL.Hostname() != cURL.Hostname(){
		return pages, nil
	}
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil{
		return pages, err
	}
	if _, ok := pages[normalizedURL]; ok {
		fmt.Println("Already visited:", normalizedURL)
		pages[normalizedURL] ++
		return pages, nil
	}
	pages[normalizedURL] = 1
	html, err := getHTML(rawCurrentURL)
	if err != nil{
		return pages, err
	}
	fmt.Println(rawCurrentURL)
	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil{
		return pages, err
	}
	for _,url := range urls{
		pages, err = crawlPage(rawBaseURL, url, pages)
		if err != nil{
			return pages, err
		}
	}
	return pages, nil
}
