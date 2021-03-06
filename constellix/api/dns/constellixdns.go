package dns

import (
	"os"
	"constellix.com/constellix/api"
)

type ConstellixDns struct {
	Domains 				Domains
	Templates 				Templates
	Pools 					Pools
	IpFilters				IpFilters
	GeoProximityLocations	GeoProximityLocations
	VanityNameservers		VanityNameservers
	ContactLists			ContactLists
	Tags					Tags
	Announcements			Announcements
}

func Init(apiKey, secretKey string) (*ConstellixDns) {
	if apiKey == "" && secretKey == "" {
		apiKey = os.Getenv("CONSTELLIX_API_KEY")
		secretKey = os.Getenv("CONSTELLIX_SECRET_KEY")
	}

	constellixDns := &ConstellixDns{}
	_ = api.GetDnsApiClient(apiKey, secretKey)
	return constellixDns
} 