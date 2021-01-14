package clerk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestClientsService_ListAll_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := "[" + dummyClientResponseJson + "]"

	mux.HandleFunc("/clients", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, expectedResponse)
	})

	var want []ClientResponse
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Clients().ListAll()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d clients to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestClientsService_ListAll_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	response, err := client.Clients().ListAll()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if response != nil {
		t.Errorf("Was not expecting any clients to be returned, instead got %v", response)
	}
}

func TestClientsService_Read_happyPath(t *testing.T) {
	token := "token"
	clientId := "someClientId"
	expectedResponse := dummyClientResponseJson

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/clients/"+clientId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, expectedResponse)
	})

	var want ClientResponse
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Clients().Read(clientId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestClientsService_Read_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	response, err := client.Clients().Read("someClientId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if response != nil {
		t.Errorf("Was not expecting any client to be returned, instead got %v", response)
	}
}

const dummyClientResponseJson = `{
        "ended": false,
        "id": "client_1mvnkzXhKhn9pDjp1f4x1id6pQZ",
        "last_active_session_id": "sess_1mvnlhBAwUv8EktIKbooqqYwsUW",
        "object": "client",
        "sign_in_attempt_id": null,
        "sign_up_attempt_id": null
    }`