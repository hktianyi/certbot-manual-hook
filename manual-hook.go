package main

import (
	"fmt"
	"manual-hook/config"
	aliyun "manual-hook/dns-provider"
	"os"
	"time"
)

func main() {
	certbotDomain := os.Getenv("CERTBOT_DOMAIN")
	fmt.Println("CERTBOT_DOMAIN: ", certbotDomain)

	switch config.Action {
	case "auth":
		certbotValidation := os.Getenv("CERTBOT_VALIDATION")
		fmt.Println("CERTBOT_VALIDATION: ", certbotValidation)

		response := aliyun.AddDomainRecord(certbotDomain, config.RR, certbotValidation)
		if len(response.RecordId) > 0 {
			time.Sleep(time.Second * 30)
		}
		r := aliyun.GetDomainRecordsSimple(certbotDomain)
		for _, record := range r.DomainRecords.Record {
			fmt.Println(record)
		}
		break

	case "clean":
		r := aliyun.GetDomainRecords(certbotDomain, config.RR, "")
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
