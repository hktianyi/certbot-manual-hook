package dns_provider

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"manual-hook/config"
)

var client *alidns.Client

func NewClient(regionId string, accessKey string, secret string) *alidns.Client {
	if client == nil {
		clientTemp, err := alidns.NewClientWithAccessKey(regionId, accessKey, secret)
		if err != nil {
			fmt.Print(err.Error())
		}
		client = clientTemp
	}
	return client
}

func GetDomainRecordsSimple(domainName string) *alidns.DescribeDomainRecordsResponse {
	return GetDomainRecords(domainName, "", "")
}

func GetDomainRecords(domainName string, rrKeyWord string, valueKeyWord string) *alidns.DescribeDomainRecordsResponse {
	if client == nil {
		NewClient(config.RegionId, config.AccessKey, config.Secret)
	}

	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = domainName
	request.RRKeyWord = rrKeyWord
	request.ValueKeyWord = valueKeyWord

	response, err := client.DescribeDomainRecords(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("GetDomainRecords response is %#v\n", response)
	return response
}

func AddDomainRecord(domainName string, rr string, value string) *alidns.AddDomainRecordResponse {
	if client == nil {
		NewClient(config.RegionId, config.AccessKey, config.Secret)
	}

	request := alidns.CreateAddDomainRecordRequest()
	request.Type = "TXT"
	request.DomainName = domainName
	request.RR = rr
	request.Value = value

	response, err := client.AddDomainRecord(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("AddDomainRecord response is %#v\n", response)
	return response
}

func DeleteDomainRecord(recordId string) *alidns.DeleteDomainRecordResponse {
	if client == nil {
		NewClient(config.RegionId, config.AccessKey, config.Secret)
	}

	request := alidns.CreateDeleteDomainRecordRequest()
	request.RecordId = recordId

	response, err := client.DeleteDomainRecord(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("DeleteDomainRecord response is %#v\n", response)
	return response
}
