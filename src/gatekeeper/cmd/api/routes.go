// Binary API
package main

import (
	systemh "github.com/kneadCODE/bridge/src/gatekeeper/internal/handler/system"
	"github.com/kneadCODE/bridge/src/golib/httpsrv"
)

func router() httpsrv.Router {
	return httpsrv.Router{
		ReadinessHandlerFunc: systemh.Readiness,
		ProfilingEnabled:     true,
	}
}
