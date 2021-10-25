package sonar

import (
	"container/list"
	"github.com/tidwall/gjson"
	"encoding/json"
	"fmt"
	"constellix.com/constellix/api"
)

type ContactAddress struct {
	Address				string			`json:"address,omitempty"`
	Active				bool			`json:"active,omitempty"`
	ContactId			int				`json:"contactId,omitempty"`
	Type				string			`json:"type,omitempty"`
}

type Contact struct {
	Id				int					`json:"id,omitempty"`
	FirstName		string				`json:"firstName,omitempty"`
	LastName		string				`json:"lastName,omitempty"`
	AccountId		int					`json:"accountId,omitempty"`
	Addresses		[]ContactAddress	`json:"addresses,omitempty"`
}

type ContactGroup struct {
	Id				int					`json:"id,omitempty"`
	Name			string				`json:"name,omitempty"`
}

func (d *Contact) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	return nil
}

func (d *ContactGroup) parse(jsonPayload string) error{
	err := json.Unmarshal([]byte(jsonPayload), d)
	if err != nil {
		return err
	}

	return nil
}

type Contacts struct {
	apiClient *api.ApiClient
}

func (d *Contacts) GetContacts() (*list.List, error) {
	d.apiClient = api.GetSonarApiClient("", "")

	contacts := list.New()

	var jsonData, err = d.apiClient.DoGet("contacts", api.CLIENTTYPE_SONAR)
	if err != nil {
		return nil, err
	}

	len := gjson.Get(string(jsonData), "@this.#")
	for i := int64(0); i < len.Int(); i++ {
		contactJson := gjson.Get(string(jsonData), fmt.Sprintf("@this.%d", i))

		var contact Contact
		err := contact.parse(contactJson.String())
		if err != nil {
			return nil, err
		}

		contacts.PushBack(contact)
	}

	return contacts, nil
}

func (d *Contacts) GetContactGroups() (*list.List, error) {
	d.apiClient = api.GetSonarApiClient("", "")

	contactGroups := list.New()

	var jsonData, err = d.apiClient.DoGet("contact/groups", api.CLIENTTYPE_SONAR)
	if err != nil {
		return nil, err
	}

	len := gjson.Get(string(jsonData), "@this.#")
	for i := int64(0); i < len.Int(); i++ {
		groupJson := gjson.Get(string(jsonData), fmt.Sprintf("@this.%d", i))

		var group ContactGroup
		err := group.parse(groupJson.String())
		if err != nil {
			return nil, err
		}

		contactGroups.PushBack(group)
	}

	return contactGroups, nil
}
