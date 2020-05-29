package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"goqmusic/src/out/outfile"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	var keyword string
	flag.StringVar(&keyword, "k", "", "搜索关键词")
	flag.Parse()
	if flag.NFlag() == 0 {
		usage()
		return
	}
	escape := url.QueryEscape(keyword)
	url := "http://localhost:3300/search?key=" + escape + "&pageSize=50"
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	all, _ := ioutil.ReadAll(response.Body)
	//json str 转map
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(all), &m); err == nil {
		data := m["data"]
		list := data.(map[string]interface{})["list"]
		for _, i2 := range list.([]interface{}) {
			outfile.Out2txt(i2.(map[string]interface{})["songmid"].(string), keyword+"下载")
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `Options:`)
	flag.PrintDefaults()
}
