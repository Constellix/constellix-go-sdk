package main

import (
	"constellix.com/constellix/api/dns"
	"container/list"
	"fmt"
)

func DomainRecordsExamples() {
	constellixDns := dns.Init("", "")
	
	//-------------------------------------------------
	// get all domains

	var domains *list.List
	var errd error
	domains, errd = constellixDns.Domains.GetAll()
	if errd != nil {
		fmt.Println("Error occured:")
		fmt.Println(errd)	
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

	var createDomainParam dns.DomainParam
	createDomainParam.Name = "goexample-testdomain.com"
	createDomainParam.Soa.PrimaryNameServer = "ns11.constellix.com"
	createDomainParam.Soa.Email = "admin.example.com"
	createDomainParam.Soa.Ttl = 86400
	createDomainParam.Soa.Refresh = 86400
	createDomainParam.Soa.Retry = 7200
	createDomainParam.Soa.Expire = 3600000
	createDomainParam.Soa.NegativeCache = 180
	createDomainParam.Note = "Sample Domain"
	createDomainParam.GeoIp = true
	createDomainParam.Gtd = true

	var newDomainId int
	var err error
	newDomainId, err = constellixDns.Domains.Create(createDomainParam)
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

	/*createParam.Value1 = new(dns.RecordValue1)
	
	var v1 dns.RecordValueShort
	v1.Enabled = true
	v1.Value = "192.168.0.1"
	createParam.Value1.Values = append(createParam.Value1.Values, v1)

	var v2 dns.RecordValueShort
	v2.Enabled = true
	v2.Value = "127.0.0.1"
	createParam.Value1.Values = append(createParam.Value1.Values, v2)*/

	/*createParam.Value2 = new(dns.RecordValue2)
	createParam.Value2.Values.Mode = dns.RECORDVALUEMODE_NORMAL
	var v1 dns.RecordValueExtended
	v1.Enabled = true
	v1.Order = 10
	v1.SonarCheckId = 101
	v1.Value = "192.168.0.1"
	createParam.Value2.Values.Values = append(createParam.Value2.Values.Values, v1)

	var v2 dns.RecordValueExtended
	v2.Enabled = true
	v2.Order = 20
	v2.SonarCheckId = 202
	v2.Value = "127.0.0.1"
	createParam.Value2.Values.Values = append(createParam.Value2.Values.Values, v2)*/

	/*createParam.Value3 = new(dns.RecordValue3)
	
	var v1 dns.RecordValueExtended
	v1.Enabled = true
	v1.Order = 10
	v1.SonarCheckId = 101
	v1.Value = "192.168.0.1"
	createParam.Value3.Values = append(createParam.Value3.Values, v1)

	var v2 dns.RecordValueExtended
	v2.Enabled = true
	v2.Order = 20
	v2.SonarCheckId = 202
	v2.Value = "127.0.0.1"
	createParam.Value3.Values = append(createParam.Value3.Values, v2)*/

	/*createParam.Value4 = new(dns.RecordValue4)
	createParam.Value4.Values = append(createParam.Value4.Values, 101)
	createParam.Value4.Values = append(createParam.Value4.Values, 102)
	createParam.Value4.Values = append(createParam.Value4.Values, 103)*/


	//-------------------------------------------------
	// get all domain records

	var domainRecords *list.List
	domainRecords, err = newDomain.Records.GetAll()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else { 
		fmt.Println()
		fmt.Printf("Count of domain recordss: %d\n", domainRecords.Len())
		for e := domainRecords.Front(); e != nil; e = e.Next() {
			domainRecord := e.Value.(dns.DomainRecord)
			fmt.Println(domainRecord)
		}
	}

	//-------------------------------------------------
	// create domain record

	var createParam dns.DomainRecordParam
	createParam.Name = "domain.com"
	createParam.Type = dns.RECORDTYPE_A
	createParam.Region = dns.RECORDREGION_EUROPE
	createParam.Mode = dns.RECORDMODE_STANDARD
	createParam.ValueStandard = new(dns.RecordValueStandard)
	
	var v1 dns.RecordValueShort
	v1.Enabled = true
	v1.Value = "192.168.0.1"
	createParam.ValueStandard.Values = append(createParam.ValueStandard.Values, v1)

	var v2 dns.RecordValueShort
	v2.Enabled = true
	v2.Value = "127.0.0.1"
	createParam.ValueStandard.Values = append(createParam.ValueStandard.Values, v2)

	var newDomainRecordId int
	newDomainRecordId, err = newDomain.Records.Create(createParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Created Domain Record Id = %d\n", newDomainRecordId)
	}

	//-------------------------------------------------
	// get domain record by id

	var newDomainRecord *dns.DomainRecord
	newDomainRecord, err = newDomain.Records.GetRecord(newDomainRecordId)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Domain Record:")
		fmt.Println(newDomainRecord)
	}

	//-------------------------------------------------
	// update domain record

	var updateParam dns.DomainRecordParam
	updateParam.Region = dns.RECORDREGION_EUROPE
	updateParam.ValueStandard = new(dns.RecordValueStandard)
	
	var vv1 dns.RecordValueShort
	vv1.Enabled = true
	vv1.Value = "192.168.0.1"
	updateParam.ValueStandard.Values = append(updateParam.ValueStandard.Values, vv1)

	var vv2 dns.RecordValueShort
	vv2.Enabled = true
	vv2.Value = "127.0.0.1"
	updateParam.ValueStandard.Values = append(updateParam.ValueStandard.Values, vv2)

	var updatedDomainRecord *dns.DomainRecord
	updatedDomainRecord, err = newDomainRecord.Update(updateParam)

	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Updated Domain Record:")
		fmt.Println(updatedDomainRecord)
	}

	//-------------------------------------------------
	// delete domain record

	err = updatedDomainRecord.Delete()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Domain Record with Id %d Deleted\n", updatedDomainRecord.Id)
	}
	
	//
	newDomain.Delete()
}
