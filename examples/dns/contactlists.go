package main

import (
	"constellix.com/constellix/api/dns"
	"container/list"
	"fmt"
)

//func ContactListsExamples() {
func main() {
	constellixDns := dns.Init("", "")

	//-------------------------------------------------
	// get all Contact Lists
	var contactLists *list.List
	var err error
	contactLists, err = constellixDns.ContactLists.GetAll()

	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of Contact Lists: %d\n", contactLists.Len())
		for e := contactLists.Front(); e != nil; e = e.Next() {
			contactList := e.Value.(dns.ContactList)
			fmt.Println(contactList)
		}
	}

	//-------------------------------------------------
	// create Contact List

	var createParam dns.ContactListParam
	createParam.Name = "Sample Contact List"
	createParam.Emails = []string{"sample@mail.com", "sample@dot.com"}

	var newContactListId int
	newContactListId, err = constellixDns.ContactLists.Create(createParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Created Contact List Id = %d\n", newContactListId)
	}

	//-------------------------------------------------
	// get Contact List by id

	var newContactList *dns.ContactList
	newContactList, err = constellixDns.ContactLists.GetContactList(newContactListId)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("ContactList:")
		fmt.Println(newContactList)
	}

	//-------------------------------------------------
	// update Contact List

	var updateParam dns.ContactListParam
	updateParam.Name = "Sample Contact List Update"
	updateParam.Emails = []string{"sample@mail.com", "sample@dot.com", "onemore@mail.com"}

	var updatedContactList *dns.ContactList
	updatedContactList, err = newContactList.Update(updateParam)

	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Updated Contact List:")
		fmt.Println(updatedContactList)
	}

	//-------------------------------------------------
	// delete Contact List

	err = updatedContactList.Delete()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Contact List with Id %d Deleted\n", updatedContactList.Id)
	}
}
