package speedtest

// Results contains the result of the test. The `Type` field is only used for
// testing, should always be "result".
type Results struct {
	found      bool
	Type       string    `json:"type"` // Used explicitly for testing.
	Timestamp  string    `json:"timestamp"`
	Ping       Ping      `json:"ping"`
	Download   Link      `json:"download"`
	Upload     Link      `json:"upload"`
	PacketLoss int       `json:"packetLoss"`
	ISP        string    `json:"isp"`
	Interface  Interface `json:"interface"`
	Server     Server    `json:"server"`
	TestLink   TestLink  `json:"result"`
}

// Link contains the results of an up or down link.
type Link struct {
	Bandwidth int `json:"bandwidth"`
	Bytes     int `json:"bytes"`
	Elapsed   int `json:"elapsed"`
}

// Ping contains the results of a ping.
type Ping struct {
	Jitter  float64 `json:"jitter"`
	Latency float64 `json:"latency"`
}

// Interface contains the information of the network interface used in the
// test.
type Interface struct {
	InternalIP string `json:"internalIp"`
	Name       string `json:"name"`
	MacAddr    string `json:"macAddr"`
	IsVpn      bool   `json:"isVpn"`
	ExternalIP string `json:"externalIp"`
}

// Server contains information regarding the server that we contacted.
type Server struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Country  string `json:"country"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	IP       string `json:"ip"`
}

// TestLink contains the speedtest.net link resources to view the results
// on their site.
type TestLink struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}
