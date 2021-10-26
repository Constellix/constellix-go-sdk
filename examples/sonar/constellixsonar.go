package main

import (
	"constellix.com/constellix/api/sonar"
)

func main() {

	//-------------------------------------------------
	// Creating Constellix Sonar object

	constellixSonar := sonar.Init("API-KEY", "SECRET-KEY")

	constellixSonar.Agents.GetAgents()			// Accessing Constellix Sonar Agents operations

	constellixSonar.Contacts.GetContacts()		// Accessing Constellix Sonar Contacts operations

	constellixSonar.DNSChecks.GetAll()			// Accessing Constellix Sonar DNS Checks operations

	constellixSonar.HTTPChecks.GetAll()			// Accessing Constellix Sonar HTTP Checks operations

	constellixSonar.TCPChecks.GetAll()			// Accessing Constellix Sonar TCP Checks operations
}