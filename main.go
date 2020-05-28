package main

import (
	"encoding/json"
	"goqmusic/src/out/outfile"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args
	keyword := ""
	if len(args) == 1 {
		println("缺失参数: -k")
		return
	}
	for index, arg := range args {
		if strings.EqualFold("-k", arg) {
			keyword = args[index+1]
		}
	}
	if len(keyword) == 0 {
		println("参数错误")
		return
	}
	request, _ := http.NewRequest("GET", "http://localhost:3300/search?key="+keyword+"&pageSize=50", nil)
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
