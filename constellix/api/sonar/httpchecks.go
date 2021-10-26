package sonar

import (
	"container/list"
	"github.com/tidwall/gjson"
	"encoding/json"
	"fmt"
	"constellix.com/constellix/api"
	"strconv"
	"strings"
)

type HTTPCheck struct {
	apiClient *api.ApiClient
	Id 							int						`json:"id,omitempty"`
	Name 						string					`json:"name,omitempty"`
	Host 						string					`json:"host,omitempty"`
	Port 						int						`json:"port,omitempty"`
	ProtocolType 				ProtocolType			`json:"protocolType,omitempty"`
	IpVersion					IPVersion				`json:"ipVersion,omitempty"`
	Fqdn 						string					`json:"fqdn,omitempty"`
	Path 						string					`json:"path,omitempty"`
	SearchString 				string					`json:"searchString,omitempty"`
	ConnectionTimeout 			int						`json:"connectionTimeout,omitempty"`
	ExpectedStatusCode 			int						`json:"expectedStatusCode,omitempty"`
	UserAgent 					string					`json:"userAgent,omitempty"`
	Note 						string					`json:"note,omitempty"`
	ScheduleInterval 			ScheduleInterval		`json:"scheduleInterval,omitempty"`
	UserId 						int						`json:"userId,omitempty"`
	Interval 					MonitoringInterval		`json:"interval,omitempty"`
	MonitorIntervalPolicy 		MonitorIntervalPolicy	`json:"monitorIntervalPolicy,omitempty"`
	CheckSites 					[]int					`json:"checkSites,omitempty"`
	NotificationGroups 			[]int					`json:"notificationGroups,omitempty"`
	ScheduleId 					int						`json:"scheduleId,omitempty"`
	NotificationReportTimeout 	int						`json:"notificationReportTimeout,omitempty"`
	VerificationPolicy 			VerificationPolicy		`json:"verificationPolicy,omitempty"`
	RunTraceroute	 			RunTraceroute			`json:"runTraceroute,omitempty"`
}

type HTTPCheckParam struct {
	Name 						string					`json:"name,omitempty"`
	Host 						string					`json:"host,omitempty"`
	Port 						int						`json:"port,omitempty"`
	ProtocolType 				ProtocolType			`json:"protocolType,omitempty"`
	IpVersion					IPVersion				`json:"ipVersion,omitempty"`
	Fqdn 						string					`json:"fqdn,omitempty"`
	Path 						string					`json:"path,omitempty"`
	SearchString 				string					`json:"searchString,omitempty"`
	ConnectionTimeout 			int						`json:"connectionTimeout,omitempty"`
	ExpectedStatusCode 			int						`json:"expectedStatusCode,omitempty"`
	UserAgent 					string					`json:"userAgent,omitempty"`
	Note 						string					`json:"note,omitempty"`
	ScheduleInterval 			ScheduleInterval		`json:"scheduleInterval,omitempty"`
	UserId 						int						`json:"userId,omitempty"`
	Interval 					MonitoringInterval		`json:"interval,omitempty"`
	MonitorIntervalPolicy 		MonitorIntervalPolicy	`json:"monitorIntervalPolicy,omitempty"`
	CheckSites 					[]int					`json:"checkSites,omitempty"`
	NotificationGroups 			[]int					`json:"notificationGroups,omitempty"`
	ScheduleId 					int						`json:"scheduleId,omitempty"`
	NotificationReportTimeout 	int						`json:"notificationReportTimeout,omitempty"`
	VerificationPolicy 			VerificationPolicy		`json:"verificationPolicy,omitempty"`
	RunTraceroute	 			RunTraceroute			`json:"runTraceroute,omitempty"`
}

func (d *HTTPCheck) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	return nil
}

func (d *HTTPCheck) Update(param HTTPCheckParam) (*HTTPCheck, error) {
	d.apiClient = api.GetSonarApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var paramJson string = string(resParam)

	var jsonData, err1 = d.apiClient.DoPut(fmt.Sprintf("http/%d", d.Id), []byte(paramJson), api.CLIENTTYPE_SONAR)
	if err1 != nil {
		return nil, err1
	}

	dataValue := string(jsonData)
	var check HTTPCheck

	err2 := check.parse(dataValue)
	if err2 != nil {
		return nil, err2
	}

	return &check, nil
}

func (d *HTTPCheck) Delete() error {
	d.apiClient = api.GetSonarApiClient("", "")

	var _, err = d.apiClient.DoDelete(fmt.Sprintf("http/%d", d.Id), api.CLIENTTYPE_SONAR)
	if err != nil {
		return err
	}

	return nil
}

