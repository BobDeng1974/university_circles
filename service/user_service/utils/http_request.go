package utils

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"university_circles/service/user_service/utils/errcode"
	"university_circles/service/user_service/utils/logger"

	"go.uber.org/zap"
)

var log = logger.Logger

// NewError is method to new ipcc error
func NewError(code string, message string) error {
	return errors.New(message)
}

func DoRequest(method string, url string, header map[string]string, body string) (response map[string]interface{}, err error) {
	client := &http.Client{}
	client.Timeout = time.Second * 20

	tr := &http.Transport{ //解决x509: certificate signed by unknown authority
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.Transport = tr //解决x509: certificate signed by unknown authority

	var jsonBody []byte

	var req *http.Request
	if method == http.MethodGet {
		req, _ = http.NewRequest(method, url, nil)
	} else {
		req, _ = http.NewRequest(method, url, strings.NewReader(body))
		if header != nil {
			for k, v := range header {
				req.Header.Set(k, v)
			}
		}
	}

	//req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Println("DoRequest:::::::", req, req.Header, string(jsonBody), err)
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		log.Error("http request failed")
		return
	}
	fmt.Println("DoRequest resp  :::::: ", resp, err)
	defer resp.Body.Close()

	var content []byte
	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Error("http response read failed")
		return
	}
	// fmt.Println(url, string(content))
	if string(content) == "" {
		log.Error("http response is not content", zap.Any("req", req), zap.Error(err))
		err = errcode.ErrHttpResponse
		return
	}

	if err = json.Unmarshal(content, &response); err != nil {
		log.Error("json unmarshal failed", zap.Error(err))
		return
	}

	// fmt.Println(response)

	return
}
