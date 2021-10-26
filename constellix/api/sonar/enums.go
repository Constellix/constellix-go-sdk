package sonar

import (
	"bytes"
	"encoding/json"
)

//=====================================================
// Protocol Type
type ProtocolType int

const (
	PROTOCOLTYPE_NONE ProtocolType = iota
	PROTOCOLTYPE_HTTP
	PROTOCOLTYPE_HTTPS
	PROTOCOLTYPE_TCP
	PROTOCOLTYPE_UDP
	PROTOCOLTYPE_DNS
	PROTOCOLTYPE_RUM
	PROTOCOLTYPE_ICMP
)

func (s ProtocolType) String() string {
	return toStringPT[s]
}

var toStringPT = map[ProtocolType]string{
	PROTOCOLTYPE_NONE:		"",
	PROTOCOLTYPE_HTTP:		"HTTP",
	PROTOCOLTYPE_HTTPS:		"HTTPS",
	PROTOCOLTYPE_TCP:		"TCP",
	PROTOCOLTYPE_UDP:		"UDP",
	PROTOCOLTYPE_DNS:		"DNS",
	PROTOCOLTYPE_RUM:		"RUM",
	PROTOCOLTYPE_ICMP:		"ICMP",
}

var toIDPT = map[string]ProtocolType{
	"":				PROTOCOLTYPE_NONE,
	"HTTP":			PROTOCOLTYPE_HTTP,
	"HTTPS":		PROTOCOLTYPE_HTTPS,
	"TCP":			PROTOCOLTYPE_TCP,
	"UDP":			PROTOCOLTYPE_UDP,
	"DNS":			PROTOCOLTYPE_DNS,
	"RUM":			PROTOCOLTYPE_RUM,
	"ICMP":			PROTOCOLTYPE_ICMP,
}

// MarshalJSON marshals the enum as a quoted json string
func (s ProtocolType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringPT[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *ProtocolType) UnmarshalJSON(b []byte) error {
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
// IP Version
type IPVersion int

const (
	IPVERSION_NONE IPVersion = iota
	IPVERSION_IPV4
	IPVERSION_IPV6
)

func (s IPVersion) String() string {
	return toStringIPV[s]
}

var toStringIPV = map[IPVersion]string{
	IPVERSION_NONE:			"",
	IPVERSION_IPV4:		"IPV4",
	IPVERSION_IPV6:		"IPV6",
}

var toIDIPV = map[string]IPVersion{
	"":				IPVERSION_NONE,
	"IPV4":			IPVERSION_IPV4,
	"IPV6":			IPVERSION_IPV6,
}

// MarshalJSON marshals the enum as a quoted json string
func (s IPVersion) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringIPV[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *IPVersion) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIDIPV[j]
	return nil
}

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
// Schedule Interval
type ScheduleInterval int

const (
	SCHEDULEINTERVAL_NONE ScheduleInterval = iota 
	SCHEDULEINTERVAL_DAILY
	SCHEDULEINTERVAL_WEEKLY
	SCHEDULEINTERVAL_MONTHLY
)

func (s ScheduleInterval) String() string {
	return toStringSI[s]
}

var toStringSI = map[ScheduleInterval]string{
	SCHEDULEINTERVAL_NONE:			"NONE",
	SCHEDULEINTERVAL_DAILY:			"DAILY",
	SCHEDULEINTERVAL_WEEKLY:		"WEEKLY",
	SCHEDULEINTERVAL_MONTHLY: 		"MONTHLY",
}

var toIDSI = map[string]ScheduleInterval{
	"NONE":				SCHEDULEINTERVAL_NONE,
	"DAILY":			SCHEDULEINTERVAL_DAILY,
	"WEEKLY":			SCHEDULEINTERVAL_WEEKLY,
	"MONTHLY":			SCHEDULEINTERVAL_MONTHLY,
}

// MarshalJSON marshals the enum as a quoted json string
func (s ScheduleInterval) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringSI[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *ScheduleInterval) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIDSI[j]
	return nil
}

//=====================================================
// Monitoring Interval
type MonitoringInterval int

const (
	MONITORINGINTERVAL_NONE MonitoringInterval = iota
	MONITORINGINTERVAL_FIVESECONDS
	MONITORINGINTERVAL_THIRTYSECONDS
	MONITORINGINTERVAL_ONEMINUTE
	MONITORINGINTERVAL_TWOMINUTES
	MONITORINGINTERVAL_THREEMINUTES
	MONITORINGINTERVAL_FOURMINUTES
	MONITORINGINTERVAL_FIVEMINUTES
	MONITORINGINTERVAL_TENMINUTES
	MONITORINGINTERVAL_THIRTYMINUTES
	MONITORINGINTERVAL_HALFDAY
	MONITORINGINTERVAL_DAY
)

