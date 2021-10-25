package dns

import (
	"container/list"
	"github.com/tidwall/gjson"
	"encoding/json"
	"fmt"
	"constellix.com/constellix/api"
)

type DomainSnapshot struct {
	apiClient *api.ApiClient
	domainId int
	Id 					int			`json:"id,omitempty"`
	Name 				string		`json:"name,omitempty"`
	Version 			int			`json:"version,omitempty"`
	UpdatedAt			string		`json:"updatedAt,omitempty"`
}

func (d *DomainSnapshot) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	return nil
}

type DomainHistory struct {
	apiClient *api.ApiClient
	domainId int
}

func (d *DomainHistory) GetHistory() (*list.List, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	domainHistory := list.New()

	var currentPage int
	for true {
		var url string
		if currentPage == 0 {
			urlP := &url
			*urlP = fmt.Sprintf("domains/%d/history", d.domainId)
		} else {
			urlP := &url
			*urlP = fmt.Sprintf("domains/%d/history?page=%d", d.domainId, currentPage)
		}

		var jsonData, err = d.apiClient.DoGet(url, api.CLIENTTYPE_DNS)
		if err != nil {
			return nil, err
		}

		len := gjson.Get(string(jsonData), "data.#")
		for i := int64(0); i < len.Int(); i++ {
			domainSnapshotJson := gjson.Get(string(jsonData), fmt.Sprintf("data.%d", i))

			var domainSnapshot DomainSnapshot
			err := domainSnapshot.parse(domainSnapshotJson.String())
			if err != nil {
				return nil, err
			}

			domainHistory.PushBack(domainSnapshot)
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
	return domainHistory, nil
}

func (d *DomainHistory) GetHistoryVersion(version int) (*DomainSnapshot, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("domains/%d/history/%d", d.domainId, version), api.CLIENTTYPE_DNS)
	if err != nil {
		return nil, err
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var domainSnapshot DomainSnapshot

	err1 := domainSnapshot.parse(dataValue.String())
	if err1 != nil {
		return nil, err1
	}

	return &domainSnapshot, nil
}

func (d *DomainHistory) Apply(version int) (error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var _, err = d.apiClient.DoPost(fmt.Sprintf("domains/%d/history/%d/apply", d.domainId, version), nil, api.CLIENTTYPE_DNS)
	if err != nil {
		return err
	}

	return nil
}

func (d *DomainHistory) Snapshot(version int) (int, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err = d.apiClient.DoPost(fmt.Sprintf("domains/%d/history/%d/snapshot", d.domainId, version), nil, api.CLIENTTYPE_DNS)
	if err != nil {
		return 0, err
	}

	v := gjson.Get(string(jsonData), "data.version")

	return int(v.Int()), nil
}