package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"encoding/json"

	"go.uber.org/zap"
	"carp.cn/whale/zaplogger"
)

func NewHttpClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
	}

	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: tr,
	}
}

func PostJSON(url string, data interface{}, result interface{}) error {
	zaplogger.Debug("post:", zap.String("url:", url), zap.Reflect(" data:", data))
	dataBytes, err := json.Marshal(&data)
	if err != nil {
		zaplogger.Error("at PostJSON  json Marshal.", zap.String("error :", err.Error()))
		return err
	}

	request, _ := http.NewRequest("POST", url, bytes.NewReader(dataBytes))
	request.Header.Add("Content-Type", "application/json")
	httpClient := NewHttpClient()
	resp, err := httpClient.Do(request)
	if err != nil {
		zaplogger.Error("at PostJSON resp error.", zap.Error(err))
		return err
	}
	return getStructResponse(resp, result)
}

func getResponse(resp *http.Response) []byte {
	defer func() {
		if resp.Body != nil {
			if err := resp.Body.Close(); err != nil {
				zaplogger.Error("getResponse close body error.",
					zap.String("Error:", err.Error()))
			}
		}
	}()

	var result []byte
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ := gzip.NewReader(resp.Body)
		for {
			buf := make([]byte, 1024)
			n, err := reader.Read(buf)

			if err != nil && err != io.EOF {
				panic(err)
			}

			if n == 0 {
				break
			}
			result = append(result, buf...)
		}
	default:
		result, _ = ioutil.ReadAll(resp.Body)
	}

	zaplogger.Debug("http",
		zap.Int("resp:", resp.StatusCode), zap.String(":", string(result)))
	return result
}

func getStructResponse(resp *http.Response, value interface{}) error {
	result := getResponse(resp)

	if value == nil {
		return nil
	}

	err := json.Unmarshal(result, &value)
	if err != nil {
		return err
	}

	return nil
}

func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("Can't find a ip addr.")
}
