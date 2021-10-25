package sonar

import (
	"container/list"
	"github.com/tidwall/gjson"
	"encoding/json"
	"fmt"
	"constellix.com/constellix/api"
)

type Agent struct {
	apiClient *api.ApiClient
	Id				int			`json:"id,omitempty"`
	Name			string		`json:"name,omitempty"`
	Label			string		`json:"label,omitempty"`
	Location		string		`json:"location,omitempty"`
	Country			string		`json:"country,omitempty"`
	Region			string		`json:"region,omitempty"`
}

type AgentStatus struct {
	Name			string 		`json:"name,omitempty"`
	Timestamp		string		`json:"timestamp,omitempty"`
	Status			string		`json:"status,omitempty"`
	DnsLookupTime	int			`json:"dnsLookupTime,omitempty"`
	ResponseTime	int			`json:"responseTime,omitempty"`
}

func (d *Agent) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	return nil
}

type Agents struct {
	apiClient *api.ApiClient
}

func (d *Agents) GetAgents() (*list.List, error) {
	d.apiClient = api.GetSonarApiClient("", "")

	agents := list.New()

	var jsonData, err = d.apiClient.DoGet("system/sites", api.CLIENTTYPE_SONAR)
	if err != nil {
		return nil, err
	}

	len := gjson.Get(string(jsonData), "@this.#")
	for i := int64(0); i < len.Int(); i++ {
		agentJson := gjson.Get(string(jsonData), fmt.Sprintf("@this.%d", i))

		var agent Agent
		err := agent.parse(agentJson.String())
		if err != nil {
			return nil, err
		}

		agents.PushBack(agent)
	}

	return agents, nil
}

func (d *Agent) HTTPTest(param HTTPTestParam) (*list.List, error) {
	d.apiClient = api.GetSonarApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var paramJson string = string(resParam)

	var jsonData, err1 = d.apiClient.DoPost(fmt.Sprintf("test/http/%d", d.Id), []byte(paramJson), api.CLIENTTYPE_SONAR)
	if err1 != nil {
		return nil, err1
	}

	results := list.New()

	len := gjson.Get(string(jsonData), "@this.#")
	for i := int64(0); i < len.Int(); i++ {
		resultJson := gjson.Get(string(jsonData), fmt.Sprintf("@this.%d", i))

		var result HTTPTestResult
		err := result.parse(resultJson.String())
		if err != nil {
			return nil, err
		}

		results.PushBack(result)
	}

	return results, nil
}

func (d *Agent) TraceTest(param TraceTestParam) (*list.List, error) {
	d.apiClient = api.GetSonarApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var paramJson string = string(resParam)

	var jsonData, err1 = d.apiClient.DoPost(fmt.Sprintf("test/trace/%d", d.Id), []byte(paramJson), api.CLIENTTYPE_SONAR)
	if err1 != nil {
		return nil, err1
	}

	results := list.New()

	len := gjson.Get(string(jsonData), "@this.#")
	for i := int64(0); i < len.Int(); i++ {
		resultJson := gjson.Get(string(jsonData), fmt.Sprintf("@this.%d", i))

		var result TraceTestResult
		err := result.parse(resultJson.String())
		if err != nil {
			return nil, err
		}

		results.PushBack(result)
	}

	return results, nil
}

func (d *Agent) TCPTest(param TCPTestParam) (*list.List, error) {
	d.apiClient = api.GetSonarApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var paramJson string = string(resParam)

	var jsonData, err1 = d.apiClient.DoPost(fmt.Sprintf("test/tcp/%d", d.Id), []byte(paramJson), api.CLIENTTYPE_SONAR)
	if err1 != nil {
		return nil, err1
	}

	results := list.New()

	len := gjson.Get(string(jsonData), "@this.#")
	for i := int64(0); i < len.Int(); i++ {
		resultJson := gjson.Get(string(jsonData), fmt.Sprintf("@this.%d", i))

		var result TCPTestResult
		err := result.parse(resultJson.String())
		if err != nil {
			return nil, err
		}

		results.PushBack(result)
	}

	return results, nil
}

func (d *Agent) DNSTest(param DNSTestParam) (*list.List, error) {
	d.apiClient = api.GetSonarApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var paramJson string = string(resParam)

	var jsonData, err1 = d.apiClient.DoPost(fmt.Sprintf("test/dns/%d", d.Id), []byte(paramJson), api.CLIENTTYPE_SONAR)
	if err1 != nil {
		return nil, err1
	}

	results := list.New()

	len := gjson.Get(string(jsonData), "@this.#")
	for i := int64(0); i < len.Int(); i++ {
		resultJson := gjson.Get(string(jsonData), fmt.Sprintf("@this.%d", i))

		var result DNSTestResult
		err := result.parse(resultJson.String())
		if err != nil {
			return nil, err
		}

		results.PushBack(result)
	}

	return results, nil
}
