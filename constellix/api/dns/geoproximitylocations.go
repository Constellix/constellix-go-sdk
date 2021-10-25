package dns

import (
	"container/list"
	"github.com/tidwall/gjson"
	"encoding/json"
	"fmt"
	"constellix.com/constellix/api"
)

type GeoProximityLocation struct {
	apiClient *api.ApiClient
	Id 					int			`json:"id,omitempty"`
	Name 				string		`json:"name,omitempty"`
	Country 			string		`json:"country,omitempty"`
	Region 				string		`json:"region,omitempty"`
	City 				string		`json:"city,omitempty"`
	Longitude			float32		`json:"longitude,omitempty"`
	Latitude 			float32		`json:"latitude,omitempty"`
}

type GeoProximityLocationParam struct {
	Name 				string		`json:"name,omitempty"`
	Country 			string		`json:"country,omitempty"`
	Region 				string		`json:"region,omitempty"`
	City 				string		`json:"city,omitempty"`
	Longitude			float32		`json:"longitude,omitempty"`
	Latitude 			float32		`json:"latitude,omitempty"`
}

func (d *GeoProximityLocation) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	return nil
}

func (d *GeoProximityLocation) Update(param GeoProximityLocationParam) (*GeoProximityLocation, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var paramJson string = string(resParam)

	var jsonData, err1 = d.apiClient.DoPut(fmt.Sprintf("geoproximities/%d", d.Id), []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return nil, err1
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var geoProximityLocation GeoProximityLocation

	err2 := geoProximityLocation.parse(dataValue.String())
	if err2 != nil {
		return nil, err2
	}

	return &geoProximityLocation, nil
}

func (d *GeoProximityLocation) Delete() error {
	d.apiClient = api.GetDnsApiClient("", "")

	var _, err = d.apiClient.DoDelete(fmt.Sprintf("geoproximities/%d", d.Id), api.CLIENTTYPE_DNS)
	if err != nil {
		return err
	}

	return nil
}

type GeoProximityLocations struct {
	apiClient *api.ApiClient
}

func (d *GeoProximityLocations) GetAll() (*list.List, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	geoProximityLocations := list.New()

	var currentPage int
	for true {
		var url string
		if currentPage == 0 {
			urlP := &url
			*urlP = "geoproximities"
		} else {
			urlP := &url
			*urlP = fmt.Sprintf("geoproximities?page=%d", currentPage)
		}

		var jsonData, err = d.apiClient.DoGet(url, api.CLIENTTYPE_DNS)
		if err != nil {
			return nil, err
		}

		len := gjson.Get(string(jsonData), "data.#")
		for i := int64(0); i < len.Int(); i++ {
			geoProximityLocationJson := gjson.Get(string(jsonData), fmt.Sprintf("data.%d", i))

			var geoProximityLocation GeoProximityLocation
			err := geoProximityLocation.parse(geoProximityLocationJson.String())
			if err != nil {
				return nil, err
			}

			geoProximityLocations.PushBack(geoProximityLocation)
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
	return geoProximityLocations, nil
}

func (d *GeoProximityLocations) GetGeoProximityLocation(id int) (*GeoProximityLocation, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("geoproximities/%d", id), api.CLIENTTYPE_DNS)
	if err != nil {
		return nil, err
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var geoProximityLocation GeoProximityLocation

	err1 := geoProximityLocation.parse(dataValue.String())
	if err1 != nil {
		return nil, err1
	}

	return &geoProximityLocation, nil
}

func (d *GeoProximityLocations) Create(param GeoProximityLocationParam) (int, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	resParam, err := json.Marshal(param)
	if err != nil {
		return 0, err
	}

	var paramJson string = string(resParam)

	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err1 = d.apiClient.DoPost("geoproximities", []byte(paramJson), api.CLIENTTYPE_DNS)
	if err1 != nil {
		return 0, err1
	}

	id := gjson.Get(string(jsonData), "data.id")

	return int(id.Int()), nil
}
