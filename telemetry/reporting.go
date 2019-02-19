package monitoring

import (
	telemetry "github.com/creativesoftwarefdn/weaviate/telemetry/utils"
)

type Reporter struct {
	if telemetry.IsEnabled() {
		// output every X (300) seconds
		// --CBOR conversion -- handled by ugorji
		// --Failsafe -- TBD
		// --get timestamp and apply to output nested objects
	}
}