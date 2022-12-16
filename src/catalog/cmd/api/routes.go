// Binary API
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kneadCODE/bridge/src/catalog/internal/handler/gql"
	systemh "github.com/kneadCODE/bridge/src/catalog/internal/handler/system"
	"github.com/kneadCODE/bridge/src/golib/httpsrv"
)

func router() httpsrv.Router {
	return httpsrv.Router{
		ReadinessHandlerFunc: systemh.Readiness,
		ProfilingEnabled:     true,
		GQLHandler:           gql.NewHandler(true), // TODO: Get introspection config from envvar
		CustomRESTRoutes: func(r chi.Router) {
			r.Get("/something", func(w http.ResponseWriter, r *http.Request) {

			})
		},
	}
}
