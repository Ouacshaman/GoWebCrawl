package main

import (
	"fmt"
	"os"
	"sync"
)

func main(){
	args := os.Args
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	fmt.Println("starting crawl of: ", args[1])
	res := make(map[string]int)
	config := config{
		pages: res,
		baseURL: args[1],
		mu: &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 1),
		wg: &sync.WaitGroup{},
	}
	go config.crawlPage(config.baseURL)
	config.wg.Wait()
	for key,val := range config.pages{
		fmt.Println(key, ": ", val)
	}
}
