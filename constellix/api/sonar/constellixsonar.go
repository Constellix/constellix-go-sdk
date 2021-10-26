package sonar

import (
	"os"
	"constellix.com/constellix/api"
	"github.com/tidwall/gjson"
	"fmt"
)

type ConstellixSonar struct {
	Contacts		Contacts
	Agents			Agents
	HTTPChecks 		HTTPChecks
	TCPChecks		TCPChecks
	DNSChecks		DNSChecks
	apiClient 		*api.ApiClient
}

func Init(apiKey, secretKey string) *ConstellixSonar {
	if apiKey == "" && secretKey == "" {
		apiKey = os.Getenv("CONSTELLIX_API_KEY")
		secretKey = os.Getenv("CONSTELLIX_SECRET_KEY")
	}
	constellixSonar := &ConstellixSonar{}
	constellixSonar.apiClient = api.GetSonarApiClient(apiKey, secretKey)
	return constellixSonar
}

func (d* ConstellixSonar) CheckType(id int) (string, error) {
	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("check/type/%d", id), api.CLIENTTYPE_SONAR)
	if err != nil {
		return "", err
	}

	t := gjson.Get(string(jsonData), "@this.type")
	return t.String(), nil
}
