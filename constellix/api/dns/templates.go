package dns

import (
	"container/list"
	"github.com/tidwall/gjson"
	"encoding/json"
	"fmt"
	"constellix.com/constellix/api"
)

type Template struct {
	apiClient *api.ApiClient
	Id 					int			`json:"id,omitempty"`
	Name 				string		`json:"name,omitempty"`
	GeoIp 				bool		`json:"geoip,omitempty"`
	Gtd 				bool		`json:"gtd,omitempty"`
	Records				TemplateRecords 	`json:"-"`
}

type TemplateParam struct {
	Name 				string		`json:"name,omitempty"`
	GeoIp 				bool		`json:"geoip"`
	Gtd 				bool		`json:"gtd"`
}

func (d *Template) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	d.Records.templateId = d.Id
	return nil
}

func (d *Template) Update(param TemplateParam) (*Template, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var paramJson string = string(resParam)

	var jsonData, err1 = d.apiClient.DoPut(fmt.Sprintf("templates/%d", d.Id), []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return nil, err1
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var template Template

	err2 := template.parse(dataValue.String())
	if err2 != nil {
		return nil, err2
	}

	return &template, nil
}

func (d *Template) Delete() error {
	d.apiClient = api.GetDnsApiClient("", "")

	var _, err = d.apiClient.DoDelete(fmt.Sprintf("templates/%d", d.Id), api.CLIENTTYPE_DNS)
	if err != nil {
		return err
	}

	return nil
}

type Templates struct {
	apiClient *api.ApiClient
}

func (d *Templates) GetAll() (*list.List, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	templates := list.New()

	var currentPage int
	for true {
		var url string
		if currentPage == 0 {
			urlP := &url
			*urlP = "templates"
		} else {
			urlP := &url
			*urlP = fmt.Sprintf("templates?page=%d", currentPage)
		}

		var jsonData, err = d.apiClient.DoGet(url, api.CLIENTTYPE_DNS)
		if err != nil {
			return nil, err
		}

		len := gjson.Get(string(jsonData), "data.#")
		for i := int64(0); i < len.Int(); i++ {
			templateJson := gjson.Get(string(jsonData), fmt.Sprintf("data.%d", i))

			var template Template
			err := template.parse(templateJson.String())
			if err != nil {
				return nil, err
			}

			templates.PushBack(template)
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
	return templates, nil
}

func (d *Templates) GetTemplate(id int) (*Template, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("templates/%d", id), api.CLIENTTYPE_DNS)
	if err != nil {
		return nil, err
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var template Template

	err1 := template.parse(dataValue.String())
	if err1 != nil {
		return nil, err1
	}

	return &template, nil
}

func (d *Templates) Create(param TemplateParam) (int, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return 0, err
	}

	var paramJson string = string(resParam)

	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err1 = d.apiClient.DoPost("templates", []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return 0, err1
	}

	id := gjson.Get(string(jsonData), "data.id")

	return int(id.Int()), nil
}