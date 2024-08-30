package service

import (
	"fmt"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	"github.com/Zucker-jex/go_ddns/utils"
	"github.com/Zucker-jex/go_ddns/utils/alibaba"
	"log"
	"strings"
)

func SyncAllDomain(domainNameList *[]string, dnsType *string) error {
	wanIp, err := utils.GetWanIpAddress(dnsType)
	if err != nil {
		return err
	}
	log.Printf("成功获取当前的公网IP地址：%v\n", *wanIp)
	domainList, err := alibaba.GetAllDomainList()
	if err != nil {
		return err
	}
	for _, domainName := range *domainNameList {
		log.Printf("开始尝试同步域名：%v\n", domainName)
		var level2Domain string
		var rr string
		for _, domain := range *domainList {
			if len(domain) > len(domainName) {
				continue
			}
			if strings.HasSuffix(domainName, fmt.Sprintf(".%v", domain)) || strings.Compare(domainName, domain) == 0 {
				level2Domain = domain
				if strings.Compare(domainName, domain) == 0 {
					rr = "@"
				} else {
					rr = strings.TrimSuffix(domainName, fmt.Sprintf(".%v", domain))
				}
				break
			}
		}
		if strings.Compare(level2Domain, "") == 0 || strings.Compare(rr, "") == 0 {
			log.Printf("非常抱歉域名%v可能不属于您，请您确认你的阿里云账户的域名信息！\n", domainName)
			continue
		}
		log.Printf("成功查询到%v域名信息信息，二级域名：%v，rr值：%v\n", domainName, level2Domain, rr)
		dnsList, err := alibaba.GetAllDNSListByDomainNameAndRR(&level2Domain, &rr)
		if err != nil {
			log.Printf("查询%v域名解析记录时候发生错误，错误信息：%v，将继续同步下一个域名\n", domainName, err)
			continue
		}
		var targetRecord *alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord = nil
		// 判断记录类型是否存在
		for _, record := range *dnsList {
			if strings.Compare(*record.Type, "A") == 0 ||
				strings.Compare(*record.Type, "AAAA") == 0 ||
				strings.Compare(*record.Type, "CNAME") == 0 ||
				strings.Compare(*record.Type, "TXT") == 0 {
				targetRecord = record
				break
			}
		}
		if targetRecord == nil {
			err = alibaba.AddDNSRecord(&level2Domain, &rr, wanIp, dnsType)
			if err != nil {
				log.Printf("新增%v解析的时候发生错误，错误信息：%v\n", domainName, err)
			} else {
				log.Printf("新增%v解析记录成功，解析到IP地址为%v\n", domainName, *wanIp)
			}
		} else if strings.Compare(*targetRecord.Type, *alibaba.GetDNSType(dnsType)) != 0 ||
			strings.Compare(*targetRecord.Value, *wanIp) != 0 ||
			strings.Compare(*targetRecord.Line, "default") != 0 ||
			*targetRecord.TTL != 600 {
			err = alibaba.UpdateDNSRecord(targetRecord.RecordId, &rr, wanIp, dnsType)
			if err != nil {
				log.Printf("修改%v解析的时候发生错误，错误信息：%v\n", domainName, err)
			} else {
				log.Printf("修改%v解析记录成功，解析到IP地址：%v，原类型：%v，原记录值：%v\n", domainName, *wanIp, *targetRecord.Type, *targetRecord.Value)
			}
		} else {
			log.Printf("无需修改%v的解析记录，记录值为：%v\n", domainName, *wanIp)
		}
	}
	return nil
}
