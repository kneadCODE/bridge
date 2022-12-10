package httpsrv

import (
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestRouter_Handler(t *testing.T) {
	type testCase struct {
		givenRouter Router
		expRoutes   []string
	}
	tcs := map[string]testCase{
		"default": {
			expRoutes: []string{
				"GET /_/ping",
			},
		},
		"with readiness": {
			givenRouter: Router{
				ReadinessHandlerFunc: func(http.ResponseWriter, *http.Request) {},
			},
			expRoutes: []string{
				"GET /_/ping",
				"GET /_/ready",
			},
		},
		"with readiness & gql": {
			givenRouter: Router{
				ReadinessHandlerFunc: func(http.ResponseWriter, *http.Request) {},
				GQLHandler:           http.NewServeMux(),
			},
			expRoutes: []string{
				"GET /_/ping",
				"GET /_/ready",
				"GET /graph",
				"POST /graph",
				"PUT /graph",
				"PATCH /graph",
				"DELETE /graph",
				"CONNECT /graph",
				"OPTIONS /graph",
				"TRACE /graph",
				"HEAD /graph",
			},
		},
		"with readiness, gql & custom rest": {
			givenRouter: Router{
				ReadinessHandlerFunc: func(http.ResponseWriter, *http.Request) {},
				GQLHandler:           http.NewServeMux(),
				CustomRESTRoutes: func(r chi.Router) {
					r.Post("/post", func(http.ResponseWriter, *http.Request) {})
				},
			},
			expRoutes: []string{
				"GET /_/ping",
				"GET /_/ready",
				"GET /graph",
				"POST /graph",
				"PUT /graph",
				"PATCH /graph",
				"DELETE /graph",
				"CONNECT /graph",
				"OPTIONS /graph",
				"TRACE /graph",
				"HEAD /graph",
				"POST /post",
			},
		},
		"with readiness, gql, custom rest & profiling": {
			givenRouter: Router{
				ReadinessHandlerFunc: func(http.ResponseWriter, *http.Request) {},
				GQLHandler:           http.NewServeMux(),
				CustomRESTRoutes: func(r chi.Router) {
					r.Post("/post", func(http.ResponseWriter, *http.Request) {})
				},
				ProfilingEnabled: true,
			},
			expRoutes: []string{
				"GET /_/ping",
				"GET /_/ready",
				"CONNECT /_/profile/*", "CONNECT /_/profile/allocs", "CONNECT /_/profile/block", "CONNECT /_/profile/cmdline", "CONNECT /_/profile/goroutine", "CONNECT /_/profile/heap", "CONNECT /_/profile/mutex", "CONNECT /_/profile/profile", "CONNECT /_/profile/symbol", "CONNECT /_/profile/threadcreate", "CONNECT /_/profile/trace",
				"DELETE /_/profile/*", "DELETE /_/profile/allocs", "DELETE /_/profile/block", "DELETE /_/profile/cmdline", "DELETE /_/profile/goroutine", "DELETE /_/profile/heap", "DELETE /_/profile/mutex", "DELETE /_/profile/profile", "DELETE /_/profile/symbol", "DELETE /_/profile/threadcreate", "DELETE /_/profile/trace",
				"GET /_/profile/*", "GET /_/profile/allocs", "GET /_/profile/block", "GET /_/profile/cmdline", "GET /_/profile/goroutine", "GET /_/profile/heap", "GET /_/profile/mutex", "GET /_/profile/profile", "GET /_/profile/symbol", "GET /_/profile/threadcreate", "GET /_/profile/trace",
				"HEAD /_/profile/*", "HEAD /_/profile/allocs", "HEAD /_/profile/block", "HEAD /_/profile/cmdline", "HEAD /_/profile/goroutine", "HEAD /_/profile/heap", "HEAD /_/profile/mutex", "HEAD /_/profile/profile", "HEAD /_/profile/symbol", "HEAD /_/profile/threadcreate", "HEAD /_/profile/trace",
				"OPTIONS /_/profile/*", "OPTIONS /_/profile/allocs", "OPTIONS /_/profile/block", "OPTIONS /_/profile/cmdline", "OPTIONS /_/profile/goroutine", "OPTIONS /_/profile/heap", "OPTIONS /_/profile/mutex", "OPTIONS /_/profile/profile", "OPTIONS /_/profile/symbol", "OPTIONS /_/profile/threadcreate", "OPTIONS /_/profile/trace",
				"PATCH /_/profile/*", "PATCH /_/profile/allocs", "PATCH /_/profile/block", "PATCH /_/profile/cmdline", "PATCH /_/profile/goroutine", "PATCH /_/profile/heap", "PATCH /_/profile/mutex", "PATCH /_/profile/profile", "PATCH /_/profile/symbol", "PATCH /_/profile/threadcreate", "PATCH /_/profile/trace",
				"POST /_/profile/*", "POST /_/profile/allocs", "POST /_/profile/block", "POST /_/profile/cmdline", "POST /_/profile/goroutine", "POST /_/profile/heap", "POST /_/profile/mutex", "POST /_/profile/profile", "POST /_/profile/symbol", "POST /_/profile/threadcreate", "POST /_/profile/trace",
				"PUT /_/profile/*", "PUT /_/profile/allocs", "PUT /_/profile/block", "PUT /_/profile/cmdline", "PUT /_/profile/goroutine", "PUT /_/profile/heap", "PUT /_/profile/mutex", "PUT /_/profile/profile", "PUT /_/profile/symbol", "PUT /_/profile/threadcreate", "PUT /_/profile/trace",
				"TRACE /_/profile/*", "TRACE /_/profile/allocs", "TRACE /_/profile/block", "TRACE /_/profile/cmdline", "TRACE /_/profile/goroutine", "TRACE /_/profile/heap", "TRACE /_/profile/mutex", "TRACE /_/profile/profile", "TRACE /_/profile/symbol", "TRACE /_/profile/threadcreate", "TRACE /_/profile/trace",
				"GET /graph",
				"POST /graph",
				"PUT /graph",
				"PATCH /graph",
				"DELETE /graph",
				"CONNECT /graph",
				"OPTIONS /graph",
				"TRACE /graph",
				"HEAD /graph",
				"POST /post",
			},
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given && When:
			var routesFound []string
			err := chi.Walk(
				tc.givenRouter.Handler(),
				func(method string, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
					routesFound = append(routesFound, method+" "+route)
					return nil
				},
			)

			// Then:
			require.NoError(t, err)

			sort.Strings(routesFound)
			sort.Strings(tc.expRoutes)

			require.Equal(t, tc.expRoutes, routesFound)
		})
	}
}

