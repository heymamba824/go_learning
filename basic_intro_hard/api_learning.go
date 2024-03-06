package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// curlconverter.com/#go

// 简单例子，通过urlc转化
func test() {
	client := &http.Client{}
	var data = strings.NewReader(`{"source":"nihao","trans_type":"auto2zh","request_id":"web_fanyi","media":"text","os_type":"web","dict":true,"cached":true,"replaced":true,"style":"formal","model":"","detect":true,"browser_id":"3f7bc0f83c3061a35eb7ae55b97e202e"}`)
	req, err := http.NewRequest("POST", "https://lingocloud.caiyunapp.com/v1/translator", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Cookie", "_gcl_au=1.1.2005912482.1709432729; _gid=GA1.2.298398224.1709432730; _gat_gtag_UA_185151443_2=1; _ga=GA1.2.59509220.1709432729; _ga_65TZCJSDBD=GS1.1.1709432729.1.1.1709432871.0.0.0; _ga_R9YPR75N68=GS1.1.1709432729.1.1.1709432871.37.0.0")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("T-Authorization", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJicm93c2VyX2lkIjoiM2Y3YmMwZjgzYzMwNjFhMzVlYjdhZTU1Yjk3ZTIwMmUiLCJpcF9hZGRyZXNzIjoiNzIuMjkuMjA0LjIzMCIsInRva2VuIjoicWdlbXY0anIxeTM4anlxNnZodmkiLCJ2ZXJzaW9uIjoxLCJleHAiOjE3MDk0MzM3NzJ9.nsX0Ybbs1TkVgcQJTxvvcouY-EVgJO4ac34n-gBzT_Q")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("app-name", "xy")
	req.Header.Set("device-id", "3f7bc0f83c3061a35eb7ae55b97e202e")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua", `"Chromium";v="122", "Not(A:Brand";v="24", "Google Chrome";v="122"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

type DictRequest struct {
	Transtype string `json:"trans_type"`
	Source    string `json:"source"`
}

func main() {

}
