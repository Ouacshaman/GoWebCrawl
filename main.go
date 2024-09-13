package main

import (
	"fmt"
	"os"
	"sync"
	"net/url"
	"strconv"
)

func main(){
	args := os.Args
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	fmt.Println("starting crawl of: ", args[1])
	url, err := url.Parse(args[1])
	if err != nil{
		fmt.Println(err)
		return
	}
	conCtrl, err := strconv.Atoi(args[2])
	if err != nil{
		fmt.Println("Invalid integer:", err)
		return
	}
	maxPagesAmt, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("Invalid integer:", err)
		return
	}
	res := make(map[string]int)
	config := config{
		pages: res,
		baseURL: url,
		mu: &sync.Mutex{},
		concurrencyControl: make(chan struct{}, conCtrl),
		wg: &sync.WaitGroup{},
		maxPages: maxPagesAmt,
	}
	config.wg.Add(1)
	go config.crawlPage(config.baseURL.String())
	config.wg.Wait()
	for key,val := range config.pages{
		fmt.Println(key, ": ", val)
	}
}
