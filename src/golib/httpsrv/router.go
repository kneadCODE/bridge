// Package httpsrv contains HTTP Server
package httpsrv

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"runtime/debug"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"
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
		r.Use(baseMiddleware())

		if rt.GQLHandler != nil {
			r.Handle("/graph", rt.GQLHandler)
		}

		if rt.CustomRESTRoutes != nil {
			r.Group(rt.CustomRESTRoutes)
		}
	})

	return rtr
}

func baseMiddleware() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			logger := slog.FromContext(r.Context())

			defer func() {
				if rcv := recover(); rcv != nil {
					if err, ok := rcv.(error); ok {
						logger.Error(fmt.Sprintf("PANIC RECOVERED. Stack: [%s]", string(debug.Stack())), err)
						return
					}
					logger.LogAttrs(
						slog.ErrorLevel,
						fmt.Sprintf("PANIC RECOVERED: [%+v]. Stack: [%s]", rcv, string(debug.Stack())),
					)
				}
			}()

			logger = logger.With(
				slog.String("http.req.method", r.Method),
				slog.String("http.req.path", r.URL.Path),
				slog.String("http.req.host", r.URL.Host),
				slog.String("http.req.user-agent", r.UserAgent()),
				slog.String("http.req.referer", r.Referer()),
				slog.String("http.req.remote_addr", r.RemoteAddr),
			)

			rw := &respWriter{ResponseWriter: w}

			logger.LogAttrs(slog.InfoLevel,
				"START HTTP Request",
				slog.Int64("http.req.content-length", r.ContentLength),
				slog.String("http.req.content-type", r.Header.Get("Content-Type")),
				slog.String("http.req.proto", r.Proto),
				slog.Time("http.req.start", start),
			)

			processStart := time.Now()
			next.ServeHTTP(w, r)

			logger.LogAttrs(slog.InfoLevel,
				"END HTTP Request",
				slog.Time("http.resp.end", time.Now()),
				slog.String("http.resp.duration", fmt.Sprintf("%dms", time.Since(start).Milliseconds())),
				slog.String("http.resp.processing_duration", fmt.Sprintf("%dms", time.Since(processStart).Milliseconds())),
				slog.Int("http.resp.bytes", rw.contentLength),
				slog.Int("http.resp.status_code", rw.statusCode),
			)
		})
	}
}

type respWriter struct {
	http.ResponseWriter

	statusCode    int
	contentLength int
}

func (w *respWriter) Write(b []byte) (int, error) {
	length, err := w.ResponseWriter.Write(b)
	w.contentLength = length
	return length, err
}

func (w *respWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
