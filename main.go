package main

import (
	"fmt"
	"os"
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
	crawl, err := crawlPage(args[1], args[1], res)
	if err != nil{
		fmt.Println(err)
	}
	for key,val := range crawl{
		fmt.Println(key, ": ", val)
	}
}
