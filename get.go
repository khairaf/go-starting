package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

func main() {
	var s Sitemapindex
	var n News
	news_map := make(map[string]NewsMap)
	resp, err := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	if err != nil {
		fmt.Printf("error yaitu: %v\n", err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error adalah: %v\n", err)
	}
	
	xml.Unmarshal(bytes, &s)
	

	for _, Location := range s.Locations {
		t := strings.TrimSpace(Location)
		resp, err := http.Get(t)
		if err != nil {
			fmt.Printf("errornya: %v\n", err)
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("error di: %v\n", err)
		}

		xml.Unmarshal(bytes, &n)
		
		
		for idx, _ := range n.Keywords {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	for idx, data := range news_map {
		fmt.Println("\n\n\n",idx)
		fmt.Println("\n",data.Keyword)
		fmt.Println("\n",data.Location)
	}

	
}