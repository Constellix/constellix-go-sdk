package main

import (
	"constellix.com/constellix/api/dns"
	"container/list"
	"fmt"
)

func TemplatesExamples() {
	constellixDns := dns.Init("", "")
	
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
		}
	}

	//-------------------------------------------------
	// create template

	var createParam dns.TemplateParam
	createParam.Name = "Sample Template"
	createParam.GeoIp = true
	createParam.Gtd = true

	var newTemplateId int
	newTemplateId, err = constellixDns.Templates.Create(createParam)
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
	// update template

	var updateParam dns.TemplateParam
	updateParam.Name = "Sample Template Updated"
	updateParam.GeoIp = false
	updateParam.Gtd = false

	var updatedTemplate *dns.Template
	updatedTemplate, err = newTemplate.Update(updateParam)

	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Updated Template:")
		fmt.Println(updatedTemplate)
	}

	//-------------------------------------------------
	// delete template

	err = updatedTemplate.Delete()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Template with Id %d Deleted\n", updatedTemplate.Id)
	}
}
