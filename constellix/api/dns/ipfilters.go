package dns

import (
	"container/list"
	"github.com/tidwall/gjson"
	"encoding/json"
	"fmt"
	"constellix.com/constellix/api"
)

type IpFilter struct {
	apiClient *api.ApiClient
	Id 					int				`json:"id,omitempty"`
	Name 				string			`json:"name,omitempty"`
	RulesLimit 			int				`json:"rulesLimit,omitempty"`
	Continents			[]Continent		`json:"continents,omitempty"`
	Countries 			[]string		`json:"countries,omitempty"`
	Asn 				[]int			`json:"asn,omitempty"`
	Ipv4 				[]string		`json:"ipv4,omitempty"`
	Ipv6 				[]string		`json:"ipv6,omitempty"`
}

type IpFilterParam struct {
	Name 				string			`json:"name,omitempty"`
	RulesLimit 			int				`json:"rulesLimit,omitempty"`
	Continents			[]Continent		`json:"continents,omitempty"`
	Countries 			[]string		`json:"countries,omitempty"`
	Asn 				[]int			`json:"asn,omitempty"`
	Ipv4 				[]string		`json:"ipv4,omitempty"`
	Ipv6 				[]string		`json:"ipv6,omitempty"`
}

func (d *IpFilter) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	return nil
}

func (d *IpFilter) Update(param IpFilterParam) (*IpFilter, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var paramJson string = string(resParam)

	var jsonData, err1 = d.apiClient.DoPut(fmt.Sprintf("ipfilters/%d", d.Id), []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return nil, err1
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var ipFilter IpFilter

	err2 := ipFilter.parse(dataValue.String())
	if err2 != nil {
		return nil, err2
	}

	return &ipFilter, nil
}

func (d *IpFilter) Delete() error {
	d.apiClient = api.GetDnsApiClient("", "")

	var _, err = d.apiClient.DoDelete(fmt.Sprintf("ipfilters/%d", d.Id), api.CLIENTTYPE_DNS)
	if err != nil {
		return err
	}

	return nil
}

type IpFilters struct {
	apiClient *api.ApiClient
}

func (d *IpFilters) GetAll() (*list.List, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	ipFilters := list.New()

	var currentPage int
	for true {
		var url string
		if currentPage == 0 {
			urlP := &url
			*urlP = "ipfilters"
		} else {
			urlP := &url
			*urlP = fmt.Sprintf("ipfilters?page=%d", currentPage)
		}

		var jsonData, err = d.apiClient.DoGet(url, api.CLIENTTYPE_DNS)
		if err != nil {
			return nil, err
		}

		len := gjson.Get(string(jsonData), "data.#")
		for i := int64(0); i < len.Int(); i++ {
			ipFilterJson := gjson.Get(string(jsonData), fmt.Sprintf("data.%d", i))

			var ipFilter IpFilter
			err := ipFilter.parse(ipFilterJson.String())
			if err != nil {
				return nil, err
			}

			ipFilters.PushBack(ipFilter)
		}

		// handle paging
		cPage := gjson.Get(string(jsonData), "meta.pagination.currentPage")
		tPage := gjson.Get(string(jsonData), "meta.pagination.totalPages")
		if cPage.Int() >= tPage.Int() {
			break
		}

		currentPage = int(cPage.Int())
	}
	currentPage = 0
	return ipFilters, nil
}

func (d *IpFilters) GetIpFilter(id int) (*IpFilter, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("ipfilters/%d", id), api.CLIENTTYPE_DNS)
	if err != nil {
		return nil, err
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var ipFilter IpFilter

	err1 := ipFilter.parse(dataValue.String())
	if err1 != nil {
		return nil, err1
	}

	return &ipFilter, nil
}

func (d *IpFilters) Create(param IpFilterParam) (int, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return 0, err
	}

	var paramJson string = string(resParam)

	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err1 = d.apiClient.DoPost("ipfilters", []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return 0, err1
	}

	id := gjson.Get(string(jsonData), "data.id")

	return int(id.Int()), nil
}
