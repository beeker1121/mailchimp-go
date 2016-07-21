package members

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	mailchimp "github.com/beeker1121/mailchimp-go"
	"github.com/beeker1121/mailchimp-go/lists"
)

var listID string
var timeString = "2020-01-02 23:59:59 +0000 UTC"
var timeType = reflect.TypeOf(time.Time{})

func TestMemberUnmarshal(t *testing.T) {
	data := []byte(`{
		"timestamp_signup": "2020-01-02T23:59:59+00:00"
	}`)

	member := &Member{}
	if err := json.NewDecoder(bytes.NewBuffer(data)).Decode(member); err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(member.TimestampSignup) != timeType {
		t.Errorf("Expected member.TimestampSignup to be of type time.Time, got %v", reflect.TypeOf(member.TimestampSignup))
	}

	if member.TimestampSignup.String() != timeString {
		t.Errorf("Expected member.TimestampSignup.String() to equal %s, got %s", timeString, member.TimestampSignup.String())
	}
}

func TestDelete(t *testing.T) {
	params := &NewParams{
		EmailAddress: "mailchimp-go-test@github.com",
		Status:       StatusPending,
	}

	member, err := New(listID, params)
	if err != nil {
		t.Error(err)
	}

	gotMember, err := GetMember(listID, member.ID, nil)
	if err != nil {
		t.Error(err)
	}

	if gotMember.ID != member.ID {
		t.Error("Expected gotMember.ID to equal member.ID")
	}

	if err = Delete(listID, member.ID); err != nil {
		t.Error(err)
	}

	_, err = GetMember(listID, member.ID, nil)

	apiErr := err.(*mailchimp.APIError)

	if apiErr.Status != 404 {
		t.Errorf("Expected err.Status to be 404, got %d", apiErr.Status)
	}
}

func TestNew(t *testing.T) {
	params := &NewParams{
		EmailAddress: "mailchimp-go-test@github.com",
		Status:       StatusPending,
	}

	member, err := New(listID, params)
	if err != nil {
		t.Error(err)
	}

	if member.EmailAddress != "mailchimp-go-test@github.com" {
		t.Errorf("Expected member.EmailAddress to equal \"mailchimp-go-test@github.com\", got %s", member.EmailAddress)
	}

	if err = Delete(listID, member.ID); err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	params := &NewParams{
		EmailAddress: "mailchimp-go-test1@github.com",
		Status:       StatusPending,
	}

	member1, err := New(listID, params)
	if err != nil {
		t.Error(err)
	}

	params.EmailAddress = "mailchimp-go-test2@github.com"

	member2, err := New(listID, params)
	if err != nil {
		t.Error(err)
	}

	members, err := Get(listID, nil)
	if err != nil {
		t.Error(nil)
	}

	if members.Members[len(members.Members)-2].EmailAddress != member1.EmailAddress {
		t.Errorf("Expected second to last list member to equal %s, got %s", member1.EmailAddress, members.Members[len(members.Members)-2].EmailAddress)
	}
	if members.Members[len(members.Members)-1].EmailAddress != member2.EmailAddress {
		t.Errorf("Expected last list member to equal %s, got %s", member2.EmailAddress, members.Members[len(members.Members)-1].EmailAddress)
	}

	if err = Delete(listID, member1.ID); err != nil {
		t.Error(err)
	}
	if err = Delete(listID, member2.ID); err != nil {
		t.Error(err)
	}
}

func TestGetMember(t *testing.T) {
	params := &NewParams{
		EmailAddress: "mailchimp-go-test@github.com",
		Status:       StatusPending,
	}

	member, err := New(listID, params)
	if err != nil {
		t.Error(err)
	}

	gotMember, err := GetMember(listID, member.ID, nil)
	if err != nil {
		t.Error(err)
	}

	if gotMember.ID != member.ID {
		t.Error("Expected gotMember.ID to equal member.ID")
	}

	if err = Delete(listID, member.ID); err != nil {
		t.Error(err)
	}
}

func TestUpdateMember(t *testing.T) {
	params := &NewParams{
		EmailAddress: "mailchimp-go-test@github.com",
		Status:       StatusPending,
	}

	member, err := New(listID, params)
	if err != nil {
		t.Error(err)
	}

	timeSignup, err := time.Parse(time.RFC3339, "2020-01-02T23:59:59+00:00")
	if err != nil {
		t.Error(err)
	}

	updateParams := &UpdateParams{
		TimestampSignup: timeSignup,
	}

	updatedMember, err := Update(listID, member.ID, updateParams)
	if err != nil {
		t.Error(err)
	}

	gotMember, err := GetMember(listID, member.ID, nil)
	if err != nil {
		t.Error(err)
	}

	if gotMember.TimestampSignup.String() != timeString {
		t.Errorf("Expected gotMember.TimestampSignup.String() to equal %s, got %s", timeString, gotMember.TimestampSignup.String())
	}
	if gotMember.TimestampSignup.String() != updatedMember.TimestampSignup.String() {
		t.Error("Expected gotMember.TimestampSignup.String() to equal updatedMember.TimestampSignup.String()")
	}

	if err = Delete(listID, member.ID); err != nil {
		t.Error(err)
	}
}

func createList() (*lists.List, error) {
	listParams := &lists.NewParams{
		Name: "mailchimp-go Test List",
		Contact: &lists.Contact{
			Company:  "Acme Corp",
			Address1: "123 Main St",
			City:     "Chicago",
			State:    "IL",
			Zip:      "60613",
			Country:  "United States",
		},
		PermissionReminder: "You opted to receive updates on Acme Corp",
		CampaignDefaults: &lists.CampaignDefaults{
			FromName:  "John Doe",
			FromEmail: "newsletter@acmecorp.com",
			Subject:   "Newsletter",
			Language:  "en",
		},
		EmailTypeOption: false,
		Visibility:      lists.VisibilityPublic,
	}

	return lists.New(listParams)
}

func TestMain(m *testing.M) {
	if err := mailchimp.SetKey(os.Getenv("MAILCHIMP_API_KEY")); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	list, err := createList()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	listID = list.ID

	code := m.Run()

	if err = lists.Delete(list.ID); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(code)
}
