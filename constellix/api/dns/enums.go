package dns

import (
	"bytes"
	"encoding/json"
)

//=====================================================
// Record Type
type RecordType int

const (
	RECORDTYPE_NONE RecordType = iota
	RECORDTYPE_A
	RECORDTYPE_AAAA
	RECORDTYPE_ANAME
	RECORDTYPE_CAA
	RECORDTYPE_CERT
	RECORDTYPE_CNAME
	RECORDTYPE_HINFO
	RECORDTYPE_HTTP
	RECORDTYPE_MX
	RECORDTYPE_NAPTR
	RECORDTYPE_NS
	RECORDTYPE_PTR
	RECORDTYPE_RP
	RECORDTYPE_SPF
	RECORDTYPE_SRV
	RECORDTYPE_TXT
)

func (s RecordType) String() string {
	return toStringRT[s]
}

var toStringRT = map[RecordType]string{
	RECORDTYPE_NONE:		"NONE",
	RECORDTYPE_AAAA:		"AAAA",
	RECORDTYPE_A:			"A", 
	RECORDTYPE_ANAME:		"ANAME",
	RECORDTYPE_CAA:			"CAA",
	RECORDTYPE_CERT:		"CERT",
	RECORDTYPE_CNAME:		"CNAME",
	RECORDTYPE_HINFO:		"HINFO",
	RECORDTYPE_HTTP:		"HTTP",
	RECORDTYPE_MX:			"MX",
	RECORDTYPE_NAPTR:		"NAPTR",
	RECORDTYPE_NS:			"NS",
	RECORDTYPE_PTR:			"PTR",
	RECORDTYPE_RP:			"RP",
	RECORDTYPE_SPF:			"SPF",
	RECORDTYPE_SRV:			"SRV",
	RECORDTYPE_TXT:			"TXT",
}

var toIDRT = map[string]RecordType{
	"":			RECORDTYPE_NONE,
	"A":		RECORDTYPE_A, 
	"AAAA":		RECORDTYPE_AAAA,
	"ANAME":	RECORDTYPE_ANAME,
	"CAA":		RECORDTYPE_CAA,
	"CERT":		RECORDTYPE_CERT,
	"CNAME":	RECORDTYPE_CNAME,
	"HINFO":	RECORDTYPE_HINFO,
	"HTTP":		RECORDTYPE_HTTP,
	"MX":		RECORDTYPE_MX,
	"NAPTR":	RECORDTYPE_NAPTR,
	"NS":		RECORDTYPE_NS,
	"PTR":		RECORDTYPE_PTR,
	"RP":		RECORDTYPE_RP,
	"SPF":		RECORDTYPE_SPF,
	"SRV":		RECORDTYPE_SRV,
	"TXT":		RECORDTYPE_TXT,
}

// MarshalJSON marshals the enum as a quoted json string
func (s RecordType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringRT[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *RecordType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIDRT[j]
	return nil
}

//=====================================================
// Record Region
type RecordRegion int

const (
	RECORDREGION_NONE RecordRegion = iota
	RECORDREGION_DEFAULT
	RECORDREGION_EUROPE
	RECORDREGION_USEAST
	RECORDREGION_USWEST
	RECORDREGION_ASIAPACIFIC
	RECORDREGION_OCEANIA
	RECORDREGION_SOUTHAMERICA
)

func (s RecordRegion) String() string {
	return toStringRR[s]
}

var toStringRR = map[RecordRegion]string{
	RECORDREGION_NONE:			"",
	RECORDREGION_DEFAULT:			"default",
	RECORDREGION_EUROPE:			"europe",
	RECORDREGION_USEAST:			"us-east",
	RECORDREGION_USWEST:			"us-west",
	RECORDREGION_ASIAPACIFIC:		"asia-pacific",
	RECORDREGION_OCEANIA:			"oceania",
	RECORDREGION_SOUTHAMERICA:	"south-america",
}

var toIDRR = map[string]RecordRegion{
	"":					RECORDREGION_NONE,					
	"default":			RECORDREGION_DEFAULT,
	"europe":			RECORDREGION_EUROPE,
	"us-east":			RECORDREGION_USEAST,
	"us-west":			RECORDREGION_USWEST,
	"asia-pacific":		RECORDREGION_ASIAPACIFIC,
	"oceania":			RECORDREGION_OCEANIA,
	"south-america":	RECORDREGION_SOUTHAMERICA,
}

// MarshalJSON marshals the enum as a quoted json string
func (s RecordRegion) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringRR[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *RecordRegion) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIDRR[j]
	return nil
}

//=====================================================
// Record Mode
type RecordMode int

const (
	RECORDMODE_NONE RecordMode = iota
	RECORDMODE_STANDARD
	RECORDMODE_FAILOVER
	RECORDMODE_ROUNDROBINFAILOVER
	RECORDMODE_POOLS
)

func (s RecordMode) String() string {
	return toStringRM[s]
}

