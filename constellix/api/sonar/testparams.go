package sonar

type HTTPTestParam struct {
	Host 				string			`json:"host,omitempty"`
    Port 				int 			`json:"port,omitempty"`
    ProtocolType 		ProtocolType	`json:"protocolType,omitempty"`
    Path 				string 			`json:"path,omitempty"`
    Fqdn 				string 			`json:"fqdn,omitempty"`
    StringToSearch 		string 			`json:"stringToSearch,omitempty"`
    ExpectedStatusCode 	int 			`json:"expectedStatusCode,omitempty"`
    IpVersion			IPVersion 		`json:"ipVersion,omitempty"`
}

type TraceTestParam struct {
	Host 				string			`json:"host,omitempty"`
    ProtocolType 		ProtocolType	`json:"type,omitempty"`
    IpVersion			IPVersion 		`json:"ipVersion,omitempty"`
}

type TCPTestParam struct {
	Host 				string			`json:"host,omitempty"`
    Port 				int 			`json:"port,omitempty"`
	StringToSearch 		string 			`json:"stringToSearch,omitempty"`
	StringToSend 		string 			`json:"stringToSend,omitempty"`
	StringToReceive		string 			`json:"stringToReceive,omitempty"`
    IpVersion			IPVersion 		`json:"ipVersion,omitempty"`
}

type DNSTestParam struct {
	Host 				string			`json:"host,omitempty"`
    NameServer 			string			`json:"nameServer,omitempty"`
    RecordType 			RecordType		`json:"recordType,omitempty"`
    ExpectedIp 			string 			`json:"expectedIp,omitempty"`
}
