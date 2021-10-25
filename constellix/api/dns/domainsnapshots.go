package dns

import (
	"container/list"
	"github.com/tidwall/gjson"
	"fmt"
	"constellix.com/constellix/api"
)

func (d*DomainSnapshot) Delete() (error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var _, err = d.apiClient.DoDelete(fmt.Sprintf("domains/%d/snapshots/%d", d.domainId, d.Id), api.CLIENTTYPE_DNS)
	if err != nil {
		return err
	}

	return nil
}

type DomainSnapshots struct {
	apiClient *api.ApiClient
	domainId int
}

func (d *DomainSnapshots) GetSnapshots() (*list.List, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	domainSnapshots := list.New()

	var currentPage int
	for true {
		var url string
		if currentPage == 0 {
			urlP := &url
			*urlP = fmt.Sprintf("domains/%d/snapshots", d.domainId)
		} else {
			urlP := &url
			*urlP = fmt.Sprintf("domains/%d/snapshots?page=%d", d.domainId, currentPage)
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

			domainSnapshot.domainId = d.domainId
			domainSnapshots.PushBack(domainSnapshot)
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
	return domainSnapshots, nil
}

func (d *DomainSnapshots) GetSnapshot(version int) (*DomainSnapshot, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("domains/%d/snapshots/%d", d.domainId, version), api.CLIENTTYPE_DNS)
	if err != nil {
		return nil, err
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var domainSnapshot DomainSnapshot

	domainSnapshot.domainId = d.domainId
	err1 := domainSnapshot.parse(dataValue.String())
	if err1 != nil {
		return nil, err1
	}

	return &domainSnapshot, nil
}

func (d *DomainSnapshots) ApplySnapshot(version int) (error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var _, err = d.apiClient.DoPost(fmt.Sprintf("domains/%d/snapshots/%d/apply", d.domainId, version), nil, api.CLIENTTYPE_DNS)
	if err != nil {
		return err
	}

	return nil
}
