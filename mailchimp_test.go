package mailchimp

import "testing"

func TestSetKey(t *testing.T) {
	if err := SetKey("123-123-123"); err != ErrAPIKeyFormat {
		t.Errorf("Expected to get ErrAPIKeyFormat, got %s", err.Error())
	}

	if err := SetKey("123-123"); err != nil {
		t.Errorf("Expected to get nil error, got %s", err.Error())
	}
}
