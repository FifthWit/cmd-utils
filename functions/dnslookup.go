package functions

import (
	"fmt"
	"net"
	"strings"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func GetIPs(domain string) (string, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return "", fmt.Errorf("could not get IPs: %v", err)
	}

	var result string
	for _, ip := range ips {
		result += ip.String() + "\n"
	}

	return result, nil
}

func WhoIs(domainOrIP string) (string, error) {

	domainOrIP = strings.TrimPrefix(domainOrIP, "www.")

	raw, err := whois.Whois(domainOrIP)
	if err != nil {
		return "", fmt.Errorf("could not perform WHOIS lookup: %v", err)
	}

	parsed, err := whoisparser.Parse(raw)
	if err != nil {
		return "", fmt.Errorf("could not parse WHOIS data: %v", err)
	}

	if parsed.Domain.Domain == "" {
		return "", fmt.Errorf("domain information not found in WHOIS data")
	}

	return fmt.Sprintf("Domain: %s\nRegistrar: %s\nCreation Date: %s\nExpiration Date: %s\n",
		parsed.Domain.Domain, parsed.Registrar.Name, parsed.Domain.CreatedDate, parsed.Domain.ExpirationDate), nil
}
