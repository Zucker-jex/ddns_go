package alibaba

import (
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	"github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"strings"
)

func GetAllDNSListByDomainNameAndRR(domainName, rr *string) (*[]*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord, error) {
	var dnsList []*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord

	var currentPageNum int64 = 1

	for {
		describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{
			DomainName: domainName,
			PageNumber: &currentPageNum,
			PageSize:   tea.Int64(10),
			RRKeyWord:  rr,
		}
		runtime := &service.RuntimeOptions{}
		dnsResult, err := dnsClient.DescribeDomainRecordsWithOptions(describeDomainRecordsRequest, runtime)
		if err != nil {
			return nil, err
		}
		dnsList = append(dnsList, dnsResult.Body.DomainRecords.Record...)

		if int(*dnsResult.Body.TotalCount) <= len(dnsList)+1 {
			break
		}
		currentPageNum++
	}
	return &dnsList, nil
}

func AddDNSRecord(domain, rr, ipAddress *string, dnsType *string) error {
	addDomainRecordRequest := &alidns20150109.AddDomainRecordRequest{
		DomainName: domain,
		RR:         rr,
		Type:       GetDNSType(dnsType),
		Value:      ipAddress,
		TTL:        tea.Int64(600),
		Line:       tea.String("default"),
	}
	runtime := &service.RuntimeOptions{}
	_, err := dnsClient.AddDomainRecordWithOptions(addDomainRecordRequest, runtime)
	return err
}

func UpdateDNSRecord(recordId, rr, ipAddress *string, dnsType *string) error {
	updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
		RecordId: recordId,
		RR:       rr,
		Type:     GetDNSType(dnsType),
		Value:    ipAddress,
		TTL:      tea.Int64(600),
		Line:     tea.String("default"),
	}
	runtime := &service.RuntimeOptions{}
	_, err := dnsClient.UpdateDomainRecordWithOptions(updateDomainRecordRequest, runtime)
	return err
}

func GetDNSType(dnsType *string) *string {
	var recordType *string = tea.String("A")
	if strings.Compare("ipv6", *dnsType) == 0 {
		recordType = tea.String("AAAA")
	}
	return recordType
}
