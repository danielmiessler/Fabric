package digitalocean

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// blockingTransport holds every request until its context is cancelled.
type blockingTransport struct{}

func (blockingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	<-req.Context().Done()
	return nil, req.Context().Err()
}

func TestListModelsControlPlaneHonorsContextCancellation(t *testing.T) {
	client := NewClient()
	client.ControlPlaneToken.Value = "token"
	client.httpClient = &http.Client{Transport: blockingTransport{}}

	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(50*time.Millisecond, cancel)

	done := make(chan error, 1)
	go func() {
		_, err := client.ListModels(ctx)
		done <- err
	}()

	select {
	case err := <-done:
		require.Error(t, err)
		require.True(t, errors.Is(err, context.Canceled), "got %v", err)
	case <-time.After(3 * time.Second):
		t.Fatal("ListModels ignored the cancelled context")
	}
}
