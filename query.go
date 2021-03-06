package excavator

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/godcong/excavator/net"
)

//type Query interface {
//	//Header() http.Header
//	SetHeader(header http.Header)
//	//URL() string
//	AJAX(body io.Reader) ([]byte, error)
//}

//func NewQuery(url string) Query {
//	return &query{url: url}
//}
//
//func (q *query) AJAX(body io.Reader) (b []byte, e error) {
//	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
//	tr := &http.Transport{
//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
//	}
//	client := &http.Client{Transport: tr}
//	//body := strings.NewReader(`wd=%E4%B9%99`)
//	req, err := http.NewRequest("POST", q.url, body)
//	if err != nil {
//		return nil, err
//	}
//	req.Header = q.header
//
//	resp, err := client.Do(req)
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//	return ioutil.ReadAll(resp.Body)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//return UnmarshalRadical(bytes)
//}

//type RequestType int
type RequestFunc func(wd string, qb string) (*http.Request, error)

//const (
//	RequestTypeHanChengBushou RequestType = iota
//	RequestTypeHanChengPinyin
//	RequestTypeKangXiBushou
//	RequestTypeKangXiPinyin
//	RequestTypeDummy
//)

var requestList = []RequestFunc{
	RadicalTypeHanChengBushou: HanChengBushouRequest,
	RadicalTypeHanChengPinyin: HanChengPinyinRequest,
	RadicalTypeHanChengBihua:  HanChengBihuaRequest,
	RadicalTypeHanChengSo:     HanChengSoRequest,
	RadicalTypeKangXiBushou:   KangXiBushouRequest,
	RadicalTypeKangXiPinyin:   KangXiPinyinRequest,
	RadicalTypeKangXiBihua:    KangXiBihuaRequest,
	RadicalTypeKangXiSo:       KangXiSoRequest,
	//RequestTypeDummy:          DummyRequest,
}

type Query struct {
	cache *net.Cache
}

type QueryOptions func(query *Query)

func NewQuery(ops ...QueryOptions) *Query {
	q := &Query{
		cache: nil,
	}

	for _, op := range ops {
		op(q)
	}
	return q
}

func CacheOption(cache *net.Cache) QueryOptions {
	return func(query *Query) {
		query.cache = cache
	}
}

type GrabReader func(wd string, qb ...string) (reader io.ReadCloser, err error)

func (q *Query) Grab(radicalType RadicalType) GrabReader {
	r := requestList[radicalType]
	return func(wd string, qb ...string) (reader io.ReadCloser, err error) {
		arg := ""
		if len(qb) > 0 {
			arg = qb[0]
		}

		request, err := r(wd, arg)
		if err != nil {
			return nil, err
		}
		response, err := net.Request(request)
		if err != nil {
			return nil, err
		}
		closer := response.Body
		if q.cache != nil {
			name := request.URL.String()
			closer, err = q.cache.Cache(response.Body, name)
		}
		return closer, err
	}
}

func DummyRequest(wd string, qb int) (*http.Request, error) {
	log.With("wd", wd, "qb", qb).Info("dummy")
	return nil, errors.New("dummy request call")
}

func HanChengBushouRequest(wd string, qb string) (*http.Request, error) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	body := strings.NewReader(`wd=` + url.QueryEscape(wd))
	req, err := http.NewRequest("POST", "http://hy.httpcn.com/bushou/zi/", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", "UM_distinctid=16c2efb5e9e134-0cfc801ee6ae06-353166-1fa400-16c2efb5e9f3c8; Hm_lvt_cd7ed86134e0e3138a7cf1994e6966c8=1564156322,1565017408; CNZZDATA1267010321=1299014713-1564151968-^%^7C1565183518; hy_so_1=^%^255B^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E6^%^259D^%^258E^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252227^%^252FPWTBILILTBTBMEILE^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522l^%^25C7^%^2590^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E6^%^259C^%^25A8^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25227^%^2522^%^257D^%^255D; hy_so_4=^%^255B^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E8^%^2592^%^258B^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252234^%^252FKOKORNKOCQXVILXVB^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522ji^%^25C7^%^258Eng^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E8^%^2589^%^25B9^%^2522^%^252C^%^2522num^%^2522^%^253A^%^252217^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E6^%^259D^%^258E^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252227^%^252FPWTBILILTBTBMEILE^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522l^%^25C7^%^2590^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E6^%^259C^%^25A8^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25227^%^2522^%^257D^%^255D; ASP.NET_SessionId=lceymsersbrsquifblfr2ewm")
	req.Header.Set("Origin", "http://hy.httpcn.com")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7,zh-TW;q=0.6")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "http://hy.httpcn.com/bushou/zi/")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	return req, nil
}

