package sonar

import (
	"encoding/json"
)

type MonitorAgent struct {
	AgentId			string		`json:"agentId,omitempty"`
	SiteId			int			`json:"siteId,omitempty"`
	Status			string		`json:"status,omitempty"`
}

type HTTPTestResult struct {
	Status				string			`json:"status,omitempty"`
	ResponseTime		float32			`json:"responseTime,omitempty"`
	DnsLookUpTime		float32			`json:"dnsLookUpTime,omitempty"`
	ResolvedIpAddress	string			`json:"resolvedIpAddress,omitempty"`
	StatusCode			int				`json:"statusCode,omitempty"`
	MonitorAgent		MonitorAgent	`json:"monitorAgent,omitempty"`
	Message				string			`json:"message,omitempty"`
}

func (d *HTTPTestResult) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	return nil
}

type TraceTestResult struct {
	Agent		MonitorAgent	`json:"agent,omitempty"`
	Result		string			`json:"result,omitempty"`
}

func (d *TraceTestResult) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	return nil
}

type TCPTestResult struct {
	Status				string			`json:"status,omitempty"`
	ResponseTime		float32			`json:"responseTime,omitempty"`
	DnsLookUpTime		float32			`json:"dnsLookUpTime,omitempty"`
	ResolvedIpAddress	string			`json:"resolvedIpAddress,omitempty"`
	MonitorAgent		MonitorAgent	`json:"monitorAgent,omitempty"`
	Message				string			`json:"message,omitempty"`
}

func (d *TCPTestResult) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	return nil
}

type DNSTestResult struct {
	Status				string			`json:"status,omitempty"`
	ResponseTime		float32			`json:"responseTime,omitempty"`
	MonitorAgent		MonitorAgent	`json:"monitorAgent,omitempty"`
	Message				string			`json:"message,omitempty"`
}

func (d *DNSTestResult) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	return nil
}
