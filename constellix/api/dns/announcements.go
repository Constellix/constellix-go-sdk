package dns

import (
	"container/list"
	"github.com/tidwall/gjson"
	"encoding/json"
	"fmt"
	"constellix.com/constellix/api"
)

type Announcement struct {
	apiClient *api.ApiClient
	Id 					int			`json:"id,omitempty"`
	Type 				string		`json:"type,omitempty"`
	Link 				string		`json:"link,omitempty"`
	Title 				string		`json:"title,omitempty"`
}

func (d *Announcement) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	return nil
}

type Announcements struct {
	apiClient *api.ApiClient
}

func (d *Announcements) GetAll() (*list.List, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	announcements := list.New()

	var currentPage int
	for true {
		var url string
		if currentPage == 0 {
			urlP := &url
			*urlP = "announcements"
		} else {
			urlP := &url
			*urlP = fmt.Sprintf("announcements?page=%d", currentPage)
		}

		var jsonData, err = d.apiClient.DoGet(url, api.CLIENTTYPE_DNS)
		if err != nil {
			return nil, err
		}

		len := gjson.Get(string(jsonData), "data.#")
		for i := int64(0); i < len.Int(); i++ {
			announcementJson := gjson.Get(string(jsonData), fmt.Sprintf("data.%d", i))

			var announcement Announcement
			err := announcement.parse(announcementJson.String())
			if err != nil {
				return nil, err
			}

			announcements.PushBack(announcement)
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
	return announcements, nil
}

func (d *Announcements) GetAnnouncement(id int) (*Announcement, error) {
	d.apiClient = api.GetDnsApiClient("", "")

	var jsonData, err = d.apiClient.DoGet(fmt.Sprintf("announcements/%d", id), api.CLIENTTYPE_DNS)
	if err != nil {
		return nil, err
	}

	dataValue := gjson.Get(string(jsonData), "data")
	var announcement Announcement

	err1 := announcement.parse(dataValue.String())
	if err1 != nil {
		return nil, err1
	}

	return &announcement, nil
}
