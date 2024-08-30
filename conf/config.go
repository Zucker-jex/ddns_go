package conf

import (
	"errors"
	"fmt"
	"github.com/go-ini/ini"
	"strings"
	"time"
)

// DnsConfig
type DnsConfig struct {
	AccessKeyId *string
	AccessKeySecret *string
	DomainEndpoint *string
	DnsEndpoint *string
	DomainList *[]string
	DnsType *string
	ExecType *string
	DurationMinute *time.Duration
}

func GetConfig(path *string) (*DnsConfig, error) {
	config, err := ini.Load(*path)
	if err != nil {
		return nil, err
	}
	accessKeyId, accessKeySecret, domainEndpoint, dnsEndpoint, err := getAliyunConfig(config)
	if err != nil {
		return nil, err
	}
	domainList, dnsType, err := getDomainList(config)
	if err != nil {
		return nil, err
	}
	if strings.Compare("ipv4", *dnsType) != 0 &&
		strings.Compare("ipv6", *dnsType) != 0 {
		return nil, errors.New(fmt.Sprintf("IP地址解析类型错误，请填写ipv4或ipv6（且只能填写小写）！您填写的值为：%v", *dnsType))
	}
	execType, durationMinute, err := getDurationMinute(config)
	if err != nil {
		return nil, err
	}
	return &DnsConfig{
		AccessKeyId: accessKeyId,
		AccessKeySecret: accessKeySecret,
		DomainEndpoint: domainEndpoint,
		DnsEndpoint: dnsEndpoint,
		DomainList: domainList,
		DnsType: dnsType,
		ExecType: execType,
		DurationMinute: durationMinute,
	}, nil
}

func getAliyunConfig(config *ini.File) (*string, *string, *string, *string, error) {
	aliyunSection, err := config.GetSection("aliyun")
	if err != nil {
		return nil, nil, nil, nil, err
	}
	accessKeyIdKey, err := aliyunSection.GetKey("accessKeyId")
	if err != nil {
		return nil, nil, nil, nil, err
	}
	accessKeyId := accessKeyIdKey.Value()
	accessKeySecretKey, err := aliyunSection.GetKey("accessKeySecret")
	if err != nil {
		return nil, nil, nil, nil, err
	}
	accessKeySecret := accessKeySecretKey.Value()
	domainEndpointKey, err := aliyunSection.GetKey("domainEndpoint")
	if err != nil {
		return nil, nil, nil, nil, err
	}
	domainEndpoint := domainEndpointKey.Value()
	dnsEndpointKey, err := aliyunSection.GetKey("dnsEndpoint")
	if err != nil {
		return nil, nil, nil, nil, err
	}
	dnsEndpoint := dnsEndpointKey.Value()
	return &accessKeyId, &accessKeySecret, &domainEndpoint, &dnsEndpoint, nil
}

func getDomainList(config *ini.File) (*[]string, *string, error) {
	domainSection, err := config.GetSection("domain")
	if err != nil {
		return nil, nil, err
	}
	domainListKey, err := domainSection.GetKey("domainList")
	if err != nil {
		return nil, nil, err
	}
	domainListStr := domainListKey.Value()
	domainList := strings.Split(domainListStr, ",")
	dnsTypeKey, err := domainSection.GetKey("dnsType")
	if err != nil {
		return nil, nil, err
	}
	dnsType := dnsTypeKey.Value()
	return &domainList, &dnsType, nil
}

func getDurationMinute(config *ini.File) (*string, *time.Duration, error) {
	timeSection, err := config.GetSection("time")
	if err != nil {
		return nil, nil, err
	}
	execTypeKey, err := timeSection.GetKey("type")
	if err != nil {
		return nil, nil, err
	}
	execType := execTypeKey.String()
	var duration time.Duration
	if strings.Compare(execType, "repetition") == 0 {
		durationMinuteKey, err := timeSection.GetKey("durationMinute")
		if err != nil {
			return nil, nil, err
		}
		durationMinute, err := durationMinuteKey.Int64()
		if err != nil {
			return nil, nil, err
		}
		duration = time.Duration(durationMinute)
	}
	return &execType, &duration, err
}
