package outfile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Out2txt(id string, filename string) {
	url := "https://api.qq.jsososo.com/song/url?type=320&id=" + id
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	all, _ := ioutil.ReadAll(response.Body)
	//json str 转map
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(all), &m); err == nil {
		m2 := m["data"]
		fmt.Println(m2)
		fl, err := os.OpenFile(filename+".txt", os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return
		}
		defer fl.Close()
		fl.Write([]byte(m2.(string) + "\n"))
		fmt.Println("追加文件成功")
	}
}