func HanChengPinyinRequest(wd string, qb string) (*http.Request, error) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	body := strings.NewReader(`wd=` + wd)
	req, err := http.NewRequest("POST", "http://hy.httpcn.com/pinyin/zi/", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", "UM_distinctid=16c2efb5e9e134-0cfc801ee6ae06-353166-1fa400-16c2efb5e9f3c8; Hm_lvt_cd7ed86134e0e3138a7cf1994e6966c8=1564156322,1565017408; CNZZDATA1267010321=1299014713-1564151968-^%^7C1565183518; hy_so_4=^%^255B^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E8^%^2592^%^258B^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252234^%^252FKOKORNKOCQXVILXVB^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522ji^%^25C7^%^258Eng^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E8^%^2589^%^25B9^%^2522^%^252C^%^2522num^%^2522^%^253A^%^252217^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E6^%^259D^%^258E^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252227^%^252FPWTBILILTBTBMEILE^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522l^%^25C7^%^2590^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E6^%^259C^%^25A8^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25227^%^2522^%^257D^%^255D; ASP.NET_SessionId=yifpdjyctkr5nzkj5bm24ag1; hy_so_1=^%^255B^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E8^%^2592^%^258B^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252234^%^252FKOKORNKOCQXVILXVB^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522ji^%^25C7^%^258Eng^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E8^%^2589^%^25B9^%^2522^%^252C^%^2522num^%^2522^%^253A^%^252212^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E6^%^259D^%^258E^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252227^%^252FPWTBILILTBTBMEILE^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522l^%^25C7^%^2590^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E6^%^259C^%^25A8^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25227^%^2522^%^257D^%^255D")
	req.Header.Set("Origin", "http://hy.httpcn.com")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7,zh-TW;q=0.6")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "http://hy.httpcn.com/pinyin/zi/")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	return req, nil
}

func HanChengBihuaRequest(wd string, qb string) (*http.Request, error) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	body := strings.NewReader(`wd=` + wd + `&qb=` + qb)
	req, err := http.NewRequest("POST", "http://hy.httpcn.com/bihua/zi/", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", "UM_distinctid=16c2efb5e9e134-0cfc801ee6ae06-353166-1fa400-16c2efb5e9f3c8; Hm_lvt_cd7ed86134e0e3138a7cf1994e6966c8=1564156322,1565017408; CNZZDATA1267010321=1299014713-1564151968-^%^7C1565183518; hy_so_4=^%^255B^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E8^%^2592^%^258B^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252234^%^252FKOKORNKOCQXVILXVB^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522ji^%^25C7^%^258Eng^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E8^%^2589^%^25B9^%^2522^%^252C^%^2522num^%^2522^%^253A^%^252217^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E6^%^259D^%^258E^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252227^%^252FPWTBILILTBTBMEILE^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522l^%^25C7^%^2590^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E6^%^259C^%^25A8^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25227^%^2522^%^257D^%^255D; hy_so_1=^%^255B^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E8^%^2592^%^258B^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252234^%^252FKOKORNKOCQXVILXVB^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522ji^%^25C7^%^258Eng^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E8^%^2589^%^25B9^%^2522^%^252C^%^2522num^%^2522^%^253A^%^252212^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E6^%^259D^%^258E^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252227^%^252FPWTBILILTBTBMEILE^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522l^%^25C7^%^2590^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E6^%^259C^%^25A8^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25227^%^2522^%^257D^%^255D; ASP.NET_SessionId=2u2rhxdtt3hux2ujokuiemjw")
	req.Header.Set("Origin", "http://hy.httpcn.com")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7,zh-TW;q=0.6")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "http://hy.httpcn.com/bihua/zi/")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	return req, nil
}

func HanChengSoRequest(wd string, qb string) (*http.Request, error) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	body := strings.NewReader(`wd=` + url.QueryEscape(wd))
	req, err := http.NewRequest("POST", "http://hy.httpcn.com/so/", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Origin", "http://hy.httpcn.com")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36 Edg/79.0.309.65")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "http://hy.httpcn.com/so/")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cookie", "ASP.NET_SessionId=3ubpl3ahk3uohsrcymecoowd; UM_distinctid=16f95b5a85b996-0e0b825e2df3bc-7e657561-1fa400-16f95b5a85c87f; CNZZDATA1267010321=1099494227-1579099638-%7C1579099638; Hm_lvt_cd7ed86134e0e3138a7cf1994e6966c8=1578764708; Hm_lpvt_cd7ed86134e0e3138a7cf1994e6966c8=1578764720")
	return req, nil
}

