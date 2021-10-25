package main

import (
	"constellix.com/constellix/api/dns"
	"container/list"
	"fmt"
)

//func DomainsExamples() {
func main() {
	constellixDns := dns.Init("b819f051-fb78-423c-bd7a-242982b52fad", "ae77965b-0aa3-4187-939e-f21be432f9b3")

	//-------------------------------------------------
	// get all domains

	var domains *list.List
	var err error
	domains, err = constellixDns.Domains.GetAll()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of domains: %d\n", domains.Len())
		for e := domains.Front(); e != nil; e = e.Next() {
			domain := e.Value.(dns.Domain)
			if domain.Name == "goexample-testdomain.com" {
				domain.Delete()
			}
			fmt.Println(domain)
		}
	}

	//-------------------------------------------------
	// create domain

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
	newDomainId, err = constellixDns.Domains.Create(createParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Created Domain Id = %d\n", newDomainId)
	}

	//-------------------------------------------------
	// get domain by id

	var newDomain *dns.Domain
	newDomain, err = constellixDns.Domains.GetDomain(newDomainId)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Domain:")
		fmt.Println(newDomain)
	}

	//-------------------------------------------------
	// update domain

	var updateParam dns.DomainParam
	updateParam.Soa.PrimaryNameServer = "ns12.constellix.com"
	updateParam.Soa.Email = "admin.example.com"
	updateParam.Soa.Ttl = 3600
	updateParam.Note = "Sample Domain Updated"
	updateParam.GeoIp = false
	updateParam.Gtd = false

	var updatedDomain *dns.Domain
	updatedDomain, err = newDomain.Update(updateParam)

	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Updated Domain:")
		fmt.Println(updatedDomain)
	}

	//-------------------------------------------------
	// delete domain

	err = updatedDomain.Delete()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Domain with Id %d Deleted\n", updatedDomain.Id)
	}
}
