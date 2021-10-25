package main

import (
	"constellix.com/constellix/api/dns"
	"container/list"
	"fmt"
)

func DomainSnapshotsExamples() {
	constellixDns := dns.Init("b819f051-fb78-423c-bd7a-242982b52fad", "ae77965b-0aa3-4187-939e-f21be432f9b3")
	
	var createParam dns.DomainParam
	createParam.Name = "goexample-testdomain.com"
	createParam.Soa.PrimaryNameServer = "ns11.constellix.com"
	createParam.Soa.Email = "admin.example.com"
	createParam.Soa.Ttl = 86400
	createParam.Soa.Refresh = 86400
	createParam.Soa.Retry = 7200
	createParam.Soa.Expire = 3600000
	createParam.Soa.NegativeCache = 180
	createParam.Note = "Sample Domain"
	createParam.GeoIp = true
	createParam.Gtd = true

	var newDomainId int
	var err error
	newDomainId, err = constellixDns.Domains.Create(createParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	}

	var newDomain *dns.Domain
	newDomain, err = constellixDns.Domains.GetDomain(newDomainId)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	}

	//-------------------------------------------------
	// get domain snapshots

	var domainSnapshots *list.List
	domainSnapshots, err = newDomain.Snapshots.GetSnapshots()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of domain snapshots: %d\n", domainSnapshots.Len())
		for e := domainSnapshots.Front(); e != nil; e = e.Next() {
			domainSnapshot := e.Value.(dns.DomainSnapshot)
			fmt.Println(domainSnapshot)
		}
	}

	//-------------------------------------------------
	// get domain snapshot version

	var domainSnapshot *dns.DomainSnapshot
	domainSnapshot, err = newDomain.Snapshots.GetSnapshot(2)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println("Domain Snapshot:")
		fmt.Println(domainSnapshot)
	}

	//-------------------------------------------------
	// apply domain snapshot of version

	err = newDomain.Snapshots.ApplySnapshot(2)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println("Snapshot applied")
	}

	//-------------------------------------------------
	// delete domain snapshot

	err = domainSnapshot.Delete()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Printf("Snapshot deleted of version %d\n", domainSnapshot.Version)
	}

	newDomain.Delete()
}