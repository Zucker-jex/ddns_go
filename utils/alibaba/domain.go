package alibaba

import (
	domain20180129 "github.com/alibabacloud-go/domain-20180129/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func GetAllDomainList() (*[]string, error) {
	var domainList []string

	var currentPageNum int32 = 1
	for {
		queryDomainListRequest := &domain20180129.QueryDomainListRequest{
			PageNum:  &currentPageNum,
			PageSize: tea.Int32(10),
		}
		runtime := &util.RuntimeOptions{}
		request, err := domainClient.QueryDomainListWithOptions(queryDomainListRequest, runtime)
		if err != nil {
			return nil, err
		}
		for _, domain := range request.Body.Data.Domain {
			domainList = append(domainList, *domain.DomainName)
		}
		if int(*request.Body.TotalItemNum) <= len(domainList)+1 {
			break
		}
		currentPageNum++
	}
	return &domainList, nil
}
