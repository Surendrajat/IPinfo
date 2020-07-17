package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type IpInfo struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Asn      struct {
		Asn    string `json:"asn"`
		Name   string `json:"name"`
		Domain string `json:"domain"`
		Route  string `json:"route"`
		Type   string `json:"type"`
	} `json:"asn"`
	Company struct {
		Name   string `json:"name"`
		Domain string `json:"domain"`
		Type   string `json:"type"`
	} `json:"company"`
	Privacy struct {
		Vpn     bool `json:"vpn"`
		Proxy   bool `json:"proxy"`
		Tor     bool `json:"tor"`
		Hosting bool `json:"hosting"`
	} `json:"privacy"`
	Abuse struct {
		Address string `json:"address"`
		Country string `json:"country"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Network string `json:"network"`
		Phone   string `json:"phone"`
	} `json:"abuse"`
	Domains struct {
		IP      string   `json:"ip"`
		Total   int      `json:"total"`
		Domains []string `json:"domains"`
	} `json:"domains"`
}

func main() {

	var ipInfo IpInfo

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://ipinfo.io/widget", nil)
	req.Header.Set("referer", "https://ipinfo.io/")
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	_ = json.NewDecoder(resp.Body).Decode(&ipInfo)

	private := false
	privacy := ""
	privacyType := map[string]bool{
		"vpn":   ipInfo.Privacy.Vpn,
		"tor": ipInfo.Privacy.Tor,
		"proxy":   ipInfo.Privacy.Proxy,
		"hosting":     ipInfo.Privacy.Hosting,
	}
	for k, v := range privacyType {
		if v {
			privacy += k + " "
			private = true
		}
	}
	if !private {
		privacy = "none"
	}

	fmt.Printf("IP	%s %s\n", ipInfo.IP, ipInfo.Hostname)
	fmt.Printf("Geo	%s %s, %s, %s %s\n", ipInfo.Loc, ipInfo.City, ipInfo.Region, ipInfo.Country, ipInfo.Postal)
	fmt.Printf("ASN	[%s %s: %s(%s)]\n", ipInfo.Asn.Asn, strings.ToUpper(ipInfo.Asn.Type), ipInfo.Asn.Name, ipInfo.Asn.Domain)
	fmt.Printf("Company	[%s: %s(%s)]\n", strings.ToUpper(ipInfo.Company.Type), ipInfo.Company.Name, ipInfo.Company.Domain)
	fmt.Printf("Privacy	%s\n", privacy)
}
