package lists

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

var timeString = "2020-01-02 23:59:59 +0000 UTC"
var timeType = reflect.TypeOf(time.Time{})

func TestListUnmarshal(t *testing.T) {
	data := []byte(`{
		"date_created": "2020-01-02 23:59:59"
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
