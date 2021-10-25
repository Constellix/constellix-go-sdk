module constellix.com/constellix/examples

require (
	constellix.com/constellix/api v0.0.0
	constellix.com/constellix/api/dns v0.0.0
	constellix.com/constellix/api/sonar v0.0.0
	github.com/tidwall/gjson v1.8.0
)

replace (
	constellix.com/constellix/api => ../constellix/api
	constellix.com/constellix/api/dns => ../constellix/api/dns
	constellix.com/constellix/api/sonar => ../constellix/api/sonar
)
