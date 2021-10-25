package dns

import (
	"container/list"
	"github.com/tidwall/gjson"
	"encoding/json"
	"fmt"
	"constellix.com/constellix/api"
)

type Tag struct {
	apiClient *api.ApiClient
	Id 					int			`json:"id,omitempty"`
	Name 				string		`json:"name,omitempty"`
}

type TagParam struct {
	Name 				string		`json:"name,omitempty"`
}

func (d *Tag) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	return nil
}

func (d *Tag) Update(param TagParam) (*Tag, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var paramJson string = string(resParam)

	var jsonData, err1 = d.apiClient.DoPut(fmt.Sprintf("tags/%d", d.Id), []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return nil, err1
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var tag Tag

	err2 := tag.parse(dataValue.String())
	if err2 != nil {
		return nil, err2
	}

	return &tag, nil
}

func (d *Tag) Delete() error {
	d.apiClient = api.GetDnsApiClient("", "")

	var _, err = d.apiClient.DoDelete(fmt.Sprintf("tags/%d", d.Id), api.CLIENTTYPE_DNS)
	if err != nil {
		return err
	}

	return nil
}

type Tags struct {
	apiClient *api.ApiClient
}

func (d *Tags) GetAll() (*list.List, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	tags := list.New()

	var currentPage int
	for true {
		var url string
		if currentPage == 0 {
			urlP := &url
			*urlP = "tags"
		} else {
			urlP := &url
			*urlP = fmt.Sprintf("tags?page=%d", currentPage)
		}

		var jsonData, err = d.apiClient.DoGet(url, api.CLIENTTYPE_DNS)
		if err != nil {
			return nil, err
		}

		len := gjson.Get(string(jsonData), "data.#")
		for i := int64(0); i < len.Int(); i++ {
			tagJson := gjson.Get(string(jsonData), fmt.Sprintf("data.%d", i))

			var tag Tag
			err := tag.parse(tagJson.String())
			if err != nil {
				return nil, err
			}

			tags.PushBack(tag)
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
	return tags, nil
}

func (d *Tags) GetTag(id int) (*Tag, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("tags/%d", id), api.CLIENTTYPE_DNS)
	if err != nil {
		return nil, err
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var tag Tag

	err1 := tag.parse(dataValue.String())
	if err1 != nil {
		return nil, err1
	}

	return &tag, nil
}

func (d *Tags) Create(param TagParam) (int, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return 0, err
	}

	var paramJson string = string(resParam)

	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err1 = d.apiClient.DoPost("tags", []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return 0, err1
	}

	id := gjson.Get(string(jsonData), "data.id")

	return int(id.Int()), nil
}