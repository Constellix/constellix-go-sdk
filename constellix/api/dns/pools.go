package dns

import (
	"container/list"
	"github.com/tidwall/gjson"
	"encoding/json"
	"fmt"
	"constellix.com/constellix/api"
)

type ItoConfig struct {
	Period 					PoolPeriod					`json:"period,omitempty"`
	MaximumNumberOfResults 	int							`json:"maximumNumberOfResults,omitempty"`
	DeviationAllowance 		PoolDeviationAllowance		`json:"deviationAllowance,omitempty"`
	MonitoringRegion 		PoolMonitoringRegion		`json:"monitoringRegion,omitempty"`
	HandicapFactor 			PoolHandicapFactor			`json:"handicapFactor,omitempty"`
}

type Ito struct {
	Enabled 	bool 		`json:"enabled,omitempty"`
	Config 		ItoConfig	`json:"config,omitempty"`
}

type PoolValue struct {
	Value 			string				`json:"value,omitempty"`
	Weight 			int					`json:"weight,omitempty"`
	Enabled 		bool				`json:"enabled,omitempty"`
	Handicap 		int					`json:"handicap,omitempty"`
	Policy 			PoolValuePolicy		`json:"policy,omitempty"`
	SonarCheckId	int					`json:"sonarCheckId,omitempty"`
}

type Pool struct {
	apiClient *api.ApiClient
	Id 					int				`json:"id,omitempty"`
	Type 				PoolType		`json:"type,omitempty"`
	Name 				string			`json:"name,omitempty"`
	Return				int				`json:"return,omitempty"`
	MinimumFailover 	int				`json:"minimumFailover,omitempty"`
	Failed				bool			`json:"failed,omitempty"`
	Enabled				bool			`json:"enabled,omitempty"`
	Values 				[]PoolValue		`json:"values,omitempty"`
	Ito 				Ito				`json:"ito,omitempty"`
	DomainIds 			[]int			`json:"-"`
	TemplateIds	 		[]int			`json:"-"`
	ContactIds 			[]int			`json:"-"`
}

type PoolParam struct {
	apiClient *api.ApiClient
	Type 				PoolType		`json:"type,omitempty"`
	Name 				string			`json:"name,omitempty"`
	Return				int				`json:"return,omitempty"`
	MinimumFailover 	int				`json:"minimumFailover,omitempty"`
	Enabled				bool			`json:"enabled,omitempty"`
	Values 				[]PoolValue		`json:"values,omitempty"`
	Contacts		 	[]int			`json:"contacts,omitempty"`
	Ito 				Ito				`json:"ito,omitempty"`
}

func (d *Pool) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	domainsLen := gjson.Get(jsonPayload, "domains.#")
	for i := int64(0); i < domainsLen.Int(); i++ {
		domainId := gjson.Get(string(jsonPayload), fmt.Sprintf("domains.%d.id", i))
		d.DomainIds = append(d.DomainIds, int(domainId.Int()))
	}

	templatesLen := gjson.Get(jsonPayload, "templates.#")
	for i := int64(0); i < templatesLen.Int(); i++ {
		templateId := gjson.Get(string(jsonPayload), fmt.Sprintf("templates.%d.id", i))
		d.TemplateIds = append(d.TemplateIds, int(templateId.Int()))
	}

	contactsLen := gjson.Get(jsonPayload, "contacts.#")
	for i := int64(0); i < contactsLen.Int(); i++ {
		contactId := gjson.Get(string(jsonPayload), fmt.Sprintf("contacts.%d.id", i))
		d.ContactIds = append(d.ContactIds, int(contactId.Int()))
	}

	return nil
}

func (d *Pool) Update(param PoolParam) (*Pool, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var paramJson string = string(resParam)

	var jsonData, err1 = d.apiClient.DoPut(fmt.Sprintf("pools/%s/%d", d.Type, d.Id), []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return nil, err1
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var pool Pool

	err2 := pool.parse(dataValue.String())
	if err2 != nil {
		return nil, err2
	}

	return &pool, nil
}

func (d *Pool) Delete() error {
	d.apiClient = api.GetDnsApiClient("", "")

	var _, err = d.apiClient.DoDelete(fmt.Sprintf("pools/%s/%d", d.Type, d.Id), api.CLIENTTYPE_DNS)
	if err != nil {
		return err
	}

	return nil
}

type Pools struct {
	apiClient *api.ApiClient
}

func (d *Pools) GetAll() (*list.List, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	pools := list.New()

	var currentPage int
	for true {
		var url string
		if currentPage == 0 {
			urlP := &url
			*urlP = "pools"
		} else {
			urlP := &url
			*urlP = fmt.Sprintf("pools?page=%d", currentPage)
		}

		var jsonData, err = d.apiClient.DoGet(url, api.CLIENTTYPE_DNS)
		if err != nil {
			return nil, err
		}

		len := gjson.Get(string(jsonData), "data.#")
		for i := int64(0); i < len.Int(); i++ {
			poolJson := gjson.Get(string(jsonData), fmt.Sprintf("data.%d", i))

			var pool Pool
			err := pool.parse(poolJson.String())
			if err != nil {
				return nil, err
			}

			pools.PushBack(pool)
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
	return pools, nil
}

func (d *Pools) GetPool(t PoolType, id int) (*Pool, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("pools/%s/%d", t, id), api.CLIENTTYPE_DNS)
	if err != nil {
		return nil, err
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var pool Pool

	err1 := pool.parse(dataValue.String())
	if err1 != nil {
		return nil, err1
	}

	return &pool, nil
}

func (d *Pools) Create(param PoolParam) (int, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return 0, err
	}

	var paramJson string = string(resParam)

	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err1 = d.apiClient.DoPost("pools", []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return 0, err1
	}

	id := gjson.Get(string(jsonData), "data.id")

	return int(id.Int()), nil
}