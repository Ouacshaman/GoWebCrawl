package main

import (
	"fmt"
	"sort"
	"strings"
)

type Page struct{
	Url   string
	Count int
}

func printReport(pages map[string]int, baseURL string){
	fmt.Println("=============================")
	fmt.Println("REPORT for", baseURL)
	fmt.Println("=============================")
	sorted := sortReport(pages)
	for _, page := range sorted{
		baseURLTrimmed := strings.TrimPrefix(baseURL, "https://")
		if !strings.HasPrefix(strings.TrimRight(page.Url, "/"), strings.TrimRight(baseURLTrimmed, "/")){
			continue
		}
		output := fmt.Sprintf("Found %d internal links to %s", page.Count, page.Url)
		fmt.Println(output)
	}
}

func sortReport(pages map[string]int) []Page{
	var res []Page
	for k, v := range pages{
		temp := Page{
			Url: k,
			Count: v,
		}
		res = append(res, temp)
	}
	sort.Slice(res, func(i,j int) bool {
		if res[i].Count != res[j].Count{
			return res[i].Count > res[j].Count
		}
		return res[i].Url < res[j].Url
	})
	return res
}
