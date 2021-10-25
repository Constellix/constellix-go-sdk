package dns

import (
	"container/list"
	"github.com/tidwall/gjson"
	"encoding/json"
	"fmt"
	"constellix.com/constellix/api"
)

type ContactEmail struct {
	Address 			string		`json:"address,omitempty"`
	Verified 			bool		`json:"verified,omitempty"`
}
type ContactList struct {
	apiClient *api.ApiClient
	Id 					int				`json:"id,omitempty"`
	Name 				string			`json:"name,omitempty"`
	EmailCount 			int				`json:"emailCount,omitempty"`
	Emails 				[]ContactEmail	`json:"emails,omitempty"`
}

type ContactListParam struct {
	Name 				string		`json:"name,omitempty"`
	Emails 				[]string	`json:"emails,omitempty"`
}

func (d *ContactList) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	return nil
}

func (d *ContactList) Update(param ContactListParam) (*ContactList, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var paramJson string = string(resParam)

	var jsonData, err1 = d.apiClient.DoPut(fmt.Sprintf("contactlists/%d", d.Id), []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return nil, err1
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var contactList ContactList

	err2 := contactList.parse(dataValue.String())
	if err2 != nil {
		return nil, err2
	}

	return &contactList, nil
}

func (d *ContactList) Delete() error {
	d.apiClient = api.GetDnsApiClient("", "")

	var _, err = d.apiClient.DoDelete(fmt.Sprintf("contactlists/%d", d.Id), api.CLIENTTYPE_DNS)
	if err != nil {
		return err
	}

	return nil
}

type ContactLists struct {
	apiClient *api.ApiClient
}

func (d *ContactLists) GetAll() (*list.List, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	contactLists := list.New()

	var currentPage int
	for true {
		var url string
		if currentPage == 0 {
			urlP := &url
			*urlP = "contactlists"
		} else {
			urlP := &url
			*urlP = fmt.Sprintf("contactlists?page=%d", currentPage)
		}

		var jsonData, err = d.apiClient.DoGet(url, api.CLIENTTYPE_DNS)
		if err != nil {
			return nil, err
		}

		len := gjson.Get(string(jsonData), "data.#")
		for i := int64(0); i < len.Int(); i++ {
			contactListJson := gjson.Get(string(jsonData), fmt.Sprintf("data.%d", i))

			var contactList ContactList
			err := contactList.parse(contactListJson.String())
			if err != nil {
				return nil, err
			}

			contactLists.PushBack(contactList)
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
	return contactLists, nil
}

func (d *ContactLists) GetContactList(id int) (*ContactList, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("contactlists/%d", id), api.CLIENTTYPE_DNS)
	if err != nil {
		return nil, err
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var contactList ContactList

	err1 := contactList.parse(dataValue.String())
	if err1 != nil {
		return nil, err1
	}

	return &contactList, nil
}

func (d *ContactLists) Create(param ContactListParam) (int, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return 0, err
	}

	var paramJson string = string(resParam)

	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err1 = d.apiClient.DoPost("contactlists", []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return 0, err1
	}

	id := gjson.Get(string(jsonData), "data.id")

	return int(id.Int()), nil
}