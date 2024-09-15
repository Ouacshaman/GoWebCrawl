package main

import (
	"fmt"
)

type Page struct{
	Url   string
	Count int
}

func printReport(pages map[string]int, baseURL string){
	fmt.Println("Report for ", baseURL)
	sorted := sortReport(pages)
	for _, page := range sorted{
		fmt.Println(page.Url, page.Count)
	}
}

func sortReport(pages map[string]int) []Page{
	var res []Page
	for k, v := range pages{
		if len(res) > 0 {
			entry := Page{
				Url: k,
				Count: v,
			}
			res = append(res, entry)
			end := len(res)
			swapping := true
			for swapping{
				swapping = false
				for i:=1; i<end; i++{
					if res[i-1].Count < res[i].Count{
						temp := res[i-1]
						res[i-1] = res[i]
						res[i] = temp
						swapping = true;
					}
				}
				end -= 1
			}
		}else{
			initial := Page{
				Url: k,
				Count: v,
			}
			res = append(res, initial)
		}
	}
	return res
}
