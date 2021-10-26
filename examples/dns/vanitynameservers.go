package main

import (
	"constellix.com/constellix/api/dns"
	"container/list"
	"fmt"
)

func VanityNameserversExamples() {
	constellixDns := dns.Init("", "")
	
	//-------------------------------------------------
	// get all vanityNameservers

	var vanityNameservers *list.List
	var err error
	vanityNameservers, err = constellixDns.VanityNameservers.GetAll()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of vanityNameservers: %d\n", vanityNameservers.Len())
		for e := vanityNameservers.Front(); e != nil; e = e.Next() {
			vanityNameserver := e.Value.(dns.VanityNameserver)
			fmt.Println(vanityNameserver)

			if vanityNameserver.Name == "Sample Vanity Nameserver" {
				vanityNameserver.Delete()
			}
		}
	}

	//-------------------------------------------------
	// create vanityNameserver

	var createParam dns.VanityNameserverParam
	createParam.Name = "Sample Vanity Nameserver"
	createParam.Default = false
	createParam.NameserverGroup.Id = 1
	createParam.NameserverGroup.Name = "Nameserver Group 1"
	createParam.Nameservers = []string{"ns1.example.com", "ns2.example.com"}

	var newVanityNameserverId int
	newVanityNameserverId, err = constellixDns.VanityNameservers.Create(createParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Created VanityNameserver Id = %d\n", newVanityNameserverId)
	}

	//-------------------------------------------------
	// get Vanity Nameserver by id

	var newVanityNameserver *dns.VanityNameserver
	newVanityNameserver, err = constellixDns.VanityNameservers.GetVanityNameserver(newVanityNameserverId)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("VanityNameserver:")
		fmt.Println(newVanityNameserver)
	}

	//-------------------------------------------------
	// update Vanity Nameserver

	var updateParam dns.VanityNameserverParam
	updateParam.Name = "Sample Vanity Nameserver Update"
	updateParam.NameserverGroup.Id = 1
	updateParam.NameserverGroup.Name = "Nameserver Group 1"
	updateParam.Default = true

	var updatedVanityNameserver *dns.VanityNameserver
	updatedVanityNameserver, err = newVanityNameserver.Update(updateParam)

	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Updated VanityNameserver:")
		fmt.Println(updatedVanityNameserver)
	}

	//-------------------------------------------------
	// delete Vanity Nameserver

	err = updatedVanityNameserver.Delete()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("VanityNameserver with Id %d Deleted\n", updatedVanityNameserver.Id)
	}
}
