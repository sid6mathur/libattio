package attio

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Attio adds response fields over time; a client built from an older spec must still decode them.
func TestNewClientAuthenticatesAndToleratesNewResponseFields(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer secret" {
			t.Errorf("Authorization = %q, want %q", got, "Bearer secret")
		}
		if r.Method != http.MethodPut || r.URL.Path != "/v2/objects/people/records" || r.URL.Query().Get("matching_attribute") != "email_addresses" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data": {
			"id": {"workspace_id": "w", "object_id": "o", "record_id": "r1", "field_added_later": 1},
			"created_at": "2026-11-23T00:00:00Z",
			"web_url": "https://app.attio.com/x/person/r1",
			"values": {},
			"another_field_added_later": {"nested": true}
		}}`))
	}))
	defer srv.Close()

	client := NewClient("secret")
	client.GetConfig().Servers = ServerConfigurations{{URL: srv.URL}}

	resp, _, err := client.RecordsAPI.V2ObjectsObjectRecordsPut(context.Background(), "people").
		MatchingAttribute("email_addresses").
		V2ObjectsObjectRecordsPutRequest(V2ObjectsObjectRecordsPutRequest{
			Data: V2ObjectsObjectRecordsPutRequestData{
				Values: map[string][]interface{}{"email_addresses": {"ada@example.com"}},
			},
		}).Execute()
	if err != nil {
		t.Fatalf("Execute: %v", err)
	}
	if resp.Data.Id.RecordId != "r1" || resp.Data.WebUrl != "https://app.attio.com/x/person/r1" {
		t.Errorf("decoded %+v", resp.Data)
	}
}
