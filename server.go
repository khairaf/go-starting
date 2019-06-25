package main

import (
	"fmt"
	"net/http"
	"html/template"
	"encoding/xml"
	"io/ioutil"
	"strings"
)

type NewsMap struct {
	Keyword string
	Location string
}

type NewsAggPage struct {
	Title string
	News map[string]NewsMap
}

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Whoaa,, look at this!</h1>")
	fmt.Fprintf(w, "<p>Go is fast!</p>")
	fmt.Fprintf(w, "<p>...and simple!</p>")	
	fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>variables</strong>")	
	fmt.Fprintf(w, `
		<h6>You can even do ...</h6>
		<h5>multiple lines ...</h5>
		<h4>in one %s</h4>`, "formatted print")
}

func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Ini adalah an about page")
}

func newsAggHandler (w http.ResponseWriter, r *http.Request){
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

	p := NewsAggPage{Title: "Awesome Golang", News: news_map}
	t, _ := template.ParseFiles("aggtable.html")
	fmt.Println(t.Execute(w, p))
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.HandleFunc("/agg", newsAggHandler)
	http.ListenAndServe(":8000", nil) //nill = nothing to pass as server argument
}

//http://127.0.0.1:8000/ or view-source:http://127.0.0.1:8000/