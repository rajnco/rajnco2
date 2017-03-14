package main

import "bytes"
import "encoding/xml"
import "fmt"
import "io/ioutil"
import "os"
import "net/http"

import "github.com/gin-gonic/gin"

func main() {
	parsedXml := DoAllXml()
	//DisplayXml(parsedXml)
    r := gin.Default()
    //r.LoadHTMLGlob("template/*")
    r.LoadHTMLFiles("template/cricinfo.tmpl")
    r.GET("/",func(c *gin.Context){
        c.JSON(200, gin.H{
            "msg": "index",
        })
    })

    r.GET("/cricinfo", func(c *gin.Context){
        c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "Cricinfo ", "items" : parsedXml.Channel.Items })
    })
    r.Run()

}


func DisplayXml(rss *RssFeed) {
	for _, item := range rss.Channel.Items {
		fmt.Println("Item Title  : ", item.Title)
		fmt.Println("Item Title  : ", item.Link)
	}
}

func DoAllXml() *RssFeed {
	//xmlUrl    := "http://feeds.reuters.com/reuters/topNew?format=xml"
    xmlUrl := "http://www.espncricinfo.com/rss/content/story/feeds/0.xml"
	xmlData   := pullXml(xmlUrl)
	xmlParsed := parseXml(xmlData)
    return xmlParsed
}

func pullXml(xmlUrl string) []byte {
	res, err := http.Get(xmlUrl)
	chkError(err)
	defer res.Body.Close()
	xmlData, err := ioutil.ReadAll(res.Body)
	chkError(err)
	return xmlData
}

func chkError(err error){
	if (err != nil) {
		fmt.Println("Error : ", err)
		os.Exit(0)
	}
}


func parseXml(x []byte) *RssFeed {
    feed := RssFeed{}
    d := xml.NewDecoder(bytes.NewReader(x))
    d.DefaultSpace = "RssDefault"
    err := d.Decode(&feed)
    chkError(err)
    return &feed
}



type RssFeed struct {
    XMLName xml.Name    `xml:"rss"`
    Channel *RssChannel `xml:"channel"`
}

type RssChannel struct {
    XMLName       xml.Name `xml:"channel"`
    Title         string   `xml:"title"`
    Description   string   `xml:"description"`
    Image         Image    `xml:"image"`
    Language      string   `xml:"language"`
    LastBuildDate string   `xml:"lastBuildDate"`
    CopyRight     string   `xml:"copyright"`
    Link          string   `xml:"RssDefault link"`
    PubDate       string   `xml:"pubDate"`
    Atom          Atom     `xml:"atom"`
    Items         []Item   `xml:"item"`
}

type Item struct {
    Title string `xml:"title"`
    Desc  string `xml:"description"`
    Link  string `xml:"link"`
    Guid  string `xml:"guid"`
    Cat   string `xml:"category"`
    Date  string `xml:"pubDate"`
    Feed  string `xml:"feedburner:origLink"`
}


type Image struct {
    XMLName xml.Name `xml:"image"`
    Url     string   `xml:"url"`
    Title   string   `xml:"title"`
    Link    string   `xml:"link"`
    Width   string   `xml:"width"`
    Height  string   `xml:"height"`
}

type Atom struct {
    XMLName xml.Name `xml:"atom"`
    Href    string   `xml:"href,attr"`
}


