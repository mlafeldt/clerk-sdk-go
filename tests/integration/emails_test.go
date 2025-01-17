//go:build integration
// +build integration

package integration

import (
	"reflect"
	"testing"

	"github.com/clerkinc/clerk-sdk-go/clerk"
)

func TestEmails(t *testing.T) {
	client := createClient()

	users, _ := client.Users().ListAll(clerk.ListAllUsersParams{})
	if users == nil || len(users) == 0 {
		return
	}

	user := users[0]

	if user.PrimaryEmailAddressID == nil {
		return
	}

	email := clerk.Email{
		FromEmailName:  "integration-test",
		Subject:        "Testing Go SDK",
		Body:           "Testing email functionality for Go SDK",
		EmailAddressID: *user.PrimaryEmailAddressID,
	}

	got, err := client.Emails().Create(email)
	if err != nil {
		t.Fatalf("Emails.Create returned error: %v", err)
	}

	want := clerk.EmailResponse{
		ID:             got.ID,
		Object:         "email",
		Status:         "queued",
		ToEmailAddress: got.ToEmailAddress,
		Email: clerk.Email{
			FromEmailName:  email.FromEmailName,
			Subject:        email.Subject,
			Body:           email.Body,
			EmailAddressID: email.EmailAddressID,
		},
	}

	if !reflect.DeepEqual(*got, want) {
		t.Fatalf("Emails.Create(%v) got: %v, wanted %v", email, got, want)
	}
}
