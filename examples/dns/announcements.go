package main

import (
	"constellix.com/constellix/api/dns"
	"container/list"
	"fmt"
)

func AnnouncementsExamples() {
	constellixDns := dns.Init("", "")

	//-------------------------------------------------
	// get all announcements

	var announcements *list.List
	var err error
	announcements, err = constellixDns.Announcements.GetAll()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of announcements: %d\n", announcements.Len())
		for e := announcements.Front(); e != nil; e = e.Next() {
			announcement := e.Value.(dns.Announcement)
			fmt.Println(announcement)
		}
	}

	//-------------------------------------------------
	// get announcement by id

	var id int = announcements.Front().Value.(dns.Announcement).Id
	var newAnnouncement *dns.Announcement
	newAnnouncement, err = constellixDns.Announcements.GetAnnouncement(id)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Announcement:")
		fmt.Println(newAnnouncement)
	}
}
