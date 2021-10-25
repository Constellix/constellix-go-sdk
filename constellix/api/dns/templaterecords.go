package dns

import (
	"container/list"
	"github.com/tidwall/gjson"
	"encoding/json"
	"fmt"
	"constellix.com/constellix/api"
)

type TemplateRecord struct {
	apiClient *api.ApiClient
	templateId int
	Id 					int				`json:"id,omitempty"`
	Type				RecordType		`json:"type,omitempty"`
	Ttl					int				`json:"ttl,omitempty"`
	Enabled				bool			`json:"enabled,omitempty"`
	Name 				string			`json:"name,omitempty"`
	Region				RecordRegion	`json:"region,omitempty"`
	Notes 				string			`json:"notes,omitempty"`
	Mode 				RecordMode		`json:"mode,omitempty"`
	IpFilterId			int				`json:"-"`
	GeoProximityId		int				`json:"-"`
	ContactsId	 		int				`json:"-"`
	TemplateId			int				`json:"-"`
	ValueStandard 				*RecordValueStandard			`json:"-"`
	ValueFailover 				*RecordValueFailover			`json:"-"`
	ValueRoundrobinFailover 	*RecordValueRoundrobinFailover	`json:"-"`
	ValuePools					*RecordValuePools				`json:"-"`
	LastValues			LastValues		`json:"lastValues,omitempty"`
}

type TemplateRecordParam struct {
	Type				RecordType		`json:"type,omitempty"`
	Ttl					int				`json:"ttl,omitempty"`
	Enabled				bool			`json:"enabled,omitempty"`
	Name 				string			`json:"name,omitempty"`
	Region				RecordRegion	`json:"region,omitempty"`
	Notes 				string			`json:"notes,omitempty"`
	Mode 				RecordMode		`json:"mode,omitempty"`
	IpfilterId			int				`json:"ipfilter,omitempty"`
	GeoProximityId		int				`json:"geoproximity,omitempty"`
	ContactIds	 		[]int			`json:"contacts,omitempty"`
	ValueStandard 				*RecordValueStandard			`json:"-"`
	ValueFailover 				*RecordValueFailover			`json:"-"`
	ValueRoundrobinFailover 	*RecordValueRoundrobinFailover	`json:"-"`
	ValuePools					*RecordValuePools				`json:"-"`
}

func (d *TemplateRecord) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	ipFilterId := gjson.Get(string(jsonPayload), "ipfilter.id")
	d.IpFilterId = int(ipFilterId.Int())

	geoProximityId := gjson.Get(string(jsonPayload), "geoproximity.id")
	d.GeoProximityId = int(geoProximityId.Int())

	templateId := gjson.Get(string(jsonPayload), "template.id")
	d.TemplateId = int(templateId.Int())

	contactId := gjson.Get(string(jsonPayload), "contacts.id")
	d.ContactsId = int(contactId.Int())

	lastValuePoolsLen := gjson.Get(jsonPayload, "lastValues.pools.#")
	for i := int64(0); i < lastValuePoolsLen.Int(); i++ {
		poolId := gjson.Get(string(jsonPayload), fmt.Sprintf("lastValues.pools.%d.id", i))
		d.LastValues.Pools = append(d.LastValues.Pools, int(poolId.Int()))
	}

	valueJson := gjson.Get(string(jsonPayload), "value")
	if !valueJson.Exists(){
		return nil
	}

	if d.Mode == RECORDMODE_STANDARD {
		var value *RecordValueStandard = new(RecordValueStandard)
		err := json.Unmarshal([]byte(valueJson.String()), &(value.Values))
		if err != nil {
			return err
		}
		d.ValueStandard = value
	} else if d.Mode == RECORDMODE_FAILOVER {
		var value *RecordValueFailover = new(RecordValueFailover)
		err := json.Unmarshal([]byte(valueJson.String()), &(value.Values))
		if err != nil {
			return err
		}
		d.ValueFailover = value
	} else if d.Mode == RECORDMODE_ROUNDROBINFAILOVER {
		var value *RecordValueRoundrobinFailover = new(RecordValueRoundrobinFailover)
		err := json.Unmarshal([]byte(valueJson.String()), &(value.Values))
		if err != nil {
			return err
		}
		d.ValueRoundrobinFailover = value
	} else if d.Mode == RECORDMODE_POOLS {
		var value *RecordValuePools = new(RecordValuePools)
		err := json.Unmarshal([]byte(valueJson.String()), &(value.Values))
		if err != nil {
			return err
		}
		d.ValuePools = value
	}

	return nil
}

