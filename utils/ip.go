package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type WanIpResponse struct {
	Ip string `json:"ip"`
}

func GetWanIpAddress(dnsType *string) (*string, error) {
	var requestURL string
	requestURL = "https://" + *dnsType + ".jsonip.com"
	resp, err := http.Get(requestURL)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("获取IP地址的时候发生错误:%v", err)
		}
	}(resp.Body)
	if err != nil {
		return nil, err
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ipResponse = &WanIpResponse{}
	err = json.Unmarshal(result, ipResponse)
	if err != nil {
		return nil, err
	}
	return &ipResponse.Ip, nil
}
