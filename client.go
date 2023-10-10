package httpclient

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"

	"github.com/gojek/heimdall/httpclient"
	"github.com/valyala/fasthttp"
)

type ReqStruct struct {
	Req     fasthttp.Request
	Method  string
	Url     string
	Body    io.Reader
	Timeout int
	Headers http.Header
}

type ReqResponse struct {
	Response *http.Response
	Error    error
}

func (req *ReqStruct) SetHeader(headers map[string]string) *ReqStruct {
	headershttp := http.Header{}
	for key, value := range headers {
		// maybe will add function for anti spoof
		headershttp.Set(key, value)
	}
	req.Headers = headershttp
	return req
}

func (req *ReqStruct) DoRequests() *ReqResponse {
	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	httpClient := httpclient.NewClient(
		httpclient.WithHTTPClient(client),
		httpclient.WithHTTPTimeout(time.Duration(time.Duration(req.Timeout).Seconds())),
	)
	request, err := http.NewRequest(req.Method, req.Url, req.Body)
	if err != nil {
		return &ReqResponse{Error: err}
	}
	request.Header = req.Headers
	response, err := httpClient.Do(request)
	if err != nil {
		return &ReqResponse{Error: err}
	}
	return &ReqResponse{
		Response: response,
	}
}
