package main

import (
	"constellix.com/constellix/api/dns"
	"container/list"
	"fmt"
)

func IpFiltersExamples() {
	constellixDns := dns.Init("", "")
	
	//-------------------------------------------------
	// get all ipFilters

	var ipFilters *list.List
	var err error
	ipFilters, err = constellixDns.IpFilters.GetAll()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of ipFilters: %d\n", ipFilters.Len())
		for e := ipFilters.Front(); e != nil; e = e.Next() {
			ipFilter := e.Value.(dns.IpFilter)
			fmt.Println(ipFilter)
			if ipFilter.Name == "Sample IP Filter" {
				ipFilter.Delete()
			}
		}
	}

	//-------------------------------------------------
	// create ipFilter

	var createParam dns.IpFilterParam
	createParam.Name = "Sample IP Filter"
	createParam.RulesLimit = 100
	createParam.Continents = []dns.Continent{dns.CONTINENT_EU, dns.CONTINENT_NA}
	createParam.Countries = []string{ "GB", "PL", "RO"}
	createParam.Ipv4 = []string{"198.51.100.0/24", "203.0.113.42"}
	createParam.Ipv6 = []string{"2001:db8:200::/64", "2001:db8:200:42::"}
	
	var newIpFilterId int
	newIpFilterId, err = constellixDns.IpFilters.Create(createParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Created IpFilter Id = %d\n", newIpFilterId)
	}

	//-------------------------------------------------
	// get ipFilter by id

	var newIpFilter *dns.IpFilter
	newIpFilter, err = constellixDns.IpFilters.GetIpFilter(newIpFilterId)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("IpFilter:")
		fmt.Println(newIpFilter)
	}

	//-------------------------------------------------
	// update ipFilter

	var updateParam dns.IpFilterParam
	updateParam.Name = "Sample IP Filter Update"
	updateParam.RulesLimit = 200
	updateParam.Continents = []dns.Continent{dns.CONTINENT_EU, dns.CONTINENT_NA, dns.CONTINENT_AF}

	var updatedIpFilter *dns.IpFilter
	updatedIpFilter, err = newIpFilter.Update(updateParam)

	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Updated IpFilter:")
		fmt.Println(updatedIpFilter)
	}

	//-------------------------------------------------
	// delete ipFilter

	err = updatedIpFilter.Delete()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("IpFilter with Id %d Deleted\n", updatedIpFilter.Id)
	}
}