func (d *HTTPCheck) RunCheck(agents []int) (*list.List, error) {
	var siteIds string = ""
	for _, id := range agents {
		if len(siteIds) == 0 {
			siteIds += "?siteIds=" + strconv.Itoa(id)
		} else {
			siteIds += "&siteIds=" + strconv.Itoa(id)
		}
	}

	d.apiClient = api.GetSonarApiClient("", "")
	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("http/%d/test%s", d.Id, siteIds), api.CLIENTTYPE_SONAR)
	if err != nil {
		return nil, err
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

func (d *HTTPCheck) RunCheckTrace(agents []int) (*list.List, error) {
	var siteIds string = ""
	for _, id := range agents {
		if len(siteIds) == 0 {
			siteIds += "?siteIds=" + strconv.Itoa(id)
		} else {
			siteIds += "&siteIds=" + strconv.Itoa(id)
		}
	}

	d.apiClient = api.GetSonarApiClient("", "")
	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("http/%d/trace%s", d.Id, siteIds), api.CLIENTTYPE_SONAR)
	if err != nil {
		return nil, err
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

func (d *HTTPCheck) Start() error {
	d.apiClient = api.GetSonarApiClient("", "")
	var _, err = d.apiClient.DoPut(fmt.Sprintf("http/%d/start", d.Id), nil, api.CLIENTTYPE_SONAR)
	if err != nil {
		return err
	}
	return nil
}

func (d *HTTPCheck) Stop() error {
	d.apiClient = api.GetSonarApiClient("", "")
	var _, err = d.apiClient.DoPut(fmt.Sprintf("http/%d/stop", d.Id), nil, api.CLIENTTYPE_SONAR)
	if err != nil {
		return err
	}
	return nil
}

func (d *HTTPCheck) CheckStatus() (string, error) {
	d.apiClient = api.GetSonarApiClient("", "")
	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("http/%d/status", d.Id), api.CLIENTTYPE_SONAR)
	if err != nil {
		return "", err
	}

	status := gjson.Get(string(jsonData), "@this.status")
	return status.String(), nil
}

func (d *HTTPCheck) CheckState() (string, error) {
	d.apiClient = api.GetSonarApiClient("", "")
	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("http/%d/state", d.Id), api.CLIENTTYPE_SONAR)
	if err != nil {
		return "", err
	}

	state := gjson.Get(string(jsonData), "@this.state")
	return state.String(), nil
}

func (d *HTTPCheck) CheckAgentStatus(agents []int) (*list.List, error) {
	var siteIds string = ""
	for _, id := range agents {
		if len(siteIds) == 0 {
			siteIds += "?siteIds=" + strconv.Itoa(id)
		} else {
			siteIds += "&siteIds=" + strconv.Itoa(id)
		}
	}

	d.apiClient = api.GetSonarApiClient("", "")
	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("http/%d/site/status%s", d.Id, siteIds), api.CLIENTTYPE_SONAR)
	if err != nil {
		return nil, err
	}

	var data map[string]AgentStatus
	err1 := json.Unmarshal([]byte(jsonData), &data)
	if err1 != nil {
		return nil, err1
	}

	res := list.New()
	for k, v := range data {
		var agentStatus AgentStatus
		agentStatus.Name = k
		agentStatus.Timestamp = v.Timestamp
		agentStatus.Status = v.Status
		agentStatus.DnsLookupTime = v.DnsLookupTime
		agentStatus.ResponseTime = v.ResponseTime
		res.PushBack(agentStatus)
	}
	
	return res, nil
}

type HTTPChecks struct {
	apiClient *api.ApiClient
}

func (d *HTTPChecks) GetAll() (*list.List, error) {
	d.apiClient = api.GetSonarApiClient("", "")

	checks := list.New()

	var jsonData, err = d.apiClient.DoGet("http", api.CLIENTTYPE_SONAR)
	if err != nil {
		return nil, err
	}

	len := gjson.Get(string(jsonData), "@this.#")
	for i := int64(0); i < len.Int(); i++ {
		checkJson := gjson.Get(string(jsonData), fmt.Sprintf("@this.%d", i))

		var check HTTPCheck
		err := check.parse(checkJson.String())
		if err != nil {
			return nil, err
		}

		checks.PushBack(check)
	}

	return checks, nil
}

func (d *HTTPChecks) GetCheck(id int) (*HTTPCheck, error) {
	d.apiClient = api.GetSonarApiClient("", "")

	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("http/%d", id), api.CLIENTTYPE_SONAR)
	if err != nil {
		return nil, err
	}

	checkValue := string(jsonData)
	var check HTTPCheck

	err1 := check.parse(checkValue)
	if err1 != nil {
		return nil, err1
	}

	return &check, nil
}

func (d *HTTPChecks) Create(param HTTPCheckParam) (int, error) {
	d.apiClient = api.GetSonarApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return 0, err
	}

	var paramJson string = string(resParam)

	var _, err1 = d.apiClient.DoPost("http", []byte(paramJson), api.CLIENTTYPE_SONAR)
	if err1 != nil {
		return 0, err1
	}

	location := d.apiClient.LastResponse.Header.Get("Location")
	s := strings.Split(location, "/")

	id, err2 := strconv.Atoi(s[len(s) - 1])
	if err2 != nil {
		return 0, err2
	}

	return id, nil
}