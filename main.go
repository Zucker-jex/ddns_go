package main

import (
	"github.com/alibabacloud-go/tea/tea"
	"github.com/Zucker-jex/go_ddns/conf"
	"github.com/Zucker-jex/go_ddns/service"
	"github.com/Zucker-jex/go_ddns/utils/alibaba"
	"log"
	"os"
	"strings"
	"time"
)

// https://next.api.aliyun.com/api/Domain/2018-01-29/SaveSingleTaskForCreatingOrderActivate?lang=GO
// https://next.api.aliyun.com/api/Alidns/2015-01-09/AddCustomLine?lang=GO

func main() {
	var configFilePath *string
	if len(os.Args) >= 2 {
		configFilePath = tea.String(os.Args[1])
	} else {
		configFilePath = tea.String("./conf/config.ini")
	}
	dnsConfig, err := conf.GetConfig(configFilePath)
	if err != nil {
		log.Fatalf("读取配置文件时候发生错误：%v\n", err)
	}
	err = alibaba.InitClient(dnsConfig.AccessKeyId, dnsConfig.AccessKeySecret, dnsConfig.DomainEndpoint, dnsConfig.DnsEndpoint)
	if err != nil {
		log.Fatalf("初始化阿里云域名客户端的时候发生了错误：%v\n", err)
	}
	log.Println("域名和DNS解析客户端初始化成功")
	if strings.Compare(*dnsConfig.ExecType, "repetition") == 0 {
		for {
			go _main(dnsConfig.DomainList, dnsConfig.DnsType)
			time.Sleep(*dnsConfig.DurationMinute * time.Minute)
		}
	} else if strings.Compare(*dnsConfig.ExecType, "single") == 0 {
		_main(dnsConfig.DomainList, dnsConfig.DnsType)
	} else {
		log.Fatalln("执行类型（time.type）配置错误，值只能为single（单次执行）和repetition（多次执行）")
	}
}

func _main(domainNameList *[]string, dnsType *string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("系统发生了异常：%v\n", err)
		}
	}()

	err := service.SyncAllDomain(domainNameList, dnsType)
	if err != nil {
		log.Printf("同步域名信息的时候发生了异常：%v\n", err)
	}
}
