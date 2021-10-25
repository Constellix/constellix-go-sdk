package main

import (
	"constellix.com/constellix/api/dns"
	"container/list"
	"fmt"
)

func TagsExamples() {
	constellixDns := dns.Init("b819f051-fb78-423c-bd7a-242982b52fad", "ae77965b-0aa3-4187-939e-f21be432f9b3")
	
	//-------------------------------------------------
	// get all tags

	var tags *list.List
	var err error
	tags, err = constellixDns.Tags.GetAll()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of tags: %d\n", tags.Len())
		for e := tags.Front(); e != nil; e = e.Next() {
			tag := e.Value.(dns.Tag)
			fmt.Println(tag)
		}
	}

	//-------------------------------------------------
	// create tag

	var createParam dns.TagParam
	createParam.Name = "Sample Tag"
	
	var newTagId int
	newTagId, err = constellixDns.Tags.Create(createParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Created Tag Id = %d\n", newTagId)
	}

	//-------------------------------------------------
	// get tag by id

	var newTag *dns.Tag
	newTag, err = constellixDns.Tags.GetTag(newTagId)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Tag:")
		fmt.Println(newTag)
	}

	//-------------------------------------------------
	// update tag

	var updateParam dns.TagParam
	updateParam.Name = "Update Tag"

	var updatedTag *dns.Tag
	updatedTag, err = newTag.Update(updateParam)

	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Updated Tag:")
		fmt.Println(updatedTag)
	}

	//-------------------------------------------------
	// delete tag

	err = updatedTag.Delete()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Tag with Id %d Deleted\n", updatedTag.Id)
	}
}
