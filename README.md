# Golang SDK for Constellix API v4

# Usage Examples

Setting apikey and secretkey over environment variables

```
export CONSTELLIX_API_KEY=<api_key>
export CONSTELLIX_SECRET_KEY=<secret_key>
```

```go
import (
    "constellix.com/constellix/api/dns"
    "constellix.com/constellix/api/sonar"
)

func main() {
    constellixDns := dns.Init("", "")
    constellixSonar := sonar.Init("", "")
}
```

Passing apikey and secretkey by values

```go
import (
    "constellix.com/constellix/api/dns"
    "constellix.com/constellix/api/sonar"
)

func main() {
    constellixDns := dns.Init("api_key", "secret_key")
    constellixSonar := sonar.Init("api_key", "secret_key")
}
```

Please check examples folder for sample usages