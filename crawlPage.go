package main

import (
	"fmt"
	"net/url"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cURL, err := url.Parse(rawCurrentURL)
	if err != nil{
		fmt.Println(err)
		return
	}
	if baseURL.Hostname() != cURL.Hostname(){
		return
	}
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil{
		fmt.Println(err)
		return
	}
	isFirst := addPageVisit(normalizedURL)
	if !isFirst{
		fmt.Println("Already visited:", normalizedURL)
		cfg.pages[normalizedURL] ++
		return
	}
	cfg.pages[normalizedURL] = 1
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
		<-cfg.concurrencyControl
		go crawlPage(cfg.baseURL, url, cfg.pages)
	}
	cfg.wg.Wait()
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool){
	cfg.mu.Lock()
	cfg.concurrencyControl <- struct{}{}
	cfg.wg.Add(1)
	defer cfg.wg.Done()
	defer cfg.mu.Unlock()
	if _, ok := cfg.pages[normalizedURL]; ok{
		return false
	}
	return true
}
