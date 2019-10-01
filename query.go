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

type RequestType int
type RequestFunc func(wd string) (*http.Request, error)

const (
	RequestTypeHanChengBushou RequestType = iota
	RequestTypeHanChengPinyin
	RequestTypeKangXiBushou
	RequestTypeKangXiPinyin
	RequestTypeDummy
)

var requestList = []RequestFunc{

	RequestTypeHanChengBushou: HanChengBushouRequest,
	RequestTypeHanChengPinyin: HanChengPinyinRequest,
	RequestTypeKangXiBushou:   KangXiBushouRequest,
	RequestTypeKangXiPinyin:   KangXiPinyinRequest,
	RequestTypeDummy:          DummyRequest,
}

type Query struct {
	req         *http.Request
	cache       *net.Cache
	requestType RequestType
}

type QueryOptions func(query *Query)

func NewQuery(ops ...QueryOptions) *Query {
	q := &Query{
		req:         nil,
		cache:       nil,
		requestType: RequestTypeDummy,
	}

	for _, op := range ops {
		op(q)
	}
	return q
}

func RequestTypeOption(requestType RequestType) QueryOptions {
	return func(query *Query) {
		query.requestType = requestType
	}
}

func (q *Query) SetRequestType(requestType RequestType) {
	q.requestType = requestType
}

func CacheOption(cache *net.Cache) QueryOptions {
	return func(query *Query) {
		query.cache = cache
	}
}

func (q *Query) SetCache(cache *net.Cache) {
	q.cache = cache
}

func (q *Query) Grab(wd string) (reader io.ReadCloser, err error) {
	request, err := requestList[q.requestType](wd)
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

func DummyRequest(wd string) (*http.Request, error) {
	log.With(wd).Info("dummy")
	return nil, errors.New("dummy request call")
}

func HanChengBushouRequest(wd string) (*http.Request, error) {
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

func HanChengPinyinRequest(wd string) (*http.Request, error) {
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

func KangXiBushouRequest(wd string) (*http.Request, error) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	body := strings.NewReader(`wd=` + url.QueryEscape(wd))
	req, err := http.NewRequest("POST", "http://hy.httpcn.com/bushou/zi/", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", "hy_so_4=%255B%257B%2522zi%2522%253A%2522%25E8%2592%258B%2522%252C%2522url%2522%253A%252234%252FKOKORNKOCQXVILXVB%252F%2522%252C%2522py%2522%253A%2522ji%25C7%258Eng%252C%2522%252C%2522bushou%2522%253A%2522%25E8%2589%25B9%2522%252C%2522num%2522%253A%252217%2522%257D%255D; ASP.NET_SessionId=zilmx52mwtr3xsq5i212pd5a; UM_distinctid=16c2efb5e9e134-0cfc801ee6ae06-353166-1fa400-16c2efb5e9f3c8; CNZZDATA1267010321=1299014713-1564151968-%7C1564151968; Hm_lvt_cd7ed86134e0e3138a7cf1994e6966c8=1564156322; Hm_lpvt_cd7ed86134e0e3138a7cf1994e6966c8=1564156322")
	req.Header.Set("Origin", "http://hy.httpcn.com")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7,zh-TW;q=0.6")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Mobile Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "http://hy.httpcn.com/bushou/kangxi/")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	return req, nil
}

func KangXiPinyinRequest(wd string) (*http.Request, error) {
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
