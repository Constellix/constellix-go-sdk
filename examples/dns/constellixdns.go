package main

import (
	"constellix.com/constellix/api/dns"
)

func main() {

	//-------------------------------------------------
	// Creating Constellix DNS object

	constellixDns := dns.Init("API-KEY", "SECRET-KEY")

	constellixDns.Domains.GetAll()					// Accessing Constellix DNS Domains operations

	constellixDns.Templates.GetAll()				// Accessing Constellix DNS Templates operations

	constellixDns.Pools.GetAll()					// Accessing Constellix DNS Pools operations

	constellixDns.IpFilters.GetAll()				// Accessing Constellix DNS IP Filters operations

	constellixDns.GeoProximityLocations.GetAll()	// Accessing Constellix DNS Geo Proximity Locations operations

	constellixDns.VanityNameservers.GetAll()		// Accessing Constellix DNS Vanity Nameservers operations
	
	constellixDns.ContactLists.GetAll()				// Accessing Constellix DNS Contact Lists operations

	constellixDns.Tags.GetAll()						// Accessing Constellix DNS Tags operations

	constellixDns.Announcements.GetAll()			// Accessing Constellix DNS Announcements operations
}