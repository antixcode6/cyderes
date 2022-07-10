package api

import (
	"fmt"
	"net"
)

type DnsRequest struct {
	Errors      string    `json:"Errors"`
	Nameservers []*net.NS `json:"Nameservers"`
	IPs         []net.IP  `json:"IP addrs"`
	MX          []*net.MX `json:"MX records"`
}

func QueryNet(server string) DnsRequest {
	ips, err := net.LookupIP(server)
	if err != nil {
		newError := fmt.Sprintf("Could not get IPs: %v", err)
		return DnsRequest{
			Errors: newError,
		}

	}
	nameserver, err := net.LookupNS(server)
	if err != nil {
		newError := fmt.Sprintf("Could not get nameserver(s): %v", err)
		return DnsRequest{
			Errors: newError,
		}

	}
	mxrecords, err := net.LookupMX(server)
	if err != nil {
		newError := fmt.Sprintf("could not get mx record(s): %v", err)
		return DnsRequest{
			Errors: newError,
		}

	}
	return DnsRequest{
		Nameservers: nameserver,
		IPs:         ips,
		MX:          mxrecords,
	}
}
