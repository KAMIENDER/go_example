package animate

import (
	// "encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	// "fmt"
	// Models "test/models"
	"github.com/gocolly/colly"
)

type AnimateController struct {
	beego.Controller
}

type AnimateInfo struct {
	ID     string `json:"id"`
	IDView string `json:"idView"`
	URL    string `json:"url"`
	Title  string `json:"title"`
}

func (this *AnimateController) Get() {
	// var back AnimateInfo
	all := test()
	this.Data["json"] = &all
	this.ServeJSON()
	//这一句会多返回一个null
}

// 下面这个函数用于爬虫
func test() []AnimateInfo {
	var all []AnimateInfo
	var photo []string
	c := colly.NewCollector()
	tt := colly.NewCollector()
	t := 0
	id := 0

	// On every a element which has href attribute call callback
	// 下面找的是具有src=""属性的img标签
	c.OnHTML("a[title][href][target]", func(e *colly.HTMLElement) {
		e.ForEach("img[src]", func(_ int, elem *colly.HTMLElement) {

			temp := elem.Attr("src")
			if strings.HasPrefix(temp, "http") == false {
				temp = "http://www.yxdm.tv" + temp
			}
			photo = append(photo, temp)
		})

		// 下面是获取标题以及对应页面链接
		var temp AnimateInfo
		temp.Title = e.Attr("title")
		temp.URL = "http://www.yxdm.tv/" + e.Attr("href")

		if t > 6 && t%2 == 0 && id < 24 {
			temp.ID = strconv.Itoa(id)
			temp.IDView = photo[id]
			all = append(all, temp)
			id++
		}

		t++
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	tt.OnResponse(func(e *colly.Response) {
		fmt.Println(e)
	})

	tt.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
		fmt.Println(tt)
	})

	tt.Visit("http://www.yxdm.tv/resource/9868-9-1.html")

	// Start scraping on https://hackerspaces.org
	c.Visit("http://www.yxdm.tv/category.html?channel=17&area=%E6%97%A5%E6%9C%AC&orderby=pubdate")
	return all
}
