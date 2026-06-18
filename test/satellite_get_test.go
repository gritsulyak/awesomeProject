//go:build test_integration

package test

import (
	"net/http"
	"testing"
	"time"
)

func TestGetSatellite(t *testing.T) {
	app := createAndRunApp(t.Context())
	defer app.Stop(t.Context())

	time.Sleep(2 * time.Second)

	t.Run("get satellite", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "http://0.0.0.0:10080/api/v1/satellite/moon", nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}
