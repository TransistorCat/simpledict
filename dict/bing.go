package dict

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func BingQuery(word string) string {
	client := &http.Client{}
	var data = fmt.Sprintf("&fromLang=en&text=%v%%0A&to=zh-Hans&token=0OYoRkfITSSclqYko18hewA8RYn3VzKW&key=1690269111078&tryFetchingGenderDebiasedTranslations=true", word)
	var reader = strings.NewReader(data)

	req, err := http.NewRequest("POST", "https://www.bing.com/ttranslatev3?isVertical=1&&IG=5AD4E08EE22442D9BB83590F14943BA2&IID=translator.5026", reader)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "www.bing.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("cookie", "MUIDB=1E3E8CC71E0F60B61B519FF81F4561B2; MUID=1E3E8CC71E0F60B61B519FF81F4561B2; SRCHD=AF=NOFORM; SRCHUID=V=2&GUID=FE5D440E81FB4DB68050EA2126A1E574&dmnchg=1; WLS=C=de4eae3eaa367403&N=Mr.; ANON=A=B2D76922F788071DA6AA74C2FFFFFFFF; SUID=A; SRCHS=PC=U531; USRLOC=HS=1&ELOC=LAT=39.00550079345703|LON=117.36669921875|N=%E6%B4%A5%E5%8D%97%E5%8C%BA%EF%BC%8C%E5%A4%A9%E6%B4%A5%E5%B8%82|ELT=4|; _tarLang=default=zh-Hans; _TTSS_IN=hist=WyJlbiIsImF1dG8tZGV0ZWN0Il0=&isADRU=0; _TTSS_OUT=hist=WyJ6aC1IYW5zIl0=; _U=1FHF9HmdxOcxpJCO9AJ-ODFPaBoc5_fVxCCDoZZ2fB7JUb2SOjtQOoi7_MuTNQyJYU5sQUUhqnCCXJciNdWc_kbqijFJQs-7XgRMByzQ8DkOIUm6PxQvFhmVwsIL5L6Il_fC5pHIQMZN4LQrpFOlBzpI1phYoHmzgUnuhRFVxfD-J4YQZVlfqfXP9XnmQPq40HxEayUTqRypEm2X4HdjACw; SRCHUSR=DOB=20230725&T=1690268772000; ipv6=hit=1690272375606&t=4; _RwBf=mta=0&rc=4832&rb=4832&gb=0&rg=0&pc=4829&mtu=0&rbb=0.0&g=0&cid=&clo=0&v=2&l=2023-07-25T07:00:00.0000000Z&lft=0001-01-01T00:00:00.0000000&aof=0&o=16&p=bingcopilotwaitlist&c=MY00IA&t=4421&s=2023-02-13T01:53:26.4294290+00:00&ts=2023-07-25T07:11:49.5373460+00:00&rwred=0&wls=2&lka=0&lkt=0&TH=&dci=3&e=1VrVOVYRvBW1YN61Ntopfg-h9DcCFUUbFLL8-blnHBbXSJnbAbY05BJPBSsQ_-vc22nWh03qOKRGJnS7JgqPmXZRByZvy733v4T0RSoMONM&A=; _SS=SID=20568CE9AC7368F73A5F9FB0AD81691B&PC=U531&R=4832&RB=4832&GB=0&RG=0&RP=4829; _EDGE_S=SID=20568CE9AC7368F73A5F9FB0AD81691B&mkt=zh-cn&ui=zh-cn; SRCHHPGUSR=SRCHLANG=zh-Hans&IG=26E6C65F6E5840749A06E9AF74ECF7D6&PV=15.0.0&BZA=0&BRW=W&BRH=M&CW=1415&CH=796&SCW=1400&SCH=3945&DPR=1.8&UTC=480&DM=0&EXLTT=19&HV=1690269112&PRVCW=1415&PRVCH=796; btstkn=fKeTNLOVsJt55Vqr8OVYN98UfRwamEfe85XRf%252Bw%252FYnyBzHXfnvGPCDXVzZMdu8SBAYLh1g1GGdYnoN52xHegZ6H9%252Bc0rx99X%252Ff4Lsr0uPfo%253D")
	req.Header.Set("origin", "https://www.bing.com")
	req.Header.Set("referer", "https://www.bing.com/translator/?h_text=msn_ctxt&setlang=zh-cn")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Microsoft Edge";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-arch", `"x86"`)
	req.Header.Set("sec-ch-ua-bitness", `"64"`)
	req.Header.Set("sec-ch-ua-full-version", `"115.0.1901.183"`)
	req.Header.Set("sec-ch-ua-full-version-list", `"Not/A)Brand";v="99.0.0.0", "Microsoft Edge";v="115.0.1901.183", "Chromium";v="115.0.5790.99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-model", `""`)
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-ch-ua-platform-version", `"15.0.0"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-ms-gec", "F0DB2948A38CB65B6745A224FFAA32E916B745164FBB483675A380F18858F95E")
	req.Header.Set("sec-ms-gec-version", "1-115.0.1901.183")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.183")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	fmt.Println(bodyText)
	if err != nil {
		log.Fatal(err)
	}
	var result string
	fmt.Printf("%s\n", bodyText)
	return result
}