func TestBaseMiddleware(t *testing.T) {
	// Given:
	rtr := Router{
		CustomRESTRoutes: func(r chi.Router) {
			r.Get("/get", func(http.ResponseWriter, *http.Request) {})
		},
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/get", nil)

	// When && Then: // Don't crash
	rtr.Handler().ServeHTTP(w, r)

	// Right now unable to test the slog.fields. TODO: Look into how to verify what log fields were recorded.
}

func TestResponseWriter(t *testing.T) {
	// Given:
	rw := respWriter{ResponseWriter: httptest.NewRecorder()}
	// When:
	rw.WriteHeader(http.StatusOK)
	// Then:
	require.Equal(t, http.StatusOK, rw.statusCode)
	require.Equal(t, 0, rw.contentLength)

	// Given:
	rw = respWriter{ResponseWriter: httptest.NewRecorder()}
	// When:
	rw.WriteHeader(http.StatusInternalServerError)
	// Then:
	require.Equal(t, http.StatusInternalServerError, rw.statusCode)
	require.Equal(t, 0, rw.contentLength)

	// Given && When:
	rw = respWriter{ResponseWriter: httptest.NewRecorder()}
	// Then:
	require.Equal(t, 0, rw.statusCode)
	require.Equal(t, 0, rw.contentLength)

	// Given:
	rw = respWriter{ResponseWriter: httptest.NewRecorder()}
	// When:
	l, err := rw.Write([]byte("hello"))
	// Then:
	require.Equal(t, 0, rw.statusCode)
	require.Equal(t, l, rw.contentLength)
	require.Equal(t, len("hello"), rw.contentLength)
	require.NoError(t, err)

	// Given:
	rw = respWriter{ResponseWriter: httptest.NewRecorder()}
	// When:
	l, err = rw.Write([]byte("hello hello hello hello hello"))
	// Then:
	require.Equal(t, 0, rw.statusCode)
	require.Equal(t, l, rw.contentLength)
	require.Equal(t, len("hello hello hello hello hello"), rw.contentLength)
	require.NoError(t, err)

	// Given:
	rw = respWriter{ResponseWriter: httptest.NewRecorder()}
	// When:
	rw.WriteHeader(http.StatusOK)
	l, err = rw.Write([]byte("hello"))
	// Then:
	require.Equal(t, http.StatusOK, rw.statusCode)
	require.Equal(t, l, rw.contentLength)
	require.Equal(t, len("hello"), rw.contentLength)
	require.NoError(t, err)

	// Given:
	rw = respWriter{ResponseWriter: httptest.NewRecorder()}
	// When:
	rw.WriteHeader(http.StatusInternalServerError)
	l, err = rw.Write([]byte("hello"))
	// Then:
	require.Equal(t, http.StatusInternalServerError, rw.statusCode)
	require.Equal(t, l, rw.contentLength)
	require.Equal(t, len("hello"), rw.contentLength)
	require.NoError(t, err)
}
