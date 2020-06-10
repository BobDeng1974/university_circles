package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
	"university_circles/service/common_service/utils/errcode"
	"university_circles/service/common_service/utils/logger"

	"go.uber.org/zap"
)

var log = logger.Logger

// NewError is method to new ipcc error
func NewError(code string, message string) error {
	return errors.New(message)
}

func DoRequest(method string, url string, body interface{}) (response map[string]interface{}, err error) {
	client := &http.Client{}
	client.Timeout = time.Second * 20

	var jsonBody []byte

	var req *http.Request
	if method == http.MethodGet {
		req, _ = http.NewRequest(method, url, nil)
	} else {
		if jsonBody, err = json.Marshal(body); err == nil {
			reqBuffer := bytes.NewBuffer(jsonBody)
			req, _ = http.NewRequest(method, url, reqBuffer)
		} else {
			log.Warn("json marshal failed", zap.Error(err))
			req, _ = http.NewRequest(method, url, nil)
		}
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		log.Error("ipcc api request failed")
		return
	}
	defer resp.Body.Close()

	var content []byte
	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Error("ipcc api response read failed")
		return
	}
	// fmt.Println(url, string(content))
	if string(content) == "" {
		log.Error("http response is not content", zap.Any("req", req), zap.Error(err))
		err = errcode.ErrHttpResponse
		return
	}

	if err = json.Unmarshal([]byte(content), &response); err != nil {
		log.Error("json unmarshal failed", zap.Error(err))
		return
	}

	// fmt.Println(response)

	return
}
