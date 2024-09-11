package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()
	cfg.concurrencyControl <- struct{}{}
	defer func() { <-cfg.concurrencyControl}()
	cURL, err := url.Parse(rawCurrentURL)
	if err != nil{
		fmt.Println(err)
		return
	}
	if cfg.baseURL.Hostname() != cURL.Hostname(){
		return
	}
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil{
		fmt.Println(err)
		return
	}
	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst{
		cfg.mu.Lock()
		fmt.Println("Already visited:", normalizedURL)
		cfg.pages[normalizedURL] ++
		cfg.mu.Unlock()
		return
	}
	html, err := getHTML(rawCurrentURL)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(rawCurrentURL)
	urls, err := getURLsFromHTML(html, cfg.baseURL.String())
	if err != nil{
		fmt.Println(err)
		return
	}
	for _,url := range urls{
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool){
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, ok := cfg.pages[normalizedURL]; ok{
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true
}
