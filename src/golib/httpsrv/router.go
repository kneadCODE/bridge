// Package httpsrv contains HTTP Server
package httpsrv

import (
	"fmt"
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi/v5"
)

// Router defines the routes for the server
type Router struct {
	ReadinessHandlerFunc http.HandlerFunc
	GQLHandler           http.Handler
	CustomRESTRoutes     func(chi.Router)
	ProfilingEnabled     bool
}

// Handler returns the http.Handler for the HTTP server
func (rt Router) Handler() chi.Router {
	rtr := chi.NewRouter()

	rtr.Get("/_/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		_, _ = fmt.Fprintln(w, "ok") // Intentionally ignoring the err as nothing to do once caught.
	})

	if rt.ReadinessHandlerFunc != nil {
		rtr.Get("/_/ready", rt.ReadinessHandlerFunc)
	}

	if rt.ProfilingEnabled {
		// Based on https: //pkg.go.dev/net/http/pprof
		rtr.HandleFunc("/_/profile/*", pprof.Index)
		rtr.HandleFunc("/_/profile/cmdline", pprof.Cmdline)
		rtr.HandleFunc("/_/profile/profile", pprof.Profile)
		rtr.HandleFunc("/_/profile/symbol", pprof.Symbol)
		rtr.HandleFunc("/_/profile/trace", pprof.Trace)
		rtr.Handle("/_/profile/goroutine", pprof.Handler("goroutine"))
		rtr.Handle("/_/profile/threadcreate", pprof.Handler("threadcreate"))
		rtr.Handle("/_/profile/mutex", pprof.Handler("mutex"))
		rtr.Handle("/_/profile/heap", pprof.Handler("heap"))
		rtr.Handle("/_/profile/block", pprof.Handler("block"))
		rtr.Handle("/_/profile/allocs", pprof.Handler("allocs"))
	}

	rtr.Group(func(r chi.Router) {
		// TODO: Add Middlewares

		if rt.GQLHandler != nil {
			rtr.Handle("/graph", rt.GQLHandler)
		}

		if rt.CustomRESTRoutes != nil {
			rtr.Group(rt.CustomRESTRoutes)
		}
	})

	return rtr
}
