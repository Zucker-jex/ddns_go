package alibaba

import (
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	domain20180129 "github.com/alibabacloud-go/domain-20180129/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

var (
	domainClient *domain20180129.Client
	dnsClient *alidns20150109.Client
)

func InitClient(accessKeyId, accessKeySecret, domainEndpoint, dnsEndpoint *string) error {
	domainConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Endpoint:        domainEndpoint,
	}
	domainResult, err := domain20180129.NewClient(domainConfig)
	if err != nil {
		return err
	}
	domainClient = domainResult

	dnsConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Endpoint:        dnsEndpoint,
	}
	dnsConfig.Endpoint = tea.String("alidns.cn-hangzhou.aliyuncs.com")
	dnsResult, err := alidns20150109.NewClient(dnsConfig)
	if err != nil {
		return err
	}
	dnsClient = dnsResult
	return nil
}

func GetDomainClient() *domain20180129.Client {
	return domainClient
}

func GetDNSClient() *alidns20150109.Client {
	return dnsClient
}