var toStringRM = map[RecordMode]string{
	RECORDMODE_NONE:				"",
	RECORDMODE_STANDARD:			"standard",
	RECORDMODE_FAILOVER:			"failover",
	RECORDMODE_ROUNDROBINFAILOVER:	"roundrobin-failover",
	RECORDMODE_POOLS:				"pools",
}

var toIDRM = map[string]RecordMode{
	"":						RECORDMODE_NONE,					
	"standard":				RECORDMODE_STANDARD,
	"failover":				RECORDMODE_FAILOVER,
	"roundrobin-failover":	RECORDMODE_ROUNDROBINFAILOVER,
	"pools":				RECORDMODE_POOLS,
}

// MarshalJSON marshals the enum as a quoted json string
func (s RecordMode) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringRM[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *RecordMode) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIDRM[j]
	return nil
}

//=====================================================
// Record Value Mode
type RecordValueMode int

const (
	RECORDVALUEMODE_NONE RecordValueMode = iota
	RECORDVALUEMODE_NORMAL
	RECORDVALUEMODE_OFF
	RECORDVALUEMODE_ONE_WAY
)

func (s RecordValueMode) String() string {
	return toStringRVM[s]
}

var toStringRVM = map[RecordValueMode]string{
	RECORDVALUEMODE_NONE:			"",
	RECORDVALUEMODE_NORMAL:			"normal",
	RECORDVALUEMODE_OFF:			"off",
	RECORDVALUEMODE_ONE_WAY:		"one-way",
}

var toIDRVM = map[string]RecordValueMode{
	"": 			RECORDVALUEMODE_NONE,
	"normal": 		RECORDVALUEMODE_NORMAL,
	"off": 			RECORDVALUEMODE_OFF,
	"one-way": 		RECORDVALUEMODE_ONE_WAY,
}

// MarshalJSON marshals the enum as a quoted json string
func (s RecordValueMode) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringRVM[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *RecordValueMode) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIDRVM[j]
	return nil
}

//=====================================================
// Pool Type
type PoolType int

const (
	POOLTYPE_NONE PoolType = iota
	POOLTYPE_A
	POOLTYPE_AAAA
	POOLTYPE_CNAME
)

func (s PoolType) String() string {
	return toStringPT[s]
}

var toStringPT = map[PoolType]string{
	POOLTYPE_NONE:		"NONE",
	POOLTYPE_AAAA:		"AAAA",
	POOLTYPE_A:			"A", 
	POOLTYPE_CNAME:		"CNAME",
}

var toIDPT = map[string]PoolType{
	"NONE":		POOLTYPE_NONE,
	"A":		POOLTYPE_A, 
	"AAAA":		POOLTYPE_AAAA,
	"CNAME":	POOLTYPE_CNAME,
}

// MarshalJSON marshals the enum as a quoted json string
func (s PoolType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringPT[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *PoolType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIDPT[j]
	return nil
}

//=====================================================
// Pool Peroiod
type PoolPeriod int

const (
	PERIOD_NONE PoolPeriod = iota
	PERIOD_30 		= 30
	PERIOD_60 		= 60
	PERIOD_120 		= 120
	PERIOD_180 		= 180
	PERIOD_240 		= 240
	PERIOD_300 		= 300
)

//=====================================================
// Pool Deviation Allowance
type PoolDeviationAllowance int

const (
	DEVIATIONALLOWANCE_NONE PoolDeviationAllowance = iota
	DEVIATIONALLOWANCE_10 		= 10
	DEVIATIONALLOWANCE_20 		= 20
	DEVIATIONALLOWANCE_30 		= 30
	DEVIATIONALLOWANCE_40 		= 40
	DEVIATIONALLOWANCE_50 		= 50
	DEVIATIONALLOWANCE_60 		= 60
	DEVIATIONALLOWANCE_70 		= 70
	DEVIATIONALLOWANCE_80 		= 80
	DEVIATIONALLOWANCE_90 		= 90
)

//=====================================================
// Pool Monitoring Region
type PoolMonitoringRegion int

const (
	MONITORINGREGION_NONE PoolMonitoringRegion = iota
	MONITORINGREGION_WORLD
	MONITORINGREGION_ASIAPAC
	MONITORINGREGION_EUROPE
	MONITORINGREGION_NACENTRAL
	MONITORINGREGION_NAEAST
	MONITORINGREGION_NAWEST
	MONITORINGREGION_OCEANIA
	MONITORINGREGION_SOUTHAMERICA
)

func (s PoolMonitoringRegion) String() string {
	return toStringPMR[s]
}

var toStringPMR = map[PoolMonitoringRegion]string{
	MONITORINGREGION_NONE:				"",
	MONITORINGREGION_WORLD:				"world",
	MONITORINGREGION_ASIAPAC:			"asiapac",
	MONITORINGREGION_EUROPE:				"europe",
	MONITORINGREGION_NACENTRAL:			"nacentral",
	MONITORINGREGION_NAEAST:				"naeast",
	MONITORINGREGION_NAWEST:				"nawest",
	MONITORINGREGION_OCEANIA:			"oceania",
	MONITORINGREGION_SOUTHAMERICA:		"southamerica",
}

var toIDPMR = map[string]PoolMonitoringRegion{
	"": 				MONITORINGREGION_NONE,
	"world": 			MONITORINGREGION_WORLD,
	"asiapac": 			MONITORINGREGION_ASIAPAC,
	"europe": 			MONITORINGREGION_EUROPE,
	"nacentral": 		MONITORINGREGION_NACENTRAL,
	"naeast": 			MONITORINGREGION_NAEAST,
	"nawest": 			MONITORINGREGION_NAWEST,
	"oceania": 			MONITORINGREGION_OCEANIA,
	"southamerica": 	MONITORINGREGION_SOUTHAMERICA,
}

// MarshalJSON marshals the enum as a quoted json string
func (s PoolMonitoringRegion) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringPMR[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *PoolMonitoringRegion) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIDPMR[j]
	return nil
}

