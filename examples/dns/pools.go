package main

import (
	"constellix.com/constellix/api/dns"
	"container/list"
	"fmt"
)

//func PoolsExamples() {
func main() {
	constellixDns := dns.Init("b819f051-fb78-423c-bd7a-242982b52fad", "ae77965b-0aa3-4187-939e-f21be432f9b3")
	
	//-------------------------------------------------
	// get all pools

	var pools *list.List
	var err error
	pools, err = constellixDns.Pools.GetAll()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Printf("Count of pools: %d\n", pools.Len())
	for e := pools.Front(); e != nil; e = e.Next() {
		pool := e.Value.(dns.Pool)
		fmt.Println(pool)
		if pool.Name == "Sample pool" {
			pool.Delete()
		}
	}

	//-------------------------------------------------
	// create pool

	var createParam dns.PoolParam
	createParam.Name = "Sample pool"
	createParam.Type = dns.POOLTYPE_A
	createParam.Return = 1
	createParam.Enabled = true

	var poolValue dns.PoolValue
	poolValue.Value = "198.51.100.42"
	poolValue.Weight = 1000
	poolValue.Enabled = true
	poolValue.Handicap = 10
	poolValue.Policy = dns.POOLVALUEPOLICY_ALWAYS_OFF

	createParam.Values = []dns.PoolValue {poolValue}

	var newPoolId int
	newPoolId, err = constellixDns.Pools.Create(createParam)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Created Pool Id = %d\n", newPoolId)
	}

	//-------------------------------------------------
	// get pool by id

	var newPool *dns.Pool
	newPool, err = constellixDns.Pools.GetPool(dns.POOLTYPE_A, newPoolId)
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Pool:")
		fmt.Println(newPool)
	}

	//-------------------------------------------------
	// update pool

	var updateParam dns.PoolParam
	updateParam.Name = "Sample Pool Update"
	updateParam.Type = dns.POOLTYPE_A
	updateParam.Return = 1
	updateParam.Enabled = false

	var updatedPool *dns.Pool
	updatedPool, err = newPool.Update(updateParam)

	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Updated Pool:")
		fmt.Println(updatedPool)
	}

	//-------------------------------------------------
	// delete pool

	err = updatedPool.Delete()
	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Printf("Pool with Id %d Deleted\n", updatedPool.Id)
	}
}
