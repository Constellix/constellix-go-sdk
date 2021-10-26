package main

import (
	"constellix.com/constellix/api/sonar"
	"container/list"
	"fmt"
)

func SonarTCPChecksExamples() {
	constellixSonar := sonar.Init("", "")

	//-------------------------------------------------
	// get all TCP Checks

	var checks *list.List
	var err error
	checks, err = constellixSonar.TCPChecks.GetAll()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Count of TCP Checks: %d\n", checks.Len())
		for e := checks.Front(); e != nil; e = e.Next() {
			check := e.Value.(sonar.TCPCheck)
			fmt.Println(check)
		}
	}

	//-------------------------------------------------
	// create TCP Check

	var createParam sonar.TCPCheckParam
	createParam.Name = "Sample TCP Check"
	createParam.Host = "constellix.com"
	createParam.IpVersion = sonar.IPVERSION_IPV4
	createParam.Port = 80
	createParam.CheckSites = []int{1, 2}

	var newCheckId int
	newCheckId, err = constellixSonar.TCPChecks.Create(createParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Created TCP Check Id = %d\n", newCheckId)
	}

	//-------------------------------------------------
	// get TCP Check by id

	var newCheck *sonar.TCPCheck
	newCheck, err = constellixSonar.TCPChecks.GetCheck(newCheckId)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("TCP Check:")
		fmt.Println(newCheck)
	}

	//-------------------------------------------------
	// update TCP Check

	var updateParam sonar.TCPCheckParam
	updateParam.Name = "Sample TCP Check Update"
	updateParam.Port = 443
	updateParam.CheckSites = []int{1, 2, 3}

	var updatedCheck *sonar.TCPCheck
	updatedCheck, err = newCheck.Update(updateParam)

	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Updated TCP Check:")
		fmt.Println(updatedCheck)
	}

	//-------------------------------------------------
	// run TCP Check

	var checkResults *list.List
	checkResults, err = updatedCheck.RunCheck([]int{1, 2})
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("Count of Run TCP Checks results: %d\n", checkResults.Len())
		for r := checkResults.Front(); r != nil; r = r.Next() {
			result := r.Value.(sonar.TCPTestResult)
			fmt.Println(result)
		}
	}

	//-------------------------------------------------
	// run TCP Check Trace

	var traceResults *list.List
	traceResults, err = updatedCheck.RunCheckTrace([]int{1, 2})
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("Count of Run TCP Checks Trace results: %d\n", traceResults.Len())
		for r := traceResults.Front(); r != nil; r = r.Next() {
			result := r.Value.(sonar.TraceTestResult)
			fmt.Println(result)
		}
	}

	//-------------------------------------------------
	// Start TCP Check

	err = updatedCheck.Start()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("TCP Check with id=%d Stared\n", updatedCheck.Id)
	}

	//-------------------------------------------------
	// Stop TCP Check

	err = updatedCheck.Stop()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("TCP Check with id=%d Stopped\n", updatedCheck.Id)
	}

	//-------------------------------------------------
	// TCP Check State

	var checkState string
	checkState, err = updatedCheck.CheckState()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("TCP Check with id=%d State is %s\n", updatedCheck.Id, checkState)
	}

	//-------------------------------------------------
	// TCP Check Status

	var checkStatus string
	checkStatus, err = updatedCheck.CheckStatus()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("TCP Check with id=%d Status is %s\n", updatedCheck.Id, checkStatus)
	}

	//-------------------------------------------------
	// TCP Check Agent Status

	var agentStatus *list.List
	agentStatus, err = updatedCheck.CheckAgentStatus([]int{1, 2, 3})
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("TCP Check id %d Agent Status len is %d\n", updatedCheck.Id, agentStatus.Len())
		for e := agentStatus.Front(); e != nil; e = e.Next() {
			status := e.Value.(sonar.AgentStatus)
			fmt.Println(status)
		}
	}

	//-------------------------------------------------
	// delete TCP Check

	err = updatedCheck.Delete()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("TCP Check with Id %d Deleted\n", updatedCheck.Id)
	}
}
