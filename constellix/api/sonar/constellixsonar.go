package sonar

import (
	"os"
	"constellix.com/constellix/api"
)

type ConstellixSonar struct {
	Contacts		Contacts
	Agents			Agents
	HTTPChecks 		HTTPChecks
	TCPChecks		TCPChecks
	DNSChecks		DNSChecks
}

func Init(apiKey, secretKey string) *ConstellixSonar {
	if apiKey == "" && secretKey == "" {
		apiKey = os.Getenv("CONSTELLIX_APIKEY")
		secretKey = os.Getenv("CONSTELLIX_APISECRET")
	}
	constellixSonar := &ConstellixSonar{}
	_ = api.GetSonarApiClient(apiKey, secretKey)
	return constellixSonar
}
