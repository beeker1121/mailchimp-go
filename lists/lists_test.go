package lists

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	mailchimp "github.com/beeker1121/mailchimp-go"
)

var timeString = "2020-01-02 23:59:59 +0000 UTC"
var timeType = reflect.TypeOf(time.Time{})

func TestListUnmarshal(t *testing.T) {
	data := []byte(`{
		"date_created": "2020-01-02T23:59:59+00:00"
	}`)

	list := &List{}
	if err := json.NewDecoder(bytes.NewBuffer(data)).Decode(list); err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(list.DateCreated) != timeType {
		t.Errorf("Expected list.DateCreated to be of type time.Time, got %v", reflect.TypeOf(list.DateCreated))
	}

	if list.DateCreated.String() != timeString {
		t.Errorf("Expected list.DateCreated.String() to equal %s, got %s", timeString, list.DateCreated.String())
	}
}

func TestDelete(t *testing.T) {
	list, err := createList()
	if err != nil {
		t.Error(err)
	}

	gotList, err := GetList(list.ID, nil)
	if err != nil {
		t.Error(err)
	}

	if gotList.ID != list.ID {
		t.Error("Expected gotList.ID to equal list.ID")
	}

	if err = Delete(list.ID); err != nil {
		t.Error(err)
	}

	_, err = GetList(list.ID, nil)

	if err == nil {
		t.Error("Expected error to be non-nil")
	}

	apiErr := err.(*mailchimp.APIError)

	if apiErr.Status != 404 {
		t.Errorf("Expected err.Status to be 404, got %d", apiErr.Status)
	}
}

func TestNew(t *testing.T) {
	list, err := createList()
	if err != nil {
		t.Error(err)
	}

	if list.Name != "mailchimp-go Test List" {
		t.Errorf("Expected list.Name to equal \"Acme Corp\", got %s", list.Name)
	}

	if err = Delete(list.ID); err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	params := &NewParams{
		Name: "mailchimp-go Test List 1",
		Contact: &Contact{
			Company:  "Acme Corp",
			Address1: "123 Main St",
			City:     "Chicago",
			State:    "IL",
			Zip:      "60613",
			Country:  "United States",
		},
		PermissionReminder: "You opted to receive updates on Acme Corp",
		CampaignDefaults: &CampaignDefaults{
			FromName:  "John Doe",
			FromEmail: "newsletter@acmecorp.com",
			Subject:   "Newsletter",
			Language:  "en",
		},
		EmailTypeOption: false,
		Visibility:      VisibilityPublic,
	}

	list1, err := New(params)
	if err != nil {
		t.Error(err)
	}

	params.Name = "mailchimp-go Test List 2"

	list2, err := New(params)
	if err != nil {
		t.Error(err)
	}

	lists, err := Get(nil)
	if err != nil {
		t.Error(err)
	}

	if lists.Lists[len(lists.Lists)-2].ID != list1.ID {
		t.Errorf("Expected second to last list to equal %s, got %s", list1.ID, lists.Lists[0].ID)
	}
	if lists.Lists[len(lists.Lists)-1].ID != list2.ID {
		t.Errorf("Expected last list to equal %s, got %s", list2.ID, lists.Lists[1].ID)
	}

	if err = Delete(list1.ID); err != nil {
		t.Error(err)
	}
	if err = Delete(list2.ID); err != nil {
		t.Error(err)
	}
}

func TestGetList(t *testing.T) {
	list, err := createList()
	if err != nil {
		t.Error(err)
	}

	gotList, err := GetList(list.ID, nil)
	if err != nil {
		t.Error(err)
	}

	if gotList.ID != list.ID {
		t.Error("Expected gotList.ID to equal list.ID")
	}

	if err = Delete(list.ID); err != nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {
	list, err := createList()
	if err != nil {
		t.Error(err)
	}

	updateParams := &UpdateParams{
		Name: "mailchimp-go Test List 2",
	}

	updatedList, err := Update(list.ID, updateParams)
	if err != nil {
		t.Error(err)
	}

	gotList, err := GetList(list.ID, nil)
	if err != nil {
		t.Error(err)
	}

	if gotList.Name != "mailchimp-go Test List 2" {
		t.Errorf("Expected gotList.Name to equal \"mailchimp-go Test List 2\", got %s", gotList.Name)
	}
	if gotList.Name != updatedList.Name {
		t.Error("Expected gotList.Name to equal updatedList.Name")
	}

	if err = Delete(list.ID); err != nil {
		t.Error(err)
	}
}

func createList() (*List, error) {
	params := &NewParams{
		Name: "mailchimp-go Test List",
		Contact: &Contact{
			Company:  "Acme Corp",
			Address1: "123 Main St",
			City:     "Chicago",
			State:    "IL",
			Zip:      "60613",
			Country:  "United States",
		},
		PermissionReminder: "You opted to receive updates on Acme Corp",
		CampaignDefaults: &CampaignDefaults{
			FromName:  "John Doe",
			FromEmail: "newsletter@acmecorp.com",
			Subject:   "Newsletter",
			Language:  "en",
		},
		EmailTypeOption: false,
		Visibility:      VisibilityPublic,
	}

	return New(params)
}

func TestMain(m *testing.M) {
	if err := mailchimp.SetKey(os.Getenv("MAILCHIMP_API_KEY")); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}
