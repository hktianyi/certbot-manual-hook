package config

import (
	"flag"
	"fmt"
	"os"
)

var (
	RR = "_acme-challenge"

	Action string

	RegionId  string
	AccessKey string
	Secret    string
)

func init() {
	flag.StringVar(&RegionId, "regionId", "cn-hangzhou", "regionId, default is 'cn-hangzhou'")
	flag.StringVar(&AccessKey, "ak", "", "AccessKey")
	flag.StringVar(&Secret, "secret", "", "AccessSecret")
	flag.StringVar(&Action, "action", "auth", "[auth, clean], default is 'auth'")
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `certbot manual-hook version: nginx/1.10.0
Usage: manual-hook [-ak accessKey] [-secret AccessSecret] [-regionId regionId]

Usage with certbot: certbot certonly [--dry-run] --manual --preferred-challenges dns --manual-auth-hook "./manual-hook -ak 'ak' -secret 'secret'" --manual-cleanup-hook "./manual-hook -action clean -ak 'ak' -secret 'secret'" -d *.example.com

Options:`)
	flag.PrintDefaults()
}