func (d *TemplateRecordParam) toJson() (string, error) {
	resParam, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	var paramJson string = string(resParam)

	var vStr string
	if d.ValueStandard != nil {
		v, err := json.Marshal(d.ValueStandard)
		if err != nil {
			return "", err
		}
		vStr = string(v)
	} else if d.ValueFailover != nil {
		v, err := json.Marshal(d.ValueFailover)
		if err != nil {
			return "", err
		}
		vStr = string(v)
	} else if d.ValueRoundrobinFailover != nil {
		v, err := json.Marshal(d.ValueRoundrobinFailover)
		if err != nil {
			return "", err
		}
		vStr = string(v)
	} else if d.ValuePools != nil {
		v, err := json.Marshal(d.ValuePools)
		if err != nil {
			return "", err
		}
		vStr = string(v)
	} else {
		return "", fmt.Errorf("TemplateRecordParam.Value is mandatory")
	}

	var res string = paramJson[0:len(paramJson)-1] + "," + vStr[1:len(vStr)-1] + "}"
	return res, nil
}

func (d *TemplateRecord) Update(param TemplateRecordParam) (*TemplateRecord, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var paramJson string
	var err error
	paramJson, err = param.toJson()
	if err != nil {
		return nil, err
	}

	var jsonData, err1 = d.apiClient.DoPut(fmt.Sprintf("templates/%d/records/%d", d.templateId, d.Id), []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return nil, err1
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var templateRecord TemplateRecord

	err2 := templateRecord.parse(dataValue.String())
	if err2 != nil {
		return nil, err2
	}

	templateRecord.templateId = d.templateId

	return &templateRecord, nil
}

func (d *TemplateRecord) Delete() error {
	d.apiClient = api.GetDnsApiClient("", "")

	var _, err = d.apiClient.DoDelete(fmt.Sprintf("templates/%d/records/%d", d.templateId, d.Id), api.CLIENTTYPE_DNS)
	if err != nil {
		return err
	}

	return nil
}

type TemplateRecords struct {
	apiClient *api.ApiClient
	templateId int
}

func (d *TemplateRecords) GetAll() (*list.List, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	templateRecords := list.New()

	var currentPage int
	for true {
		var url string
		if currentPage == 0 {
			urlP := &url
			*urlP = fmt.Sprintf("templates/%d/records", d.templateId)
		} else {
			urlP := &url
			*urlP = fmt.Sprintf("templates/%d/records?page=%d", d.templateId, currentPage)
		}

		var jsonData, err = d.apiClient.DoGet(url, api.CLIENTTYPE_DNS)
		if err != nil {
			return nil, err
		}

		len := gjson.Get(string(jsonData), "data.#")
		for i := int64(0); i < len.Int(); i++ {
			templateRecordJson := gjson.Get(string(jsonData), fmt.Sprintf("data.%d", i))

			var templateRecord TemplateRecord
			err := templateRecord.parse(templateRecordJson.String())
			if err != nil {
				return nil, err
			}

			templateRecord.templateId = d.templateId
			templateRecords.PushBack(templateRecord)
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
	return templateRecords, nil
}

func (d *TemplateRecords) GetRecord(id int) (*TemplateRecord, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("templates/%d/records/%d", d.templateId, id), api.CLIENTTYPE_DNS)
	if err != nil {
		return nil, err
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var templateRecord TemplateRecord

	err1 := templateRecord.parse(dataValue.String())
	if err1 != nil {
		return nil, err1
	}

	templateRecord.templateId = d.templateId

	return &templateRecord, nil
}

func (d *TemplateRecords) Create(param TemplateRecordParam) (int, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var paramJson string
	var err error
	paramJson, err = param.toJson()
	if err != nil {
		return 0, err
	}

	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err1 = d.apiClient.DoPost(fmt.Sprintf("templates/%d/records", d.templateId), []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return 0, err1
	}

	id := gjson.Get(string(jsonData), "data.id")

	return int(id.Int()), nil
}
