package elasticsearch

func FakeAPIIPDataResponse() []byte {
	return []byte(`{
		"ip": "189.216.190.197",
		"is_eu": false,
		"city": "Mexico City",
		"region": "Mexico City",
		"region_code": "CMX",
		"country_name": "Mexico",
		"country_code": "MX",
		"continent_name": "North America",
		"continent_code": "NA",
		"latitude": 19.4471,
		"longitude": -99.1599,
		"asn": "AS28548",
		"organisation": "Cablevisi\u00f3n, S.A. de C.V.",
		"postal": "06400",
		"calling_code": "52",
		"flag": "https://ipdata.co/flags/mx.png",
		"emoji_flag": "\ud83c\uddf2\ud83c\uddfd",
		"emoji_unicode": "U+1F1F2 U+1F1FD",
		"languages": [
			{
				"name": "Spanish",
				"native": "Espa\u00f1ol"
			}
		],
		"currency": {
			"name": "Mexican Peso",
			"code": "MXN",
			"symbol": "MX$",
			"native": "$",
			"plural": "Mexican pesos"
		},
		"time_zone": {
			"name": "America/Mexico_City",
			"abbr": "CST",
			"offset": "-0600",
			"is_dst": false,
			"current_time": "2018-12-19T10:50:37.909404-06:00"
		},
		"threat": {
			"is_tor": false,
			"is_proxy": false,
			"is_anonymous": false,
			"is_known_attacker": false,
			"is_known_abuser": false,
			"is_threat": false,
			"is_bogon": false
		},
		"count": "1511"
	}`)
}