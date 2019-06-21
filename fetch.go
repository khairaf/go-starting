package main

import (
	"fmt"
	"io/ioutil" //untuk membaca data
	"net/http" //untuk membuat request
)

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body) //untuk mengakses body data
	string_body := string(bytes) //untuk rubah body data menjadi string
	fmt.Println(string_body)
	fmt.Println(resp.Body)
	resp.Body.Close()
}