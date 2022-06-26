package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ismael3s/go-tests/01/handler"
)

func TestHealthCheckHandler(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/health-check", nil)

	if err != nil {
		t.Fatal("Failed to create request")
	}

	rr := httptest.NewRecorder()
	h := http.HandlerFunc(handler.HealthCheckHandler)

	h.ServeHTTP(rr, request)

	if rr.Code != http.StatusOK {
		t.Errorf("expected %d, got %d\n", http.StatusOK, rr.Code)
	}

	expected := `{"alive": true}`

	if rr.Body.String() != expected {
		t.Errorf("expected %s, got %s\n", expected, rr.Body.String())
	}

}

func TestRedirectHandler(t *testing.T) {
	t.Run("should redirect to correct url", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/google", nil)
		if err != nil {
			t.Fatal("Failed to create request")
		}

		paths := map[string]string{
			"/google": "http://127.0.0.1",
		}

		rr := httptest.NewRecorder()
		h := http.HandlerFunc(handler.RedirectHandler(paths))

		h.ServeHTTP(rr, request)

		if rr.Code != http.StatusFound {
			t.Errorf("expected %d, got %d\n", http.StatusFound, rr.Code)
		}

		expected := `http://127.0.0.1`

		got, _ := rr.Result().Location()

		if got.String() != expected {
			t.Errorf("expected %s, got %s\n", expected, rr.Body.String())
		}
	})

	t.Run("should redirect to health-check when no path is found", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/teste", nil)
		if err != nil {
			t.Fatal("Failed to create request")
		}

		paths := map[string]string{
			"/google": "http://127.0.0.1",
		}

		rr := httptest.NewRecorder()
		h := http.HandlerFunc(handler.RedirectHandler(paths))

		h.ServeHTTP(rr, request)

		if rr.Code != http.StatusOK {
			t.Errorf("expected %d, got %d\n", http.StatusOK, rr.Code)
		}

	})
}
