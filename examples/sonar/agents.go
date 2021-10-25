package main

import (
	"constellix.com/constellix/api/sonar"
	"container/list"
	"fmt"
)

//func SonarAgentsExamples() {
func main() {
	constellixSonar := sonar.Init("b819f051-fb78-423c-bd7a-242982b52fad", "ae77965b-0aa3-4187-939e-f21be432f9b3")

	//-------------------------------------------------
	// get all Agents

	var agents *list.List
	var err error
	agents, err = constellixSonar.Agents.GetAgents()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of Agents: %d\n", agents.Len())
		for e := agents.Front(); e != nil; e = e.Next() {
			agent := e.Value.(sonar.Agent)
			fmt.Println(agent)
		}
	}

	//-------------------------------------------------
	// Sample HTTP Test on Agent
	
	var httpTestParam sonar.HTTPTestParam
	httpTestParam.Host = "www.google.com"
	httpTestParam.Port = 443
	httpTestParam.ExpectedStatusCode = 200
	httpTestParam.IpVersion = sonar.IPVERSION_IPV4
	httpTestParam.ProtocolType = sonar.PROTOCOLTYPE_HTTPS

	var httpTestResults *list.List
	agentForHTTPTest := agents.Front().Value.(sonar.Agent)
	httpTestResults, err = agentForHTTPTest.HTTPTest(httpTestParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of Results: %d\n", httpTestResults.Len())
		for e := httpTestResults.Front(); e != nil; e = e.Next() {
			result := e.Value.(sonar.HTTPTestResult)
			fmt.Println(result)
		}
	}

	//-------------------------------------------------
	// Sample Trace Test on Agent
	
	var traceTestParam sonar.TraceTestParam
	traceTestParam.Host = "www.google.com"
	traceTestParam.ProtocolType = sonar.PROTOCOLTYPE_UDP
	traceTestParam.IpVersion = sonar.IPVERSION_IPV4

	var traceTestResults *list.List
	agentForTraceTest := agents.Front().Value.(sonar.Agent)
	traceTestResults, err = agentForTraceTest.TraceTest(traceTestParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of Results: %d\n", traceTestResults.Len())
		for e := traceTestResults.Front(); e != nil; e = e.Next() {
			result := e.Value.(sonar.TraceTestResult)
			fmt.Println(result)
		}
	}

	//-------------------------------------------------
	// Sample TCP Test on Agent
	
	var tcpTestParam sonar.TCPTestParam
	tcpTestParam.Host = "www.google.com"
	tcpTestParam.Port = 80
	tcpTestParam.IpVersion = sonar.IPVERSION_IPV4

	var tcpTestResults *list.List
	agentForTcpTest := agents.Front().Value.(sonar.Agent)
	tcpTestResults, err = agentForTcpTest.TCPTest(tcpTestParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of Results: %d\n", tcpTestResults.Len())
		for e := tcpTestResults.Front(); e != nil; e = e.Next() {
			result := e.Value.(sonar.TCPTestResult)
			fmt.Println(result)
		}
	}

	//-------------------------------------------------
	// Sample DNS Test on Agent
	
	var dnsTestParam sonar.DNSTestParam
	dnsTestParam.Host = "www.google.com"
	dnsTestParam.NameServer = "ns2.google.com"
	dnsTestParam.RecordType = sonar.RECORDTYPE_A

	var dnsTestResults *list.List
	agentForDNSTest := agents.Front().Value.(sonar.Agent)
	dnsTestResults, err = agentForDNSTest.DNSTest(dnsTestParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of Results: %d\n", dnsTestResults.Len())
		for e := dnsTestResults.Front(); e != nil; e = e.Next() {
			result := e.Value.(sonar.DNSTestResult)
			fmt.Println(result)
		}
	}

}
