package dns

import (
	"container/list"
	"github.com/tidwall/gjson"
	"encoding/json"
	"fmt"
	"constellix.com/constellix/api"
)

type Soa struct {
	PrimaryNameServer string `json:"primaryNameserver,omitempty"`
	Email             string `json:"email,omitempty"`
	Ttl               int    `json:"ttl,omitempty"`
	Serial            int    `json:"serial,omitempty"`
	Refresh           int    `json:"refresh,omitempty"`
	Retry             int    `json:"retry,omitempty"`
	Expire            int    `json:"expire,omitempty"`
	NegativeCache     int    `json:"negativeCache,omitempty"`
}

type Domain struct {
	apiClient *api.ApiClient
	Id 					int			`json:"id,omitempty"`
	Name 				string		`json:"name,omitempty"`
	Note 				string		`json:"note,omitempty"`
	Status 				string		`json:"status,omitempty"`
	Version 			int			`json:"version,omitempty"`
	Soa					Soa			`json:"soa,omitempty"`
	GeoIp 				bool		`json:"geoip,omitempty"`
	Gtd 				bool		`json:"gtd,omitempty"`
	NameServers 		[]string	`json:"nameservers,omitempty"`
	TagIds 				[]int		`json:"-"`
	TemplateId	 		int			`json:"-"`
	VanityNameserverId	int			`json:"-"`
	ContactIds 			[]int		`json:"-"`
	Records				DomainRecords 	`json:"-"`
	History				DomainHistory 	`json:"-"`
	Snapshots			DomainSnapshots	`json:"-"`
}

type DomainParam struct {
	Name 				string		`json:"name,omitempty"`
	Note 				string		`json:"note,omitempty"`
	Soa					Soa			`json:"soa,omitempty"`
	GeoIp 				bool		`json:"geoip,omitempty"`
	Gtd 				bool		`json:"gtd,omitempty"`
	NameServers 		[]string	`json:"nameservers,omitempty"`
	TagId 				[]int		`json:"tags,omitempty"`
	TemplateId	 		int			`json:"template,omitempty"`
	VanityNameserverId	int			`json:"vanity_nameserver,omitempty"`
	ContactIds 			[]int		`json:"contacts,omitempty"`
}

func (d *Domain) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	tagsLen := gjson.Get(jsonPayload, "tags.#")
	for i := int64(0); i < tagsLen.Int(); i++ {
		tagId := gjson.Get(string(jsonPayload), fmt.Sprintf("tags.%d.id", i))
		d.TagIds = append(d.TagIds, int(tagId.Int()))
	}

	templateId := gjson.Get(string(jsonPayload), "template.id")
	d.TemplateId = int(templateId.Int())

	vanityNameserverId := gjson.Get(string(jsonPayload), "vanityNameserver.id")
	d.VanityNameserverId = int(vanityNameserverId.Int())

	contactsLen := gjson.Get(jsonPayload, "contacts.#")
	for i := int64(0); i < contactsLen.Int(); i++ {
		contactId := gjson.Get(string(jsonPayload), fmt.Sprintf("contacts.%d.id", i))
		d.ContactIds = append(d.ContactIds, int(contactId.Int()))
	}

	d.Records.domainId = d.Id
	d.History.domainId = d.Id
	d.Snapshots.domainId = d.Id

	return nil
}

func (d *Domain) Update(param DomainParam) (*Domain, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var paramJson string = string(resParam)

	var jsonData, err1 = d.apiClient.DoPut(fmt.Sprintf("domains/%d", d.Id), []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return nil, err1
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var domain Domain

	err2 := domain.parse(dataValue.String())
	if err2 != nil {
		return nil, err2
	}

	return &domain, nil
}

func (d *Domain) Delete() error {
	d.apiClient = api.GetDnsApiClient("", "")

	var _, err = d.apiClient.DoDelete(fmt.Sprintf("domains/%d", d.Id), api.CLIENTTYPE_DNS)
	if err != nil {
		return err
	}

	return nil
}

type Domains struct {
	apiClient *api.ApiClient
}

func (d *Domains) GetAll() (*list.List, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	domains := list.New()

	var currentPage int
	for true {
		var url string
		if currentPage == 0 {
			urlP := &url
			*urlP = "domains"
		} else {
			urlP := &url
			*urlP = fmt.Sprintf("domains?page=%d", currentPage)
		}

		var jsonData, err = d.apiClient.DoGet(url, api.CLIENTTYPE_DNS)
		if err != nil {
			return nil, err
		}

		len := gjson.Get(string(jsonData), "data.#")
		for i := int64(0); i < len.Int(); i++ {
			domainJson := gjson.Get(string(jsonData), fmt.Sprintf("data.%d", i))

			var domain Domain
			err := domain.parse(domainJson.String())
			if err != nil {
				return nil, err
			}

			domains.PushBack(domain)
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
	return domains, nil
}

func (d *Domains) GetDomain(id int) (*Domain, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("domains/%d", id), api.CLIENTTYPE_DNS)
	if err != nil {
		return nil, err
	}
/*
	jsonData := `
	{
		"data": {
		  "id": 366246,
		  "name": "example.com",
		  "note": "My Domain",
		  "status": "ACTIVE",
		  "version": 3,
		  "soa": {
			"primaryNameserver": "ns11.constellix.com",
			"email": "admin.example.com",
			"ttl": 86400,
			"serial": 2020061601,
			"refresh": 86400,
			"retry": 7200,
			"expire": 3600000,
			"negativeCache": 180
		  },
		  "geoip": true,
		  "gtd": true,
		  "nameservers": [
			"ns11.constellix.com",
			"ns21.constellix.com",
			"ns31.constellix.com"
		  ],
		  "tags": [
			{
			  "id": 824,
			  "name": "My Tag",
			  "links": {
				"self": "/api/v4/tags/824"
			  }
			}
		  ],
		  "template": {
			"id": 83675283,
			"name": "My Template",
			"version": 3,
			"links": {
			  "self": "/api/v4/templates/83675283",
			  "records": "/api/v4/records?template=83675283"
			}
		  },
		  "vanityNameserver": {
			"id": 82648967,
			"links": {
			  "self": "/api/v4/vanitynameservers/82648967"
			}
		  },
		  "contacts": [
			{
			  "id": 2668228,
			  "links": {
				"self": "/api/v4/contactlists/2668228"
			  }
			}
		  ],
		  "createdAt": "2021-06-07T17:52:00Z",
		  "updatedAt": "2021-06-07T17:52:00Z",
		  "links": {
			"self": "/api/v4/domains/366246",
			"records": "/api/v4/domains/366246/records"
		  }
		}
	  }
	`
*/
	dataValue := gjson.Get(string(jsonData), "data")
	var domain Domain

	err1 := domain.parse(dataValue.String())
	if err1 != nil {
		return nil, err1
	}

	return &domain, nil
}

func (d *Domains) Create(param DomainParam) (int, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return 0, err
	}

	var paramJson string = string(resParam)

	var jsonData, err1 = d.apiClient.DoPost("domains", []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return 0, err1
	}

	id := gjson.Get(string(jsonData), "data.id")

	return int(id.Int()), nil
}