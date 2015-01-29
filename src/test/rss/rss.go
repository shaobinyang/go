package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	// "net/http"
)

func main() {
	// response, _ := http.Get("http://blog.atime.me/feed.xml")
	// defer response.Body.Close()
	// body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(body))
	xmlFile, err := os.Open("/home/ysb/tmp/feed.xml")
	if err != nil {
		fmt.Println("read file error: ", err)
		return
	}

	// 	data := `<?xml version="1.0" encoding="utf-8"?>
	// <rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
	// <channel>
	// <title>MWB日常笔记</title>
	// <link>http://blog.atime.me/</link>
	// <description>zzdsfd</description>
	// <atom:link href="http://blog.atime.me/feed.xml" rel="self"></atom:link>
	// <lastBuildDate>Wed, 07 Jan 2015 12:50:00 +0800</lastBuildDate>
	// </channel>
	// </rss>
	// `

	type Item struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
	}
	type ChannelInfo struct {
		Title         string `xml:"title"`
		Description   string `xml:"description"`
		Link          string `xml:"RssDefault link"`
		LastBuildDate string `xml:"lastBuildDate"`
		Items         []Item `xml:"item"`
	}
	type RSS struct {
		XMLName xml.Name    `xml:"rss"`
		Channel ChannelInfo `xml:"channel"`
	}
	var rss RSS
	decoder := xml.NewDecoder(xmlFile)
	decoder.DefaultSpace = "RssDefault"
	err2 := decoder.Decode(&rss)
	if err2 != nil {
		fmt.Println("error:%w", err2)
		return
	}
	// fmt.Printf("%+v\n", rss)

	var filtItems []Item
	for _, v := range rss.Channel.Items {
		if strings.Contains(v.Description, "install") {
			filtItems = append(filtItems, v)
		}
	}
	fmt.Printf("%+v\n", filtItems)
}
