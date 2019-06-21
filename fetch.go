package main

import (
	"fmt"
	"io/ioutil" //untuk membaca data
	"net/http" //untuk membuat request
	"encoding/xml"
)

// var washPostXML = []byte(`
// <sitemapindex>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
//    </sitemap>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
//    </sitemap>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
//    </sitemap>
// </sitemapindex>
// `)


type Sitemapindex struct {
  Locations []Location `xml:"sitemap"`
}

//membuat list/slice bernama Locations (musti huruf besar L nya) dengan type Location
//Location bukan build-in type, jadi perlu bikin struct Location dulu di bawah
//xml: sitemap artinya dia akan diarahkan buat nyari si tag <sitemap>

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
  return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body) //untuk mengakses body data
	//string_body := string(bytes) //untuk rubah body data menjadi string
	//fmt.Println(string_body)
	//fmt.Println(resp.Body)
	resp.Body.Close()

	//cara 2
	// bytes := washPostXML

	var s Sitemapindex
	xml.Unmarshal(bytes, &s)
	
	//fmt.Println(s.Locations)

	for _, Location := range s.Locations {
		fmt.Printf("%s\n", Location)
	}	
}

//An array with 6 integers might look something like: var ArrExample [6]int.
//How about a slice? We could make a slice with SliceExample []float32. This would be a slice that's made up for float32 values.

/*
// package main

// import (
//   "encoding/xml"
//   "fmt"
// )

// var washPostXML = []byte(`
// <sitemapindex>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
//    </sitemap>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
//    </sitemap>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
//    </sitemap>
// </sitemapindex>
// `)

// type Sitemapindex struct {
//   Locations []Location `xml:"sitemap"`
// }

// type Location struct {
//   Loc string `xml:"loc"`
// }

// func (e Location) String () string {
//   return fmt.Sprintf(e.Loc)
// }

// func main() {
//   bytes := washPostXML
//   var s Sitemapindex
//   xml.Unmarshal(bytes, &s)
//   fmt.Println(s.Locations)
// }
*/