func KangXiBushouRequest(wd string, qb string) (*http.Request, error) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	body := strings.NewReader(`wd=` + url.QueryEscape(wd))
	req, err := http.NewRequest("POST", "http://hy.httpcn.com/bushou/kangxi/", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", "UM_distinctid=16c2efb5e9e134-0cfc801ee6ae06-353166-1fa400-16c2efb5e9f3c8; Hm_lvt_cd7ed86134e0e3138a7cf1994e6966c8=1564156322,1565017408; CNZZDATA1267010321=1299014713-1564151968-^%^7C1565183518; ASP.NET_SessionId=2u2rhxdtt3hux2ujokuiemjw; hy_so_1=^%^255B^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E8^%^2592^%^258B^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252234^%^252FKOKORNKOCQXVILXVB^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522ji^%^25C7^%^258Eng^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E8^%^2589^%^25B9^%^2522^%^252C^%^2522num^%^2522^%^253A^%^252212^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E4^%^25B8^%^2591^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252220^%^252FCQRNRNXVUYILECQCQ^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522ch^%^25C7^%^2592u^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E4^%^25B8^%^2580^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25224^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E5^%^2590^%^2583^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252222^%^252FPWCQUYAZMEUYILAZKO^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522ch^%^25C4^%^25AB^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E5^%^258F^%^25A3^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25226^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E6^%^259D^%^258E^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252227^%^252FPWTBILILTBTBMEILE^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522l^%^25C7^%^2590^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E6^%^259C^%^25A8^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25227^%^2522^%^257D^%^255D; hy_so_4=^%^255B^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E9^%^2582^%^25AA^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252238^%^252FKOMEAZKOILRNAZAA^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522xi^%^25C3^%^25A9^%^252Cy^%^25C3^%^25A9^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E9^%^2598^%^259D^%^2522^%^252C^%^2522num^%^2522^%^253A^%^252211^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E4^%^25B8^%^2591^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252220^%^252FCQRNRNXVUYILECQCQ^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522ch^%^25C7^%^2592u^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E4^%^25B8^%^2580^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25224^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E5^%^2590^%^2583^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252222^%^252FPWCQUYAZMEUYILAZKO^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522ch^%^25C4^%^25AB^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E5^%^258F^%^25A3^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25226^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E8^%^2592^%^258B^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252234^%^252FKOKORNKOCQXVILXVB^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522ji^%^25C7^%^258Eng^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E8^%^2589^%^25B9^%^2522^%^252C^%^2522num^%^2522^%^253A^%^252217^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E6^%^259D^%^258E^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252227^%^252FPWTBILILTBTBMEILE^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522l^%^25C7^%^2590^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E6^%^259C^%^25A8^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25227^%^2522^%^257D^%^255D")
	req.Header.Set("Origin", "http://hy.httpcn.com")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7,zh-TW;q=0.6")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "http://hy.httpcn.com/bushou/kangxi/")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	return req, nil

}

func KangXiPinyinRequest(wd string, qb string) (*http.Request, error) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	body := strings.NewReader(`wd=` + wd)
	req, err := http.NewRequest("POST", "http://hy.httpcn.com/pinyin/kangxi/", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", "UM_distinctid=16c2efb5e9e134-0cfc801ee6ae06-353166-1fa400-16c2efb5e9f3c8; Hm_lvt_cd7ed86134e0e3138a7cf1994e6966c8=1564156322,1565017408; CNZZDATA1267010321=1299014713-1564151968-^%^7C1565183518; hy_so_4=^%^255B^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E8^%^2592^%^258B^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252234^%^252FKOKORNKOCQXVILXVB^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522ji^%^25C7^%^258Eng^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E8^%^2589^%^25B9^%^2522^%^252C^%^2522num^%^2522^%^253A^%^252217^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E6^%^259D^%^258E^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252227^%^252FPWTBILILTBTBMEILE^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522l^%^25C7^%^2590^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E6^%^259C^%^25A8^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25227^%^2522^%^257D^%^255D; ASP.NET_SessionId=yifpdjyctkr5nzkj5bm24ag1; hy_so_1=^%^255B^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E8^%^2592^%^258B^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252234^%^252FKOKORNKOCQXVILXVB^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522ji^%^25C7^%^258Eng^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E8^%^2589^%^25B9^%^2522^%^252C^%^2522num^%^2522^%^253A^%^252212^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E6^%^259D^%^258E^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252227^%^252FPWTBILILTBTBMEILE^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522l^%^25C7^%^2590^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E6^%^259C^%^25A8^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25227^%^2522^%^257D^%^255D")
	req.Header.Set("Origin", "http://hy.httpcn.com")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7,zh-TW;q=0.6")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "http://hy.httpcn.com/pinyin/kangxi/")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	return req, nil
}

