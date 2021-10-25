package main

import (
	"constellix.com/constellix/api/sonar"
	"container/list"
	"fmt"
)

//func SonarContactsExamples() {
func main() {
	constellixSonar := sonar.Init("b819f051-fb78-423c-bd7a-242982b52fad", "ae77965b-0aa3-4187-939e-f21be432f9b3")

	//-------------------------------------------------
	// get all contacts

	var contacts *list.List
	var err error
	contacts, err = constellixSonar.Contacts.GetContacts()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of Contacts: %d\n", contacts.Len())
		for e := contacts.Front(); e != nil; e = e.Next() {
			contact := e.Value.(sonar.Contact)
			fmt.Println(contact)
		}
	}

	//-------------------------------------------------
	// get all contact groups

	var groups *list.List
	var err1 error
	groups, err = constellixSonar.Contacts.GetContactGroups()
	if err1 != nil {
		fmt.Println("Error occured:")
		fmt.Println(err1)	
	} else {
		fmt.Println()
		fmt.Printf("Count of Contact Groups: %d\n", groups.Len())
		for e := groups.Front(); e != nil; e = e.Next() {
			group := e.Value.(sonar.ContactGroup)
			fmt.Println(group)
		}
	}
}
