package keenable

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

type captured struct {
	method, path, apiKey, title string
	body                        []byte
}

// withServer points both endpoints at a test server. The handler only records
// the request and writes the response; assertions run in the test goroutine.
func withServer(t *testing.T, respond func(w http.ResponseWriter, r *http.Request)) <-chan captured {
	t.Helper()
	reqs := make(chan captured, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		reqs <- captured{r.Method, r.URL.Path, r.Header.Get("X-API-Key"), r.Header.Get("X-Keenable-Title"), body}
		respond(w, r)
	}))
	usePrefix(t, server.URL)
	t.Cleanup(server.Close)
	return reqs
}

func usePrefix(t *testing.T, prefix string) {
	oldPublic, oldKeyed := publicSearchURL, keyedSearchURL
	publicSearchURL, keyedSearchURL = prefix+"/public", prefix+"/keyed"
	t.Cleanup(func() { publicSearchURL, keyedSearchURL = oldPublic, oldKeyed })
}

func newClient(key string) *Client {
	c := NewClient()
	c.ApiKey.Value = key
	return c
}

func TestSearchWithoutKeyUsesPublicEndpoint(t *testing.T) {
	reqs := withServer(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"results":[
			{"title":"Go generics","url":"https://go.dev/doc/tutorial/generics","description":"Getting started with generics.","snippet":"long"},
			{"title":"No URL","description":"dropped"},
			{"title":"Snippet only","url":"https://example.com","snippet":"` + strings.Repeat("x", 1500) + `"}
		]}`))
	})

	got, err := newClient("").Search("golang generics")
	require.NoError(t, err)
	r := <-reqs
	require.Equal(t, http.MethodPost, r.method)
	require.Equal(t, "/public", r.path)
	require.Empty(t, r.apiKey)
	require.Equal(t, "Fabric", r.title)
	var body map[string]any
	require.NoError(t, json.Unmarshal(r.body, &body))
	require.Equal(t, "golang generics", body["query"])
	require.EqualValues(t, 10, body["max_results"])

	require.True(t, strings.HasPrefix(got, "## Go generics\nhttps://go.dev/doc/tutorial/generics\n\nGetting started with generics.\n\n## Snippet only\nhttps://example.com\n\n"))
	require.NotContains(t, got, "No URL")
	require.True(t, strings.HasSuffix(got, strings.Repeat("x", 1000)+"…"))
}

func TestSearchWithKeyUsesKeyedEndpoint(t *testing.T) {
	reqs := withServer(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"results":[]}`))
	})

	_, err := newClient("secret").Search("anything")
	require.NoError(t, err)
	r := <-reqs
	require.Equal(t, "/keyed", r.path)
	require.Equal(t, "secret", r.apiKey)
}

