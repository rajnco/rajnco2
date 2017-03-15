package main

import "bytes"
import "encoding/xml"
import "fmt"
import "io/ioutil"
import "os"
import "net/http"

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    //r.LoadHTMLGlob("template/*")
    r.LoadHTMLFiles("template/cricinfo.tmpl", "template/index.tmpl")

    /*
    r.GET("/",func(c *gin.Context){
        c.JSON(200, gin.H{
            "msg": "index",
        })
    })
    */

    r.GET("/", func(c *gin.Context){
        c.HTML(http.StatusOK, "index.tmpl", gin.H{ "title": "Online RSS Reader" })
    })

    r.GET("/cricinfo", Cricinfo)
    r.GET("/bbc", BBC)
    r.GET("/cnn", CNN)
    r.GET("/yahoo", Yahoo)
    r.GET("/nytimes", NYTimes)
    r.GET("/ft", FT)
    r.GET("/wsj", WSJ)
    r.GET("/cnet", CNET)
    r.GET("/pcworld", PCWORLD)
    r.GET("/cweekly", CWEEKLY)
    r.GET("/pcmag", PCMAG)
    r.GET("/nworld", Nworld)
    r.GET("/techr", TechR)
    r.GET("/cio", CIO)
    r.GET("/pcq", PCQ)
    r.GET("/toi", TOI)
    r.GET("/et", ET)
    r.Run()

}

func TOI(c *gin.Context) {
	parsedXml := DoAllXml("http://timesofindia.indiatimes.com/rssfeedstopstories.cms")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "Times of India News", "items" : parsedXml.Channel.Items })
}

func ET(c *gin.Context) {
	parsedXml := DoAllXml("http://economictimes.indiatimes.com/rssfeedsdefault.cms")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "Economic Times News", "items" : parsedXml.Channel.Items })
}

func PCQ(c *gin.Context) {
	parsedXml := DoAllXml("http://www.pcquest.com/rss-2-2/?cat_slug=open-source-and-linux")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "PC Quest News", "items" : parsedXml.Channel.Items })
}

func CIO(c *gin.Context) {
	parsedXml := DoAllXml("http://www.cio.com/index.rss")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "CIO News", "items" : parsedXml.Channel.Items })
}

func TechR(c *gin.Context) {
	parsedXml := DoAllXml("http://www.techrepublic.com/rssfeeds/articles/latest/")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "Tech Republic News", "items" : parsedXml.Channel.Items })
}

func Nworld(c *gin.Context) {
	parsedXml := DoAllXml("http://www.networkworld.com/category/opensource-subnet/index.rss")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "Network World News", "items" : parsedXml.Channel.Items })
}

func SlashDot(c *gin.Context) {
	parsedXml := DoAllXml("http://rss.slashdot.org/Slashdot/slashdotMain")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "Slashdot News", "items" : parsedXml.Channel.Items })
}

func PCMAG(c *gin.Context) {
	parsedXml := DoAllXml("http://feeds.pcmag.com/Rss.aspx/SectionArticles?sectionId=1489")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "PG Magazine News", "items" : parsedXml.Channel.Items })
}

func CWEEKLY(c *gin.Context) {
	parsedXml := DoAllXml("http://www.computerweekly.com/rss/All-Computer-Weekly-content.xml")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "Computer Weekly News", "items" : parsedXml.Channel.Items })
}

func PCWORLD(c *gin.Context) {
	parsedXml := DoAllXml("http://www.pcworld.com/index.rss")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "PC World News", "items" : parsedXml.Channel.Items })
}

func CNET(c *gin.Context) {
	parsedXml := DoAllXml("https://www.cnet.com/rss/news/")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "CNET News", "items" : parsedXml.Channel.Items })
}

func WSJ(c *gin.Context) {
	parsedXml := DoAllXml("http://www.wsj.com/xml/rss/3_7085.xml")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "Wall Street Journal News", "items" : parsedXml.Channel.Items })
}

func FT(c *gin.Context) {
	parsedXml := DoAllXml("http://www.ft.com/rss/home/asia")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "Financial Times News", "items" : parsedXml.Channel.Items })
}

func NYTimes(c *gin.Context) {
	parsedXml := DoAllXml("http://www.nytimes.com/services/xml/rss/nyt/HomePage.xml")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "NY Times News", "items" : parsedXml.Channel.Items })
}

func Cricinfo(c *gin.Context) {
	parsedXml := DoAllXml("http://www.espncricinfo.com/rss/content/story/feeds/0.xml")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "Cricinfo News ", "items" : parsedXml.Channel.Items })
}

func CNN(c *gin.Context) {
	parsedXml := DoAllXml("http://rss.cnn.com/rss/edition.rss")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "CNN News", "items" : parsedXml.Channel.Items })
}

func Yahoo(c *gin.Context) {
	parsedXml := DoAllXml("https://www.yahoo.com/news/rss")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "Yahoo News ", "items" : parsedXml.Channel.Items })
}

func BBC(c *gin.Context) {
	//parsedXml := DoAllXml("http://feeds.reuters.com/reuters/topNew?format=xml")
	parsedXml := DoAllXml("http://feeds.bbci.co.uk/news/world/asia/rss.xml")
    c.HTML(http.StatusOK, "cricinfo.tmpl", gin.H{ "title" : "BBC News ", "items" : parsedXml.Channel.Items })

}

func DisplayXml(rss *RssFeed) {
	for _, item := range rss.Channel.Items {
		fmt.Println("Item Title  : ", item.Title)
		fmt.Println("Item Title  : ", item.Link)
	}
}

func DoAllXml(xmlUrl string) *RssFeed {
    //xmlUrl := "http://www.espncricinfo.com/rss/content/story/feeds/0.xml"
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


