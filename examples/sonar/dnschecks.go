package main

import (
	"constellix.com/constellix/api/sonar"
	"container/list"
	"fmt"
)

func SonarDNSChecksExamples() {
	constellixSonar := sonar.Init("", "")

	//-------------------------------------------------
	// get all DNS Checks

	var checks *list.List
	var err error
	checks, err = constellixSonar.DNSChecks.GetAll()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)	
	} else {
		fmt.Println()
		fmt.Printf("Count of DNS Checks: %d\n", checks.Len())
		for e := checks.Front(); e != nil; e = e.Next() {
			check := e.Value.(sonar.DNSCheck)
			fmt.Println(check)
		}
	}

	//-------------------------------------------------
	// create DNS Check

	var createParam sonar.DNSCheckParam
	createParam.Name = "Sample DNS Check"
	createParam.Fqdn = "example.com"
	createParam.Port = 53
	createParam.Resolver = "8.8.8.8"
	createParam.ResolverIPVersion = sonar.IPVERSION_IPV4
	createParam.ScheduleInterval = sonar.SCHEDULEINTERVAL_NONE
	createParam.RecordType = sonar.RECORDTYPE_A
	createParam.QueryProtocol = sonar.DNSQUERYPROTOCOL_UDP
	createParam.CompareOptions = sonar.DNSCOMPAREOPTION_ANYMATCH
	createParam.CheckSites = []int{1, 2}
	createParam.Dnssec = false
	createParam.Interval = sonar.MONITORINGINTERVAL_THIRTYSECONDS
	createParam.MonitorIntervalPolicy = sonar.MONITORINTERVALPOLICY_PARALLEL
	createParam.NotificationReportTimeout = 1440
	createParam.VerificationPolicy = sonar.VERIFICATIONPOLICY_SIMPLE

	var newCheckId int
	newCheckId, err = constellixSonar.DNSChecks.Create(createParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Created DNS Check Id = %d\n", newCheckId)
	}

	var checkType string
	checkType, err = constellixSonar.CheckType(newCheckId)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Created DNS Check Type = %s\n", checkType)
	}

	//-------------------------------------------------
	// get DNS Check by id

	var newCheck *sonar.DNSCheck
	newCheck, err = constellixSonar.DNSChecks.GetCheck(newCheckId)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("DNS Check:")
		fmt.Println(newCheck)
	}

	//-------------------------------------------------
	// update DNS Check

	var updateParam sonar.DNSCheckParam
	updateParam.Name = "Sample DNS Check Update"
	updateParam.QueryProtocol = sonar.DNSQUERYPROTOCOL_TCP
	updateParam.CheckSites = []int{1, 2, 3}
	updateParam.ResolverIPVersion = sonar.IPVERSION_IPV4
	updateParam.RecordType = sonar.RECORDTYPE_A
	updateParam.CompareOptions = sonar.DNSCOMPAREOPTION_ANYMATCH
	updateParam.Dnssec = false
	updateParam.Interval = sonar.MONITORINGINTERVAL_THIRTYSECONDS
	updateParam.MonitorIntervalPolicy = sonar.MONITORINTERVALPOLICY_PARALLEL

	var updatedCheck *sonar.DNSCheck
	updatedCheck, err = newCheck.Update(updateParam)

	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Updated DNS Check:")
		fmt.Println(updatedCheck)
	}

	//-------------------------------------------------
	// run DNS Check

	var checkResults *list.List
	checkResults, err = updatedCheck.RunCheck([]int{1, 2})
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("Count of Run DNS Checks results: %d\n", checkResults.Len())
		for r := checkResults.Front(); r != nil; r = r.Next() {
			result := r.Value.(sonar.DNSTestResult)
			fmt.Println(result)
		}
	}

	//-------------------------------------------------
	// run DNS Check Trace

	var traceResults *list.List
	traceResults, err = updatedCheck.RunCheckTrace([]int{1, 2})
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("Count of Run DNS Checks Trace results: %d\n", traceResults.Len())
		for r := traceResults.Front(); r != nil; r = r.Next() {
			result := r.Value.(sonar.TraceTestResult)
			fmt.Println(result)
		}
	}

	//-------------------------------------------------
	// Start DNS Check

	err = updatedCheck.Start()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("DNS Check with id=%d Stared\n", updatedCheck.Id)
	}

	//-------------------------------------------------
	// Stop DNS Check

	err = updatedCheck.Stop()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("DNS Check with id=%d Stopped\n", updatedCheck.Id)
	}

	//-------------------------------------------------
	// DNS Check State

	var checkState string
	checkState, err = updatedCheck.CheckState()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("DNS Check with id=%d State is %s\n", updatedCheck.Id, checkState)
	}

	//-------------------------------------------------
	// DNS Check Status

	var checkStatus string
	checkStatus, err = updatedCheck.CheckStatus()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("DNS Check with id=%d Status is %s\n", updatedCheck.Id, checkStatus)
	}

	//-------------------------------------------------
	// DNS Check Agent Status

	var agentStatus *list.List
	agentStatus, err = updatedCheck.CheckAgentStatus([]int{1, 2, 3})
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Printf("DNS Check id %d Agent Status len is %d\n", updatedCheck.Id, agentStatus.Len())
		for e := agentStatus.Front(); e != nil; e = e.Next() {
			status := e.Value.(sonar.AgentStatus)
			fmt.Println(status)
		}
	}

	//-------------------------------------------------
	// delete DNS Check

	err = updatedCheck.Delete()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("DNS Check with Id %d Deleted\n", updatedCheck.Id)
	}
}
