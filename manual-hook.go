package main

import (
	"fmt"
	"manual-hook/config"
	aliyun "manual-hook/dns-provider"
	"os"
	"time"
)

func main() {
	CERTBOT_DOMAIN := os.Getenv("CERTBOT_DOMAIN")
	fmt.Println("CERTBOT_DOMAIN: ", CERTBOT_DOMAIN)

	switch config.Action {
	case "auth":
		CERTBOT_VALIDATION := os.Getenv("CERTBOT_VALIDATION")
		fmt.Println("CERTBOT_VALIDATION: ", CERTBOT_VALIDATION)

		response := aliyun.AddDomainRecord(CERTBOT_DOMAIN, config.RR, CERTBOT_VALIDATION)
		if len(response.RecordId) > 0 {
			time.Sleep(time.Second * 30)
		}
		r := aliyun.GetDomainRecordsSimple(CERTBOT_DOMAIN)
		for _, record := range r.DomainRecords.Record {
			fmt.Println(record)
		}
		break

	case "clean":
		r := aliyun.GetDomainRecords(CERTBOT_DOMAIN, config.RR, "")
		for _, record := range r.DomainRecords.Record {
			fmt.Println(record)
			aliyun.DeleteDomainRecord(record.RecordId)
		}
		break

	default:
		fmt.Println("Invalid Action")
		break
	}
}
