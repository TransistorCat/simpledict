package dict

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type TencentDictRequest struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Word   string `json:"sourceText"`
}

type tencentDictResponse struct {
	SessionUUID string `json:"sessionUuid"`
	Translate   struct {
		ErrCode     int    `json:"errCode"`
		ErrMsg      string `json:"errMsg"`
		SessionUUID string `json:"sessionUuid"`
		Source      string `json:"source"`
		Target      string `json:"target"`
		Records     []struct {
			SourceText string `json:"sourceText"`
			TargetText string `json:"targetText"`
			TraceID    string `json:"traceId"`
		} `json:"records"`
		Full    bool `json:"full"`
		Options struct {
		} `json:"options"`
	} `json:"translate"`
	Dict    interface{} `json:"dict"`
	Suggest interface{} `json:"suggest"`
	ErrCode int         `json:"errCode"`
	ErrMsg  string      `json:"errMsg"`
}

func TencentQuery(word string) string {
	client := &http.Client{}

	var data = fmt.Sprintf(`source=auto&target=zh&sourceText=%v&qtv=5d94939721cd5d83&qtk=wgoABMjwqwRTOA8mIglYINu24XRtctJF%%2BCeRb0x0mrdcIlBNHH%%2BEAN8QXSFci0KBRU6l0mEiKzxtU8t6kcR3g0CcKcILoELs60TAc3EcRBaC6wz2J0WnpqnK6Ln8e0NS8ep2b%%2BsdLaTkolgO92Kp%%2FQ%%3D%%3D&ticket=&randstr=&sessionUuid=translate_uuid1690264966595`, word)
	var reader = strings.NewReader(data)
	req, err := http.NewRequest("POST", "https://fanyi.qq.com/api/translate", reader)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "pac_uid=0_db87cdfebced3; iip=0; RK=gS3U0tw87w; ptcz=0ae1c2712b41401f3f4415c466d10122474b6e613cb2a59d79e0921de7b15e38; eas_sid=c1D63848Q4T4Y9H2P2b685H700; pgv_pvid=2761494972; ptui_loginuin=2328988259; uin=o1820737440; skey=@9X0qaPxWg; fy_guid=8ca96ccc-5598-4c65-8cfc-c6b511a0376d; ADHOC_MEMBERSHIP_CLIENT_ID1.0=40e390f3-042b-7b85-052b-34184747b2fb; openCount=1; qtv=5d94939721cd5d83; qtk=wgoABMjwqwRTOA8mIglYINu24XRtctJF+CeRb0x0mrdcIlBNHH+EAN8QXSFci0KBRU6l0mEiKzxtU8t6kcR3g0CcKcILoELs60TAc3EcRBaC6wz2J0WnpqnK6Ln8e0NS8ep2b+sdLaTkolgO92Kp/Q==")
	req.Header.Set("Origin", "https://fanyi.qq.com")
	req.Header.Set("Referer", "https://fanyi.qq.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.183")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Microsoft Edge";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("uc", "nrflfIMpW2bjVlg4tapAzkRs60kwXc4K1UlrujUaOm8=")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", bodyText)
	// return ""
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictResponse tencentDictResponse
	var result string

	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	result = "腾讯翻译:" + "\n" + dictResponse.Translate.Records[0].TargetText + "\n"
	return result
}
