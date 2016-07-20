package members

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

var timeString = "2020-01-02 23:59:59 +0000 UTC"
var timeType = reflect.TypeOf(time.Time{})

func TestMemberUnmarshal(t *testing.T) {
	data := []byte(`{
		"timestamp_signup": "2020-01-02 23:59:59"
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
