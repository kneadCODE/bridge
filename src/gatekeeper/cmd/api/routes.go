// Binary API
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	systemh "github.com/kneadCODE/bridge/src/gatekeeper/internal/handler/system"
	"github.com/kneadCODE/bridge/src/golib/httpsrv"
)

func router() httpsrv.Router {
	return httpsrv.Router{
		ReadinessHandlerFunc: systemh.Readiness,
		ProfilingEnabled:     true,
		CustomRESTRoutes: func(r chi.Router) {
			r.Get("/something", func(w http.ResponseWriter, r *http.Request) {

			})
		},
	}
}
