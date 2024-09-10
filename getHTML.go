package main

import (
	"net/http"
	"fmt"
	"io"
	"strings"
)

func getHTML(rawURL string) (string, error){
	resp, err := http.Get(rawURL)
	if err != nil{
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400{
		statusError := fmt.Errorf("Status Code is 400+")
		return "", statusError
	}
	values := resp.Header.Values("content-type")
	exist := false
	for _, v := range values{
		if strings.Contains(v,"text/html"){
			exist = true
		}
	}
	if !exist{
		contentError := fmt.Errorf("The content type does not have text/html")
		return "", contentError
	} 
	data, err := io.ReadAll(resp.Body)
	if err != nil{
		return "", err
	}
	return string(data), nil
}
