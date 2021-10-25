package dns

import (
	"container/list"
	"github.com/tidwall/gjson"
	"encoding/json"
	"fmt"
	"constellix.com/constellix/api"
)

type NameserverGroup struct {
	Id 			int 		`json:"id,omitempty"`
	Name        string 		`json:"name,omitempty"`
}

type VanityNameserver struct {
	apiClient *api.ApiClient
	Id 					int					`json:"id,omitempty"`
	Name 				string				`json:"name,omitempty"`
	Default 			bool				`json:"default,omitempty"`
	Public 				bool				`json:"public,omitempty"`
	NameserverGroup 	NameserverGroup		`json:"nameserverGroup,omitempty"`
	Nameservers 		[]string			`json:"nameservers,omitempty"`
}

type VanityNameserverParam struct {
	Name 				string				`json:"name,omitempty"`
	Default 			bool				`json:"default,omitempty"`
	NameserverGroup 	NameserverGroup		`json:"nameserverGroup,omitempty"`
	Nameservers 		[]string			`json:"nameservers,omitempty"`
}

func (d *VanityNameserver) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	return nil
}

func (d *VanityNameserver) Update(param VanityNameserverParam) (*VanityNameserver, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var paramJson string = string(resParam)

	var jsonData, err1 = d.apiClient.DoPut(fmt.Sprintf("vanitynameservers/%d", d.Id), []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return nil, err1
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var vanityNameserver VanityNameserver

	err2 := vanityNameserver.parse(dataValue.String())
	if err2 != nil {
		return nil, err2
	}

	return &vanityNameserver, nil
}

func (d *VanityNameserver) Delete() error {
	d.apiClient = api.GetDnsApiClient("", "")

	var _, err = d.apiClient.DoDelete(fmt.Sprintf("vanitynameservers/%d", d.Id), api.CLIENTTYPE_DNS)
	if err != nil {
		return err
	}

	return nil
}

type VanityNameservers struct {
	apiClient *api.ApiClient
}

func (d *VanityNameservers) GetAll() (*list.List, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	vanityNameservers := list.New()

	var currentPage int
	for true {
		var url string
		if currentPage == 0 {
			urlP := &url
			*urlP = "vanitynameservers"
		} else {
			urlP := &url
			*urlP = fmt.Sprintf("vanitynameservers?page=%d", currentPage)
		}

		var jsonData, err = d.apiClient.DoGet(url, api.CLIENTTYPE_DNS)
		if err != nil {
			return nil, err
		}

		len := gjson.Get(string(jsonData), "data.#")
		for i := int64(0); i < len.Int(); i++ {
			vanityNameserverJson := gjson.Get(string(jsonData), fmt.Sprintf("data.%d", i))

			var vanityNameserver VanityNameserver
			err := vanityNameserver.parse(vanityNameserverJson.String())
			if err != nil {
				return nil, err
			}

			vanityNameservers.PushBack(vanityNameserver)
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
	return vanityNameservers, nil
}

func (d *VanityNameservers) GetVanityNameserver(id int) (*VanityNameserver, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("vanitynameservers/%d", id), api.CLIENTTYPE_DNS)
	if err != nil {
		return nil, err
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var vanityNameserver VanityNameserver

	err1 := vanityNameserver.parse(dataValue.String())
	if err1 != nil {
		return nil, err1
	}

	return &vanityNameserver, nil
}

func (d *VanityNameservers) Create(param VanityNameserverParam) (int, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return 0, err
	}

	var paramJson string = string(resParam)

	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err1 = d.apiClient.DoPost("vanitynameservers", []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return 0, err1
	}

	id := gjson.Get(string(jsonData), "data.id")

	return int(id.Int()), nil
}