func (s MonitoringInterval) String() string {
	return toStringMI[s]
}

var toStringMI = map[MonitoringInterval]string{
	MONITORINGINTERVAL_NONE:				"",
	MONITORINGINTERVAL_FIVESECONDS:			"FIVESECONDS",
	MONITORINGINTERVAL_THIRTYSECONDS:		"THIRTYSECONDS",
	MONITORINGINTERVAL_ONEMINUTE:			"ONEMINUTE",
	MONITORINGINTERVAL_TWOMINUTES:			"TWOMINUTES",
	MONITORINGINTERVAL_THREEMINUTES:		"THREEMINUTES",
	MONITORINGINTERVAL_FOURMINUTES:			"FOURMINUTES",
	MONITORINGINTERVAL_FIVEMINUTES:			"FIVEMINUTES",
	MONITORINGINTERVAL_TENMINUTES:			"TENMINUTES",
	MONITORINGINTERVAL_THIRTYMINUTES:		"THIRTYMINUTES",
	MONITORINGINTERVAL_HALFDAY:				"HALFDAY",
	MONITORINGINTERVAL_DAY:					"DAY",
}

var toIDMI = map[string]MonitoringInterval{
	"":					MONITORINGINTERVAL_NONE,
	"FIVESECONDS":		MONITORINGINTERVAL_FIVESECONDS,
	"THIRTYSECONDS":	MONITORINGINTERVAL_THIRTYSECONDS,
	"ONEMINUTE":		MONITORINGINTERVAL_ONEMINUTE,
	"TWOMINUTES":		MONITORINGINTERVAL_TWOMINUTES,
	"THREEMINUTES":		MONITORINGINTERVAL_THREEMINUTES,
	"FOURMINUTES":		MONITORINGINTERVAL_FOURMINUTES,
	"FIVEMINUTES":		MONITORINGINTERVAL_FIVEMINUTES,
	"TENMINUTES":		MONITORINGINTERVAL_TENMINUTES,
	"THIRTYMINUTES":	MONITORINGINTERVAL_THIRTYMINUTES,
	"HALFDAY":			MONITORINGINTERVAL_HALFDAY,
	"DAY":				MONITORINGINTERVAL_DAY,
}

// MarshalJSON marshals the enum as a quoted json string
func (s MonitoringInterval) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringMI[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *MonitoringInterval) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIDMI[j]
	return nil
}

//=====================================================
// Monitor Interval Policy
type MonitorIntervalPolicy int

const (
	MONITORINTERVALPOLICY_NONE MonitorIntervalPolicy = iota 
	MONITORINTERVALPOLICY_PARALLEL
	MONITORINTERVALPOLICY_ONCEPERSITE
	MONITORINTERVALPOLICY_ONCEPERREGION
)

func (s MonitorIntervalPolicy) String() string {
	return toStringMIP[s]
}

var toStringMIP = map[MonitorIntervalPolicy]string{
	MONITORINTERVALPOLICY_NONE:				"",
	MONITORINTERVALPOLICY_PARALLEL:			"PARALLEL",
	MONITORINTERVALPOLICY_ONCEPERSITE:		"ONCEPERSITE",
	MONITORINTERVALPOLICY_ONCEPERREGION: 	"ONCEPERREGION",
}

var toIMIP = map[string]MonitorIntervalPolicy{
	"":					MONITORINTERVALPOLICY_NONE,
	"PARALLEL":			MONITORINTERVALPOLICY_PARALLEL,
	"ONCEPERSITE":		MONITORINTERVALPOLICY_ONCEPERSITE,
	"ONCEPERREGION":	MONITORINTERVALPOLICY_ONCEPERREGION,
}

// MarshalJSON marshals the enum as a quoted json string
func (s MonitorIntervalPolicy) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringMIP[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *MonitorIntervalPolicy) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIMIP[j]
	return nil
}

//=====================================================
// Verification Policy
type VerificationPolicy int

const (
	VERIFICATIONPOLICY_NONE VerificationPolicy = iota 
	VERIFICATIONPOLICY_SIMPLE
	VERIFICATIONPOLICY_MAJORITY
)

func (s VerificationPolicy) String() string {
	return toStringVP[s]
}

