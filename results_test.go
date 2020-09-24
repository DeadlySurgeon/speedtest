package speedtest

var (
	validTestResult = Results{

		Type:       "result",
		PacketLoss: 0,
		Timestamp:  "2020-09-19T15:02:30Z",
		Download: Link{
			Bandwidth: 32277398,
			Bytes:     386257144,
			Elapsed:   12403,
		},
		Upload: Link{
			Bandwidth: 1463943,
			Bytes:     9942424,
			Elapsed:   6814,
		},
		Ping: Ping{
			Jitter:  0.925,
			Latency: 7.411,
		},
		Server: Server{
			ID:       2391,
			Name:     "Spectrum",
			Location: "Olivette, MO",
			Country:  "United States",
			Host:     "spt01olvemo.stls.mo.charter.com",
			Port:     8080,
			IP:       "24.217.2.210",
		},
		ISP: "Spectrum",
		Interface: Interface{
			ExternalIP: "0.0.0.0",
			InternalIP: "192.168.1.69",
			IsVpn:      false,
			MacAddr:    "123 Big Mac Dr, Atlanta, Georgia 30318",
			Name:       "wlan0",
		},
		TestLink: TestLink{
			ID:  "123",
			URL: "speedtest.net/123",
		},
	}

	invalidTypeResult = Results{
		Type: "Not it",
	}
)
