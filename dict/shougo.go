package dict

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func SougouQuery(word string) string {
	client := &http.Client{}
	var data = fmt.Sprintf("{\"from\":\"auto\",\"to\":\"zh-CHS\",\"text\":\"%v\",\"client\":\"pc\",\"fr\":\"browser_pc\",\"needQc\":1,\"s\":\"8dcdbc5f48b9b9f85a43f0307e471252\",\"uuid\":\"043ec12e-951a-4da2-830c-19ed278c1da8\",\"exchange\":false}", word)
	var reader = strings.NewReader(data)
	req, err := http.NewRequest("POST", "https://fanyi.sogou.com/api/transpc/text/result", reader)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Cookie", "ABTEST=7|1690272146|v17; SNUID=B2EA8618606661144688DE5F60ACFE90; SUID=D28AE678EF53A00A0000000064BF8192; wuid=1690272146201; FQV=c3910ab6646b96474ff93685f354a09c; translate.sess=538d064b-a58f-46d2-9713-8a5798f92958; SUV=1690272146179; SGINPUT_UPSCREEN=1690272146194")
	req.Header.Set("Origin", "https://fanyi.sogou.com")
	req.Header.Set("Referer", "https://fanyi.sogou.com/text?channel=document")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.183")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Microsoft Edge";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
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
	return ""
}
