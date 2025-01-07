package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/miekg/dns"
	"github.com/semihalev/sdns/config"
	"github.com/semihalev/sdns/dnsutil"
	"github.com/semihalev/sdns/middleware/resolver"
)

func main() {
	flag.Parse()
	if flag.Arg(0) == "" {
		fmt.Println("Usage: test <domain>")
		os.Exit(1)
	}

	cfg := config.Config{
		Version:      "1.3.7",
		Directory:    "db",
		BlockLists:   []string{},
		BlockListDir: "",
		RootServers: []string{
			"198.41.0.4:53",
			"199.9.14.201:53",
			"192.33.4.12:53",
			"199.7.91.13:53",
			"192.203.230.10:53",
			"192.5.5.241:53",
			"192.112.36.4:53",
			"198.97.190.53:53",
			"192.36.148.17:53",
			"192.58.128.30:53",
			"193.0.14.129:53",
			"199.7.83.42:53",
			"202.12.27.33:53",
		},
		Root6Servers: []string{
			"[2001:503:ba3e::2:30]:53",
			"[2001:500:200::b]:53",
			"[2001:500:2::c]:53",
			"[2001:500:2d::d]:53",
			"[2001:500:a8::e]:53",
			"[2001:500:2f::f]:53",
			"[2001:500:12::d0d]:53",
			"[2001:500:1::53]:53",
			"[2001:7fe::53]:53",
			"[2001:503:c27::2:30]:53",
			"[2001:7fd::1]:53",
			"[2001:500:9f::42]:53",
			"[2001:dc3::35]:53",
		},
		DNSSEC: "on",
		RootKeys: []string{
			".			172800	IN	DNSKEY	257 3 8 AwEAAaz/tAm8yTn4Mfeh5eyI96WSVexTBAvkMgJzkKTOiW1vkIbzxeF3+/4RgWOq7HrxRixHlFlExOLAJr5emLvN7SWXgnLh4+B5xQlNVz8Og8kvArMtNROxVQuCaSnIDdD5LKyWbRd2n9WGe2R8PzgCmr3EgVLrjyBxWezF0jLHwVN8efS3rCj/EWgvIWgb9tarpVUDK/b58Da+sqqls3eNbuv7pr+eoZG+SrDK6nWeL3c6H5Apxz7LjVc1uTIdsIXxuOLYA4/ilBmSVIzuDWfdRUfhHdY6+cn8HFRm+2hM8AnXGXws9555KrUB5qihylGa8subX2Nn6UwNR1AkUTV74bU=",
		},
		FallbackServers:  []string{},
		ForwarderServers: []string{},
		AccessList: []string{
			"0.0.0.0/0",
			"::0/0",
		},
		LogLevel:       "info",
		AccessLog:      "",
		Bind:           ":53",
		BindTLS:        ":853",
		BindDOH:        ":443",
		BindDOQ:        ":853",
		TLSCertificate: "server.crt",
		TLSPrivateKey:  "server.key",
		API:            "127.0.0.1:8080",
		BearerToken:    "",
		Nullroute:      "0.0.0.0",
		Nullroutev6:    "::0",
		Hostsfile:      "",
		OutboundIPs:    []string{},
		OutboundIP6s:   []string{},
		Timeout: config.Duration{
			Duration: time.Duration(2 * time.Second),
		},
		QueryTimeout: config.Duration{
			Duration: time.Duration(10 * time.Second),
		},
		Expire:          600,
		CacheSize:       256000,
		Prefetch:        10,
		Maxdepth:        30,
		RateLimit:       0,
		ClientRateLimit: 0,
		NSID:            "",
		Blocklist:       []string{},
		Whitelist:       []string{},
		Chaos:           true,
		QnameMinLevel:   5,
		EmptyZones:      []string{},
		// Plugins:         map[string]config.Plugin{},
		// CookieSecret:    "",
		// IPv6Access:      false,
	}
	resolve := resolver.NewResolver(&cfg)

	ctx := context.Background()
	var req dns.Msg
	req.SetQuestion(dns.Fqdn(flag.Arg(0)), dns.TypeA)
	req.SetEdns0(dnsutil.DefaultMsgSize, true)
	req.CheckingDisabled = true
	var rr []dns.RR
	res, err := resolve.Resolve(ctx, &req, nil, true, 30, 5, false, rr)
	if err != nil {
		log.Fatalf("Resolve error: %v", err)
	}

	fmt.Println(res)
}
