package httpsrv

import (
	"context"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	type testCase struct {
		givenOpts []ServerOption
		expErr    error
		expSrv    *Server
	}
	tcs := map[string]testCase{
		"default": {
			expSrv: &Server{
				srv: &http.Server{
					Addr:         ":9000",
					ReadTimeout:  5 * time.Second,
					WriteTimeout: 10 * time.Second,
					IdleTimeout:  120 * time.Second,
				},
				gracefulShutdownTimeout: 10 * time.Second,
			},
		},
		"override port": {
			givenOpts: []ServerOption{ServerPort(3000)},
			expSrv: &Server{
				srv: &http.Server{
					Addr:         ":3000",
					ReadTimeout:  5 * time.Second,
					WriteTimeout: 10 * time.Second,
					IdleTimeout:  120 * time.Second,
				},
				gracefulShutdownTimeout: 10 * time.Second,
			},
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given && When:
			srv, err := New(nil, tc.givenOpts...)

			// Then:
			require.Equal(t, tc.expErr, err)
			if tc.expSrv != nil {
				require.EqualValues(t, tc.expSrv.srv, srv.srv)
				require.EqualValues(t, tc.expSrv.gracefulShutdownTimeout, srv.gracefulShutdownTimeout)
			} else {
				require.Nil(t, srv)
			}
		})
	}
}

func TestServer_Start(t *testing.T) {
	// Given:
	srv, err := New(nil)
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	// When:
	go func() {
		defer wg.Done()
		err = srv.Start(ctx)
	}()

	// Then:
	time.Sleep(2 * time.Second)

	cancel()

	wg.Wait()

	require.NoError(t, err)
}
