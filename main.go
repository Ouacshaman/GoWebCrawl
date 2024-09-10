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
	htmlBody, err := getHTML(args[1])
	if err != nil{
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(htmlBody)
}
