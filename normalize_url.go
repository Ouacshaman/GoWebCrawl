package main

import (
	"net/url"
)

func normalizeURL(str string) (string, error){
	u, err := url.Parse(str)
	if err != nil{
		return "", nil
	}
	normalized := u.Host + u.Path
	return normalized, nil
}
