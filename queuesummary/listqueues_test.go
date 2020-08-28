package queuesummary

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListQueues(t *testing.T) {
	t.Run("Sends the request and receives the response", func(t *testing.T) {
		request := newListQueuesRequest()
		response := httptest.NewRecorder()
		ListQueues(response, request)
		fmt.Printf(response.Body.String())
		assertStatus(t, response.Code, 200)
	})
}

func newListQueuesRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/queues-summary/abc", nil)
	return req
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
