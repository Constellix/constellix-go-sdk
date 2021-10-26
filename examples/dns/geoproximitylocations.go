package main

import (
	"constellix.com/constellix/api/dns"
	"container/list"
	"fmt"
)

func GeoProximityLocationsExamples() {
	constellixDns := dns.Init("", "")

	//-------------------------------------------------
	// get all geoProximityLocations

	var geoProximityLocations *list.List
	var err error
	geoProximityLocations, err = constellixDns.GeoProximityLocations.GetAll()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of Geo Proximity Locations: %d\n", geoProximityLocations.Len())
		for e := geoProximityLocations.Front(); e != nil; e = e.Next() {
			geoProximityLocation := e.Value.(dns.GeoProximityLocation)
			fmt.Println(geoProximityLocation)
			if geoProximityLocation.Name == "Sample Geo Proximity Location"{
				geoProximityLocation.Delete()
			}
		}
	}

	//-------------------------------------------------
	// create geoProximityLocation

	var createParam dns.GeoProximityLocationParam
	createParam.Name = "Sample Geo Proximity Location"
	createParam.Longitude = 22.7
	createParam.Latitude = 56.8333

	var newGeoProximityLocationId int
	newGeoProximityLocationId, err = constellixDns.GeoProximityLocations.Create(createParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Created Geo Proximity Location Id = %d\n", newGeoProximityLocationId)
	}

	//-------------------------------------------------
	// get geoProximityLocation by id

	var newGeoProximityLocation *dns.GeoProximityLocation
	newGeoProximityLocation, err = constellixDns.GeoProximityLocations.GetGeoProximityLocation(newGeoProximityLocationId)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Geo Proximity Location:")
		fmt.Println(newGeoProximityLocation)
	}

	//-------------------------------------------------
	// update geoProximityLocation

	var updateParam dns.GeoProximityLocationParam
	updateParam.Country = "GB"
	//updateParam.City = "London"

	var updatedGeoProximityLocation *dns.GeoProximityLocation
	updatedGeoProximityLocation, err = newGeoProximityLocation.Update(updateParam)

	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Updated Geo Proximity Location:")
		fmt.Println(updatedGeoProximityLocation)
	}

	//-------------------------------------------------
	// delete geoProximityLocation

	err = updatedGeoProximityLocation.Delete()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("GeoProximityLocation with Id %d Deleted\n", updatedGeoProximityLocation.Id)
	}
}