var toStringVP = map[VerificationPolicy]string{
	VERIFICATIONPOLICY_NONE:			"",
	VERIFICATIONPOLICY_SIMPLE:			"SIMPLE",
	VERIFICATIONPOLICY_MAJORITY:		"MAJORITY",
}

var toIVP = map[string]VerificationPolicy{
	"":					VERIFICATIONPOLICY_NONE,
	"SIMPLE":			VERIFICATIONPOLICY_SIMPLE,
	"MAJORITY":			VERIFICATIONPOLICY_MAJORITY,
}

// MarshalJSON marshals the enum as a quoted json string
func (s VerificationPolicy) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringVP[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *VerificationPolicy) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIVP[j]
	return nil
}

//=====================================================
// DNS Query Protocol
type DNSQueryProtocol int

const (
	DNSQUERYPROTOCOL_NONE DNSQueryProtocol = iota
	DNSQUERYPROTOCOL_TCP
	DNSQUERYPROTOCOL_UDP
)

func (s DNSQueryProtocol) String() string {
	return toStringDNSQP[s]
}

var toStringDNSQP = map[DNSQueryProtocol]string{
	DNSQUERYPROTOCOL_NONE:		"",
	DNSQUERYPROTOCOL_TCP:		"TCP",
	DNSQUERYPROTOCOL_UDP:		"UDP",
}

var toIDDNSQP = map[string]DNSQueryProtocol{
	"":				DNSQUERYPROTOCOL_NONE,
	"TCP":			DNSQUERYPROTOCOL_TCP,
	"UDP":			DNSQUERYPROTOCOL_UDP,
}

// MarshalJSON marshals the enum as a quoted json string
func (s DNSQueryProtocol) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringDNSQP[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *DNSQueryProtocol) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIDDNSQP[j]
	return nil
}

//=====================================================
// DNS Compare Option
type DNSCompareOption int

const (
	DNSCOMPAREOPTION_NONE DNSCompareOption = iota
	DNSCOMPAREOPTION_EQUALS
	DNSCOMPAREOPTION_CONTAINS
	DNSCOMPAREOPTION_ONEOFF
	DNSCOMPAREOPTION_ANYMATCH
)

func (s DNSCompareOption) String() string {
	return toStringDNSCO[s]
}

var toStringDNSCO = map[DNSCompareOption]string{
	DNSCOMPAREOPTION_NONE:			"",
	DNSCOMPAREOPTION_EQUALS:		"EQUALS",
	DNSCOMPAREOPTION_CONTAINS:		"CONTAINS",
	DNSCOMPAREOPTION_ONEOFF:		"ONEOFF",
	DNSCOMPAREOPTION_ANYMATCH:		"ANYMATCH",
}

var toIDDNSCO = map[string]DNSCompareOption{
	"": 			DNSCOMPAREOPTION_NONE,
	"EQUALS": 		DNSCOMPAREOPTION_EQUALS,
	"CONTAINS": 	DNSCOMPAREOPTION_CONTAINS,
	"ONEOFF": 		DNSCOMPAREOPTION_ONEOFF,
	"ANYMATCH": 	DNSCOMPAREOPTION_ANYMATCH,
}

// MarshalJSON marshals the enum as a quoted json string
func (s DNSCompareOption) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringDNSCO[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *DNSCompareOption) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toIDDNSCO[j]
	return nil
}

//=====================================================
// run Traceroute
type RunTraceroute int

const (
	RUNTRACEROUTE_NONE RunTraceroute = iota
	RUNTRACEROUTE_DISABLED
	RUNTRACEROUTE_ON_STATUS_CHANGE
	RUNTRACEROUTE_WITH_CHECK
)

func (s RunTraceroute) String() string {
	return toStringRunTR[s]
}

var toStringRunTR = map[RunTraceroute]string{
	RUNTRACEROUTE_NONE:				"",
	RUNTRACEROUTE_DISABLED:			"DISABLED",
	RUNTRACEROUTE_ON_STATUS_CHANGE:	"ON_STATUS_CHANGE",
	RUNTRACEROUTE_WITH_CHECK:		"WITH_CHECK",
}

var toRunTR = map[string]RunTraceroute{
	"":					RUNTRACEROUTE_NONE,
	"DISABLED":			RUNTRACEROUTE_DISABLED,
	"ON_STATUS_CHANGE":	RUNTRACEROUTE_ON_STATUS_CHANGE,
	"WITH_CHECK":		RUNTRACEROUTE_WITH_CHECK,
}

// MarshalJSON marshals the enum as a quoted json string
func (s RunTraceroute) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringRunTR[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *RunTraceroute) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value.
	*s = toRunTR[j]
	return nil
}
