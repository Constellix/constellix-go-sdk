package main

import (
	"constellix.com/constellix/api/dns"
	"container/list"
	"fmt"
)

func TemplateRecordsExamples() {
	constellixDns := dns.Init("b819f051-fb78-423c-bd7a-242982b52fad", "ae77965b-0aa3-4187-939e-f21be432f9b3")

	//-------------------------------------------------
	// get all templates

	var templates *list.List
	var err error
	templates, err = constellixDns.Templates.GetAll()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of templates: %d\n", templates.Len())
		for e := templates.Front(); e != nil; e = e.Next() {
			template := e.Value.(dns.Template)
			fmt.Println(template)
			if template.Name == "Sample Template" {
				template.Delete()
			}
		}
	}

	//-------------------------------------------------
	// create template

	var createTemplateParam dns.TemplateParam
	createTemplateParam.Name = "Sample Template"
	createTemplateParam.GeoIp = true
	createTemplateParam.Gtd = true

	var newTemplateId int
	newTemplateId, err = constellixDns.Templates.Create(createTemplateParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Created Template Id = %d\n", newTemplateId)
	}

	//-------------------------------------------------
	// get template by id

	var newTemplate *dns.Template
	newTemplate, err = constellixDns.Templates.GetTemplate(newTemplateId)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Template:")
		fmt.Println(newTemplate)
	}

	//-------------------------------------------------
	// get all template records

	var templateRecords *list.List
	templateRecords, err = newTemplate.Records.GetAll()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of template recordss: %d\n", templateRecords.Len())
		for e := templateRecords.Front(); e != nil; e = e.Next() {
			templateRecord := e.Value.(dns.TemplateRecord)
			fmt.Println(templateRecord)
		}
	}

	//-------------------------------------------------
	// create template record

	var createParam dns.TemplateRecordParam
	createParam.Name = "template.com"
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

	var newTemplateRecordId int
	newTemplateRecordId, err = newTemplate.Records.Create(createParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Created Template Record Id = %d\n", newTemplateRecordId)
	}

	//-------------------------------------------------
	// get template record by id

	var newTemplateRecord *dns.TemplateRecord
	newTemplateRecord, err = newTemplate.Records.GetRecord(newTemplateRecordId)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Template Record:")
		fmt.Println(newTemplateRecord)
	}

	//-------------------------------------------------
	// update template record

	var updateParam dns.TemplateRecordParam
	updateParam.Region = dns.RECORDREGION_EUROPE
	updateParam.Mode = dns.RECORDMODE_STANDARD
	updateParam.ValueStandard = new(dns.RecordValueStandard)
	
	var vv1 dns.RecordValueShort
	vv1.Enabled = true
	vv1.Value = "192.168.0.1"
	updateParam.ValueStandard.Values = append(updateParam.ValueStandard.Values, vv1)

	var vv2 dns.RecordValueShort
	vv2.Enabled = true
	vv2.Value = "127.0.0.1"
	updateParam.ValueStandard.Values = append(updateParam.ValueStandard.Values, vv2)

	var updatedTemplateRecord *dns.TemplateRecord
	updatedTemplateRecord, err = newTemplateRecord.Update(updateParam)

	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Updated Template Record:")
		fmt.Println(updatedTemplateRecord)
	}

	//-------------------------------------------------
	// delete template record

	err = updatedTemplateRecord.Delete()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Template Record with Id %d Deleted\n", updatedTemplateRecord.Id)
	}

	newTemplate.Delete()
}
