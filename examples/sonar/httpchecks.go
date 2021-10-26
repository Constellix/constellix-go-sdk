package main

import (
	"constellix.com/constellix/api/sonar"
	"container/list"
	"fmt"
)

func SonarHTTPChecksExamples() {
	constellixSonar := sonar.Init("", "")

	//-------------------------------------------------
	// get all HTTP Checks

	var checks *list.List
	var err error
	checks, err = constellixSonar.HTTPChecks.GetAll()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of HTTP Checks: %d\n", checks.Len())
		for e := checks.Front(); e != nil; e = e.Next() {
			check := e.Value.(sonar.HTTPCheck)
			fmt.Println(check)
		}
	}

	//-------------------------------------------------
	// create HTTP Check

	var createParam sonar.HTTPCheckParam
	createParam.Name = "Sample HTTP Check"
	createParam.Host = "constellix.com"
	createParam.IpVersion = sonar.IPVERSION_IPV4
	createParam.Port = 80
	createParam.ProtocolType = sonar.PROTOCOLTYPE_HTTPS
	createParam.CheckSites = []int{1, 2}

	var newCheckId int
	newCheckId, err = constellixSonar.HTTPChecks.Create(createParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Created HTTP Check Id = %d\n", newCheckId)
	}

	//-------------------------------------------------
	// get HTTP Check by id

	var newCheck *sonar.HTTPCheck
	newCheck, err = constellixSonar.HTTPChecks.GetCheck(newCheckId)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("HTTP Check:")
		fmt.Println(newCheck)
	}

	//-------------------------------------------------
	// update HTTP Check

	var updateParam sonar.HTTPCheckParam
	updateParam.Name = "Sample HTTP Check Update"
	updateParam.Port = 443
	updateParam.ProtocolType = sonar.PROTOCOLTYPE_HTTP
	updateParam.CheckSites = []int{1, 2, 3}

	var updatedCheck *sonar.HTTPCheck
	updatedCheck, err = newCheck.Update(updateParam)

	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Updated HTTP Check:")
		fmt.Println(updatedCheck)
	}

	//-------------------------------------------------
	// run HTTP Check

	var checkResults *list.List
	checkResults, err = updatedCheck.RunCheck([]int{1, 2})
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("Count of Run HTTP Checks results: %d\n", checkResults.Len())
		for r := checkResults.Front(); r != nil; r = r.Next() {
			result := r.Value.(sonar.HTTPTestResult)
			fmt.Println(result)
		}
	}

	//-------------------------------------------------
	// run HTTP Check Trace

	var traceResults *list.List
	traceResults, err = updatedCheck.RunCheckTrace([]int{1, 2})
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("Count of Run HTTP Checks Trace results: %d\n", traceResults.Len())
		for r := traceResults.Front(); r != nil; r = r.Next() {
			result := r.Value.(sonar.TraceTestResult)
			fmt.Println(result)
		}
	}

	//-------------------------------------------------
	// Start HTTP Check

	err = updatedCheck.Start()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("HTTP Check with id=%d Stared\n", updatedCheck.Id)
	}

	//-------------------------------------------------
	// Stop HTTP Check

	err = updatedCheck.Stop()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("HTTP Check with id=%d Stopped\n", updatedCheck.Id)
	}

	//-------------------------------------------------
	// HTTP Check State

	var checkState string
	checkState, err = updatedCheck.CheckState()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("HTTP Check with id=%d State is %s\n", updatedCheck.Id, checkState)
	}

	//-------------------------------------------------
	// HTTP Check Status

	var checkStatus string
	checkStatus, err = updatedCheck.CheckStatus()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("HTTP Check with id=%d Status is %s\n", updatedCheck.Id, checkStatus)
	}

	//-------------------------------------------------
	// HTTP Check Agent Status

	var agentStatus *list.List
	agentStatus, err = updatedCheck.CheckAgentStatus([]int{1, 2, 3})
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("HTTP Check id %d Agent Status len is %d\n", updatedCheck.Id, agentStatus.Len())
		for e := agentStatus.Front(); e != nil; e = e.Next() {
			status := e.Value.(sonar.AgentStatus)
			fmt.Println(status)
		}
	}

	//-------------------------------------------------
	// delete HTTP Check

	err = updatedCheck.Delete()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("HTTP Check with Id %d Deleted\n", updatedCheck.Id)
	}
}
