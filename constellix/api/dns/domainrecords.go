package dns

import (
	"container/list"
	"github.com/tidwall/gjson"
	"encoding/json"
	"fmt"
	"constellix.com/constellix/api"
)

type RecordValueShort struct {
	Enabled			bool 		`json:"enabled"`
	Value 			string 		`json:"value"`
}

type RecordValueExtended struct {
	Enabled			bool 		`json:"enabled"`
	Order			int 		`json:"order"`
	SonarCheckId	int 		`json:"sonarCheckId"`
	Value 			string 		`json:"value"`
}

type RecordValueExtendedLast struct {
	Enabled			bool 		`json:"enabled"`
	Order			int 		`json:"order"`
	SonarCheckId	int 		`json:"sonarCheckId"`
	Active			bool 		`json:"active"`
	Failed			bool 		`json:"failed"`
	Status 			string 		`json:"status"`
	Value 			string 		`json:"value"`
}

type RecordValueObject struct {
	Mode		RecordValueMode  		`json:"mode"`
	Values      []RecordValueExtended   `json:"values"`
}

type LastValues struct {
	Standard 			[]RecordValueShort 			`json:"standard"`
	Failover			RecordValueObject			`json:"failover"`
	RoundRobinFailover	[]RecordValueExtendedLast	`json:"roundRobinFailover"`
	Pools				[]int						`json:"pools"`
}

type RecordValueStandard struct {
	Values 		[]RecordValueShort 			`json:"value"`
}

type RecordValueFailover struct {
	Values 		RecordValueObject 			`json:"value"`
}

type RecordValueRoundrobinFailover struct {
	Values 		[]RecordValueExtended 		`json:"value"`
}

type RecordValuePools struct {
	Values 		[]int 						`json:"value"`
}

type DomainRecord struct {
	apiClient *api.ApiClient
	domainId int
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
	DomainId			int				`json:"-"`
	ValueStandard 				*RecordValueStandard			`json:"-"`
	ValueFailover 				*RecordValueFailover			`json:"-"`
	ValueRoundrobinFailover 	*RecordValueRoundrobinFailover	`json:"-"`
	ValuePools					*RecordValuePools				`json:"-"`
	LastValues			LastValues		`json:"lastValues,omitempty"`
}

type DomainRecordParam struct {
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

func (d *DomainRecord) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	ipFilterId := gjson.Get(string(jsonPayload), "ipfilter.id")
	d.IpFilterId = int(ipFilterId.Int())

	geoProximityId := gjson.Get(string(jsonPayload), "geoproximity.id")
	d.GeoProximityId = int(geoProximityId.Int())

	domainId := gjson.Get(string(jsonPayload), "domain.id")
	d.DomainId = int(domainId.Int())

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

	/*modeValue := gjson.Get(string(jsonPayload), "value.mode")
	if modeValue.Exists() {
		// Value 2
		var value2 *RecordValue2 = new(RecordValue2)
		err2 := json.Unmarshal([]byte(value.String()), &(value2.Values))
		if err2 != nil {
			return err2
		}
		d.Value2 = value2
	} else {
		orderValue := gjson.Get(string(jsonPayload), "value.0.order")
		if orderValue.Exists() {
			// Value 3
			var value3 *RecordValue3 = new(RecordValue3)
			err3 := json.Unmarshal([]byte(value.String()), &(value3.Values))
			if err3 != nil {
				return err3
			}
			d.Value3 = value3
		} else {
			enabledValue := gjson.Get(string(jsonPayload), "value.0.enabled")
			if enabledValue.Exists() {
				// Value 1
				var value1 *RecordValue1 = new(RecordValue1)
				err1 := json.Unmarshal([]byte(value.String()), &(value1.Values))
				if err1 != nil {
					return err1
				}
				d.Value1 = value1
			} else {
				// Value 4
				var value4 *RecordValue4 = new(RecordValue4)
				err4 := json.Unmarshal([]byte(value.String()), &(value4.Values))
				if err4 != nil {
					return err4
				}
				d.Value4 = value4
			}
		}
	}*/

	return nil
}

func (d *DomainRecordParam) toJson() (string, error) {
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
		return "", fmt.Errorf("DomainRecordParam.Value is mandatory")
	}

	var res string = paramJson[0:len(paramJson)-1] + "," + vStr[1:len(vStr)-1] + "}"
	return res, nil
}

