package outfile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Out2txt(arr string, filename string) {
	url := "https://api.qq.jsososo.com/song/urls?id=" + arr
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	all, _ := ioutil.ReadAll(response.Body)
	//json str 转map
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(all), &m); err == nil {
		m2 := m["data"].(map[string]interface{})
		durl := []string{}
		for _, s2 := range m2 {
			durl = append(durl, s2.(string))
		}
		join := strings.Join(durl, "\n")
		if ioutil.WriteFile(filename+".txt", []byte(join), 0644) == nil {
			fmt.Println("写入文件成功")
		}
	}
}