func TestSearchReturnsAPIError(t *testing.T) {
	withServer(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"detail":"rate limited"}`))
	})

	_, err := newClient("").Search("anything")
	require.ErrorContains(t, err, "status 429")
	require.ErrorContains(t, err, "rate limited")
}

func TestSearchRejectsRedirects(t *testing.T) {
	for _, key := range []string{"", "secret"} {
		for _, code := range []int{http.StatusFound, http.StatusTemporaryRedirect, http.StatusPermanentRedirect} {
			t.Run(fmt.Sprintf("key=%t/%d", key != "", code), func(t *testing.T) {
				var hits atomic.Int32
				other := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
					hits.Add(1)
					_, _ = w.Write([]byte(`{"results":[]}`))
				}))
				t.Cleanup(other.Close)
				withServer(t, func(w http.ResponseWriter, r *http.Request) {
					http.Redirect(w, r, other.URL+"/steal", code)
				})

				_, err := newClient(key).Search("private query")
				require.ErrorIs(t, err, errRedirectRejected)
				require.Zero(t, hits.Load(), "redirect destination must not be contacted")
			})
		}
	}
}

func TestSearchRejectsHTTPSToHTTPDowngrade(t *testing.T) {
	var hits atomic.Int32
	plain := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		hits.Add(1)
	}))
	t.Cleanup(plain.Close)
	tlsServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, plain.URL+"/keyed", http.StatusPermanentRedirect)
	}))
	t.Cleanup(tlsServer.Close)
	usePrefix(t, tlsServer.URL)
	oldClient := httpClient
	httpClient = newHTTPClient(tlsServer.Client().Transport)
	t.Cleanup(func() { httpClient = oldClient })

	for _, key := range []string{"", "secret"} {
		_, err := newClient(key).Search("private query")
		require.ErrorIs(t, err, errRedirectRejected)
	}
	require.Zero(t, hits.Load())
}

func paddedJSON(size int) string {
	const doc = `{"results":[{"title":"t","url":"https://example.com","description":"d"}]}`
	return doc + strings.Repeat(" ", size-len(doc))
}

func TestSearchResponseSizeBoundary(t *testing.T) {
	cases := []struct {
		name    string
		body    string
		tooBig  bool
		wantErr bool
	}{
		{"at limit", paddedJSON(maxResponseBytes), false, false},
		{"one byte over", paddedJSON(maxResponseBytes + 1), true, true},
		{"valid prefix then excess bytes", paddedJSON(maxResponseBytes) + "x", true, true},
		{"valid JSON then trailing data", `{"results":[]}garbage`, false, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			withServer(t, func(w http.ResponseWriter, _ *http.Request) {
				_, _ = io.WriteString(w, tc.body)
			})
			got, err := newClient("").Search("q")
			if !tc.wantErr {
				require.NoError(t, err)
				require.Contains(t, got, "https://example.com")
				return
			}
			require.Error(t, err)
			if tc.tooBig {
				require.ErrorIs(t, err, errResponseTooLarge)
			} else {
				require.NotErrorIs(t, err, errResponseTooLarge)
			}
		})
	}
}

func TestSearchErrorDiagnosticIsBounded(t *testing.T) {
	for _, size := range []int{maxErrorBodyBytes, maxErrorBodyBytes + 1, 10 * maxResponseBytes} {
		t.Run(fmt.Sprint(size), func(t *testing.T) {
			withServer(t, func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusBadGateway)
				_, _ = io.WriteString(w, strings.Repeat("e", size))
			})
			_, err := newClient("").Search("q")
			require.ErrorContains(t, err, "status 502")
			msg := err.Error()
			require.Less(t, len(msg), maxErrorBodyBytes+200)
			require.Equal(t, size > maxErrorBodyBytes, strings.Contains(msg, "(truncated)"))
		})
	}
}

func TestSearchErrorDiagnosticIsPrintable(t *testing.T) {
	withServer(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("bad\x1b[31m\nrequest\xff"))
	})
	_, err := newClient("").Search("q")
	require.ErrorContains(t, err, "status 400")
	require.NotContains(t, err.Error(), "\x1b")
	require.NotContains(t, err.Error(), "\n")
}

func TestSearchPreservesBodyReadFailures(t *testing.T) {
	for _, status := range []int{http.StatusOK, http.StatusInternalServerError} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			withServer(t, func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Length", "1000")
				w.WriteHeader(status)
				_, _ = w.Write([]byte(`{"results":`))
				conn, _, err := http.NewResponseController(w).Hijack()
				if err == nil {
					_ = conn.Close()
				}
			})
			_, err := newClient("").Search("q")
			require.ErrorContains(t, err, "reading response body")
			require.NotErrorIs(t, err, errResponseTooLarge)
			if status != http.StatusOK {
				require.ErrorContains(t, err, "status 500")
			}
		})
	}
}

func TestSearchLimitsResultsAndFields(t *testing.T) {
	var sb strings.Builder
	sb.WriteString(`{"results":[`)
	for i := range 50 {
		if i > 0 {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, `{"title":%q,"url":"https://example.com/%d","description":%q}`,
			strings.Repeat("T", 500), i, strings.Repeat("d", 5000))
	}
	sb.WriteString(`]}`)
	withServer(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, sb.String())
	})

	got, err := newClient("").Search("q")
	require.NoError(t, err)
	require.Equal(t, maxResults, strings.Count(got, "## "))
	require.NotContains(t, got, strings.Repeat("T", titleMaxChars+1))
	require.NotContains(t, got, strings.Repeat("d", snippetMaxChars+1))
}