func (d *DomainRecord) Update(param DomainRecordParam) (*DomainRecord, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var paramJson string
	var err error
	paramJson, err = param.toJson()
	if err != nil {
		return nil, err
	}

	var jsonData, err1 = d.apiClient.DoPut(fmt.Sprintf("domains/%d/records/%d", d.domainId, d.Id), []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return nil, err1
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var domainRecord DomainRecord

	err2 := domainRecord.parse(dataValue.String())
	if err2 != nil {
		return nil, err2
	}

	domainRecord.domainId = d.domainId

	return &domainRecord, nil
}

func (d *DomainRecord) Delete() error {
	d.apiClient = api.GetDnsApiClient("", "")

	var _, err = d.apiClient.DoDelete(fmt.Sprintf("domains/%d/records/%d", d.domainId, d.Id), api.CLIENTTYPE_DNS)
	if err != nil {
		return err
	}

	return nil
}

type DomainRecords struct {
	apiClient *api.ApiClient
	domainId int
}

func (d *DomainRecords) GetAll() (*list.List, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	domainRecords := list.New()

	var currentPage int
	for true {
		var url string
		if currentPage == 0 {
			urlP := &url
			*urlP = fmt.Sprintf("domains/%d/records", d.domainId)
		} else {
			urlP := &url
			*urlP = fmt.Sprintf("domains/%d/records?page=%d", d.domainId, currentPage)
		}

		var jsonData, err = d.apiClient.DoGet(url, api.CLIENTTYPE_DNS)
		if err != nil {
			return nil, err
		}

		len := gjson.Get(string(jsonData), "data.#")
		for i := int64(0); i < len.Int(); i++ {
			domainRecordJson := gjson.Get(string(jsonData), fmt.Sprintf("data.%d", i))

			var domainRecord DomainRecord
			err := domainRecord.parse(domainRecordJson.String())
			if err != nil {
				return nil, err
			}

			domainRecord.domainId = d.domainId
			domainRecords.PushBack(domainRecord)
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
	return domainRecords, nil
}

func (d *DomainRecords) GetRecord(id int) (*DomainRecord, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("domains/%d/records/%d", d.domainId, id), api.CLIENTTYPE_DNS)
	if err != nil {
		return nil, err
	}

/*
	jsonData := `
	{
		"data": {
		  "id": 732673,
		  "type": "A",
		  "ttl": 3600,
		  "enabled": true,
		  "name": "www",
		  "region": "default",
		  "ipfilter": {
			"id": 47345837,
			"name": "My IP filter",
			"links": {
			  "self": "/api/v4/ipfilters/47345837"
			}
		  },
		  "geoproximity": {
			"id": 4367769,
			"name": "My Geo Proximity Location",
			"links": {
			  "self": "/api/v4/geoproximities/4367769"
			}
		  },
		  "ipfilterDrop": true,
		  "notes": "This is my DNS record",
		  "contacts": {
			"id": 2668228,
			"links": {
			  "self": "/api/v4/contactlists/2668228"
			}
		  },
		  "mode": "standard",
		  "value":  [
			  {
				"enabled": true,
				"value": "127.0.0.1" 
			  }
				
			  ],
		
		
		  "lastValues": {
			"standard": [
			  {
				"enabled": true,
				"value": "198.51.100.42"
			  }
			],
			"failover": {
			  "mode": "normal",
			  "values": [
				{
				  "enabled": true,
				  "order": 1,
				  "sonarCheckId": 76627,
				  "active": true,
				  "failed": false,
				  "status": "N/A",
				  "value": "198.51.100.42"
				}
			  ]
			},
			"pools": [
			  {
				"id": 7665,
				"links": {
				  "self": "/api/v4/pools/7665"
				}
			  }
			],
			"roundRobinFailover": [
			  {
				"enabled": true,
				"order": 1,
				"sonarCheckId": 76627,
				"active": true,
				"failed": false,
				"status": "N/A",
				"value": "198.51.100.42"
			  }
			]
		  },
		  "domain": {
			"id": 366246,
			"name": "example.com",
			"note": "My Domain",
			"status": "ACTIVE",
			"version": 3,
			"geoip": true,
			"gtd": true,
			"tags": [
			  {
				"id": 824,
				"name": "My Tag",
				"links": {
				  "self": "/api/v4/tags/824"
				}
			  }
			],
			"createdAt": "2021-06-07T17:52:00Z",
			"updatedAt": "2021-06-07T17:52:00Z",
			"links": {
			  "self": "/api/v4/domains/366246",
			  "records": "/api/v4/domains/366246/records"
			}
		  },
		  "links": {
			"self": "/api/v4/domains/366246/records/732673"
		  }
		}
	  }
	`

	jsonData := `
	{
		"data": {
		  "id": 732673,
		  "type": "A",
		  "ttl": 3600,
		  "enabled": true,
		  "name": "www",
		  "region": "default",
		  "ipfilter": {
			"id": 47345837,
			"name": "My IP filter",
			"links": {
			  "self": "/api/v4/ipfilters/47345837"
			}
		  },
		  "geoproximity": {
			"id": 4367769,
			"name": "My Geo Proximity Location",
			"links": {
			  "self": "/api/v4/geoproximities/4367769"
			}
		  },
		  "ipfilterDrop": true,
		  "notes": "This is my DNS record",
		  "contacts": {
			"id": 2668228,
			"links": {
			  "self": "/api/v4/contactlists/2668228"
			}
		  },
		  "mode": "standard",
		  "value":  {
			  "mode": "one-way",
			  "values": [
			  {
				"enabled": true,
				"order": 25,
				"sonarCheckId": 169,
				"value": "127.0.0.1" 
			  }
				
			  ]
			},
		
		  "lastValues": {
			"standard": [
			  {
				"enabled": true,
				"value": "198.51.100.42"
			  }
			],
			"failover": {
			  "mode": "normal",
			  "values": [
				{
				  "enabled": true,
				  "order": 1,
				  "sonarCheckId": 76627,
				  "active": true,
				  "failed": false,
				  "status": "N/A",
				  "value": "198.51.100.42"
				}
			  ]
			},
			"pools": [
			  {
				"id": 7665,
				"links": {
				  "self": "/api/v4/pools/7665"
				}
			  }
			],
			"roundRobinFailover": [
			  {
				"enabled": true,
				"order": 1,
				"sonarCheckId": 76627,
				"active": true,
				"failed": false,
				"status": "N/A",
				"value": "198.51.100.42"
			  }
			]
		  },
		  "domain": {
			"id": 366246,
			"name": "example.com",
			"note": "My Domain",
			"status": "ACTIVE",
			"version": 3,
			"geoip": true,
			"gtd": true,
			"tags": [
			  {
				"id": 824,
				"name": "My Tag",
				"links": {
				  "self": "/api/v4/tags/824"
				}
			  }
			],
			"createdAt": "2021-06-07T17:52:00Z",
			"updatedAt": "2021-06-07T17:52:00Z",
			"links": {
			  "self": "/api/v4/domains/366246",
			  "records": "/api/v4/domains/366246/records"
			}
		  },
		  "links": {
			"self": "/api/v4/domains/366246/records/732673"
		  }
		}
	  }
	`

	jsonData := `
	{
		"data": {
		  "id": 732673,
		  "type": "A",
		  "ttl": 3600,
		  "enabled": true,
		  "name": "www",
		  "region": "default",
		  "ipfilter": {
			"id": 47345837,
			"name": "My IP filter",
			"links": {
			  "self": "/api/v4/ipfilters/47345837"
			}
		  },
		  "geoproximity": {
			"id": 4367769,
			"name": "My Geo Proximity Location",
			"links": {
			  "self": "/api/v4/geoproximities/4367769"
			}
		  },
		  "ipfilterDrop": true,
		  "notes": "This is my DNS record",
		  "contacts": {
			"id": 2668228,
			"links": {
			  "self": "/api/v4/contactlists/2668228"
			}
		  },
		  "mode": "standard",
		  "value":  [
			  {
				"enabled": true,
				"order": 25,
				"sonarCheckId": 169,
				"value": "127.0.0.1" 
			  }
				
			  ],
		
		  "lastValues": {
			"standard": [
			  {
				"enabled": true,
				"value": "198.51.100.42"
			  }
			],
			"failover": {
			  "mode": "normal",
			  "values": [
				{
				  "enabled": true,
				  "order": 1,
				  "sonarCheckId": 76627,
				  "active": true,
				  "failed": false,
				  "status": "N/A",
				  "value": "198.51.100.42"
				}
			  ]
			},
			"pools": [
			  {
				"id": 7665,
				"links": {
				  "self": "/api/v4/pools/7665"
				}
			  }
			],
			"roundRobinFailover": [
			  {
				"enabled": true,
				"order": 1,
				"sonarCheckId": 76627,
				"active": true,
				"failed": false,
				"status": "N/A",
				"value": "198.51.100.42"
			  }
			]
		  },
		  "domain": {
			"id": 366246,
			"name": "example.com",
			"note": "My Domain",
			"status": "ACTIVE",
			"version": 3,
			"geoip": true,
			"gtd": true,
			"tags": [
			  {
				"id": 824,
				"name": "My Tag",
				"links": {
				  "self": "/api/v4/tags/824"
				}
			  }
			],
			"createdAt": "2021-06-07T17:52:00Z",
			"updatedAt": "2021-06-07T17:52:00Z",
			"links": {
			  "self": "/api/v4/domains/366246",
			  "records": "/api/v4/domains/366246/records"
			}
		  },
		  "links": {
			"self": "/api/v4/domains/366246/records/732673"
		  }
		}
	  }
	`

	jsonData := `
	{
		"data": {
		  "id": 732673,
		  "type": "A",
		  "ttl": 3600,
		  "enabled": true,
		  "name": "www",
		  "region": "default",
		  "ipfilter": {
			"id": 47345837,
			"name": "My IP filter",
			"links": {
			  "self": "/api/v4/ipfilters/47345837"
			}
		  },
		  "geoproximity": {
			"id": 4367769,
			"name": "My Geo Proximity Location",
			"links": {
			  "self": "/api/v4/geoproximities/4367769"
			}
		  },
		  "ipfilterDrop": true,
		  "notes": "This is my DNS record",
		  "contacts": {
			"id": 2668228,
			"links": {
			  "self": "/api/v4/contactlists/2668228"
			}
		  },
		  "mode": "standard",
		  "value":  [
				1, 2, 3
			  ],
		
		  "lastValues": {
			"standard": [
			  {
				"enabled": true,
				"value": "198.51.100.42"
			  }
			],
			"failover": {
			  "mode": "normal",
			  "values": [
				{
				  "enabled": true,
				  "order": 1,
				  "sonarCheckId": 76627,
				  "active": true,
				  "failed": false,
				  "status": "N/A",
				  "value": "198.51.100.42"
				}
			  ]
			},
			"pools": [
			  {
				"id": 7665,
				"links": {
				  "self": "/api/v4/pools/7665"
				}
			  }
			],
			"roundRobinFailover": [
			  {
				"enabled": true,
				"order": 1,
				"sonarCheckId": 76627,
				"active": true,
				"failed": false,
				"status": "N/A",
				"value": "198.51.100.42"
			  }
			]
		  },
		  "domain": {
			"id": 366246,
			"name": "example.com",
			"note": "My Domain",
			"status": "ACTIVE",
			"version": 3,
			"geoip": true,
			"gtd": true,
			"tags": [
			  {
				"id": 824,
				"name": "My Tag",
				"links": {
				  "self": "/api/v4/tags/824"
				}
			  }
			],
			"createdAt": "2021-06-07T17:52:00Z",
			"updatedAt": "2021-06-07T17:52:00Z",
			"links": {
			  "self": "/api/v4/domains/366246",
			  "records": "/api/v4/domains/366246/records"
			}
		  },
		  "links": {
			"self": "/api/v4/domains/366246/records/732673"
		  }
		}
	  }
	`
*/


	dataValue := gjson.Get(string(jsonData), "data")
	var domainRecord DomainRecord

	err1 := domainRecord.parse(dataValue.String())
	if err1 != nil {
		return nil, err1
	}

	domainRecord.domainId = d.domainId

	return &domainRecord, nil
}

func (d *DomainRecords) Create(param DomainRecordParam) (int, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var paramJson string
	var err error
	paramJson, err = param.toJson()
	if err != nil {
		return 0, err
	}

	var jsonData, err1 = d.apiClient.DoPost(fmt.Sprintf("domains/%d/records", d.domainId), []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return 0, err1
	}

	id := gjson.Get(string(jsonData), "data.id")

	return int(id.Int()), nil
}