//=====================================================
// Pool Handicap Factor
type PoolHandicapFactor int

const (
	HANDICAPFACTOR_NONE PoolHandicapFactor = iota
	HANDICAPFACTOR_PERCENT
	HANDICAPFACTOR_SPEED
)

func (s PoolHandicapFactor) String() string {
	return toStringPHF[s]
}

var toStringPHF = map[PoolHandicapFactor]string{
	HANDICAPFACTOR_NONE:			"none",
	HANDICAPFACTOR_PERCENT:		"percent",
	HANDICAPFACTOR_SPEED:			"speed",
}

var toIDPHF = map[string]PoolHandicapFactor{
	"none": 			HANDICAPFACTOR_NONE,
	"percent": 			HANDICAPFACTOR_PERCENT,
	"speed": 			HANDICAPFACTOR_SPEED,
}

// MarshalJSON marshals the enum as a quoted json string
func (s PoolHandicapFactor) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringPHF[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *PoolHandicapFactor) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIDPHF[j]
	return nil
}

//=====================================================
// Continent
type Continent int

const (
	CONTINENT_NONE Continent = iota
	CONTINENT_WORLD
	CONTINENT_AF
	CONTINENT_AN
	CONTINENT_AS
	CONTINENT_EU
	CONTINENT_NA
	CONTINENT_OC
	CONTINENT_SA
)

func (s Continent) String() string {
	return toStringC[s]
}

var toStringC = map[Continent]string{
	CONTINENT_NONE:		"",
	CONTINENT_WORLD:	"world",
	CONTINENT_AF:		"AF",
	CONTINENT_AN:		"AN",
	CONTINENT_AS:		"AS",
	CONTINENT_EU:		"EU",
	CONTINENT_NA:		"NA",
	CONTINENT_OC:		"OC",
	CONTINENT_SA:		"SA",
}

var toIDC = map[string]Continent{
	"":			CONTINENT_NONE,
	"world":	CONTINENT_WORLD, 
	"AF":		CONTINENT_AF,
	"AN":		CONTINENT_AN,
	"AS":		CONTINENT_AS,
	"EU":		CONTINENT_EU,
	"NA":		CONTINENT_NA,
	"OC":		CONTINENT_OC,
	"SA":		CONTINENT_SA,
}

// MarshalJSON marshals the enum as a quoted json string
func (s Continent) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringC[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *Continent) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIDC[j]
	return nil
}

//=====================================================
// Pool Value Policy
type PoolValuePolicy int

const (
	POOLVALUEPOLICY_NONE PoolValuePolicy = iota
	POOLVALUEPOLICY_FOLLOW_SONAR
	POOLVALUEPOLICY_ALWAYS_OFF
	POOLVALUEPOLICY_ALWAYS_ON
	POOLVALUEPOLICY_OFF_ON_FAILURE
)

func (s PoolValuePolicy) String() string {
	return toStringPVP[s]
}

var toStringPVP = map[PoolValuePolicy]string{
	POOLVALUEPOLICY_NONE:				"",
	POOLVALUEPOLICY_FOLLOW_SONAR:		"follow_sonar",
	POOLVALUEPOLICY_ALWAYS_OFF:			"always_off",
	POOLVALUEPOLICY_ALWAYS_ON:			"always_on",
	POOLVALUEPOLICY_OFF_ON_FAILURE:		"off_on_failure",
}

var toIDPVP = map[string]PoolValuePolicy{
	"":					POOLVALUEPOLICY_NONE,
	"follow_sonar":		POOLVALUEPOLICY_FOLLOW_SONAR, 
	"always_off":		POOLVALUEPOLICY_ALWAYS_OFF,
	"always_on":		POOLVALUEPOLICY_ALWAYS_ON,
	"off_on_failure":	POOLVALUEPOLICY_OFF_ON_FAILURE,
}

// MarshalJSON marshals the enum as a quoted json string
func (s PoolValuePolicy) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringPVP[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *PoolValuePolicy) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIDPVP[j]
	return nil
}