func KangXiBihuaRequest(wd string, qb string) (*http.Request, error) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	body := strings.NewReader(`wd=` + wd + `&qb=` + qb)
	req, err := http.NewRequest("POST", "http://hy.httpcn.com/bihua/kangxi/", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", "UM_distinctid=16c2efb5e9e134-0cfc801ee6ae06-353166-1fa400-16c2efb5e9f3c8; Hm_lvt_cd7ed86134e0e3138a7cf1994e6966c8=1564156322,1565017408; CNZZDATA1267010321=1299014713-1564151968-^%^7C1565183518; hy_so_4=^%^255B^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E8^%^2592^%^258B^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252234^%^252FKOKORNKOCQXVILXVB^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522ji^%^25C7^%^258Eng^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E8^%^2589^%^25B9^%^2522^%^252C^%^2522num^%^2522^%^253A^%^252217^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E6^%^259D^%^258E^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252227^%^252FPWTBILILTBTBMEILE^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522l^%^25C7^%^2590^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E6^%^259C^%^25A8^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25227^%^2522^%^257D^%^255D; hy_so_1=^%^255B^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E8^%^2592^%^258B^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252234^%^252FKOKORNKOCQXVILXVB^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522ji^%^25C7^%^258Eng^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E8^%^2589^%^25B9^%^2522^%^252C^%^2522num^%^2522^%^253A^%^252212^%^2522^%^257D^%^252C^%^257B^%^2522zi^%^2522^%^253A^%^2522^%^25E6^%^259D^%^258E^%^2522^%^252C^%^2522url^%^2522^%^253A^%^252227^%^252FPWTBILILTBTBMEILE^%^252F^%^2522^%^252C^%^2522py^%^2522^%^253A^%^2522l^%^25C7^%^2590^%^252C^%^2522^%^252C^%^2522bushou^%^2522^%^253A^%^2522^%^25E6^%^259C^%^25A8^%^2522^%^252C^%^2522num^%^2522^%^253A^%^25227^%^2522^%^257D^%^255D; ASP.NET_SessionId=2u2rhxdtt3hux2ujokuiemjw")
	req.Header.Set("Origin", "http://hy.httpcn.com")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7,zh-TW;q=0.6")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "http://hy.httpcn.com/bihua/kangxi/")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")

	return req, nil

}

func KangXiSoRequest(wd string, qb string) (*http.Request, error) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	body := strings.NewReader(`wd=` + url.QueryEscape(wd))
	req, err := http.NewRequest("POST", "http://hy.httpcn.com/so/kangxi/", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Origin", "http://hy.httpcn.com")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "http://hy.httpcn.com/so/kangxi/")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7,zh-TW;q=0.6")
	req.Header.Set("Cookie", "UM_distinctid=16c2efb5e9e134-0cfc801ee6ae06-353166-1fa400-16c2efb5e9f3c8; Hm_lvt_cd7ed86134e0e3138a7cf1994e6966c8=1564156322,1565017408; CNZZDATA1267010321=1299014713-1564151968-%7C1565183518; hy_so_4=%255B%257B%2522zi%2522%253A%2522%25E6%25AC%25A3%2522%252C%2522url%2522%253A%252228%252FPWMEILPWMETBBPWKO%252F%2522%252C%2522py%2522%253A%2522x%25C4%25ABn%252C%2522%252C%2522bushou%2522%253A%2522%25E6%25AC%25A0%2522%252C%2522num%2522%253A%25228%2522%257D%252C%257B%2522zi%2522%253A%2522%25E8%258A%2582%2522%252C%2522url%2522%253A%252234%252FKOKOILCQAZXVPWXVPW%252F%2522%252C%2522py%2522%253A%2522ji%25C3%25A9%252Cji%25C4%2593%252C%2522%252C%2522bushou%2522%253A%2522%25E8%2589%25B9%2522%252C%2522num%2522%253A%252215%2522%257D%252C%257B%2522zi%2522%253A%2522%25E7%2591%2580%2522%252C%2522url%2522%253A%252230%252FPWRNMETBAZMEILILAZ%252F%2522%252C%2522py%2522%253A%2522y%25C7%2594%252C%2522%252C%2522bushou%2522%253A%2522%25E7%258E%258B%2522%252C%2522num%2522%253A%252214%2522%257D%252C%257B%2522zi%2522%253A%2522%25E9%2595%259C%2522%252C%2522url%2522%253A%252239%252FKOXVPWKOTBRNUYUYC%252F%2522%252C%2522py%2522%253A%2522j%25C3%25ACng%252C%2522%252C%2522bushou%2522%253A%2522%25E9%2592%2585%2522%252C%2522num%2522%253A%252219%2522%257D%255D; ASP.NET_SessionId=35wig01yof2qfumenvbxtqdu")

	return req, nil
}
