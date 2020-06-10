package contactflowssummary

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListContactFlows(t *testing.T) {
	t.Run("Sends the request and receives the response", func(t *testing.T) {
		request := newListContactFlowsRequest()
		response := httptest.NewRecorder()
		ListContactFlows(response, request)
		assertStatus(t, response.Code, 200)
	})
}

func newListContactFlowsRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/contact-flows-summary/abc", nil)
	return req
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
