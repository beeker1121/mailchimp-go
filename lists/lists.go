package lists

import (
	"encoding/json"
	"time"

	mailchimp "github.com/beeker1121/mailchimp-go"
)

// timeFormat defines the MailChimp timestamp format.
const timeFormat string = "2006-01-02 15:04:05"

// Contact defines the contact information of the list owner, which
// is displayed in the footer of campaigns.
type Contact struct {
	Company  string `json:"company"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2,omitempty"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
	Country  string `json:"country"`
	Phone    string `json:"phone,omitempty"`
}

// CampaignDefaults defines the default values for campaigns.
type CampaignDefaults struct {
	FromName  string `json:"from_name"`
	FromEmail string `json:"from_email"`
	Subject   string `json:"subject"`
	Language  string `json:"language"`
}

// Visibility defines whether a list is public or private
type Visibility string

// The visibility definitions.
const (
	VisibilityPublic  Visibility = "pub"
	VisibilityPrivate            = "prv"
)

// Stats defines statistics for a list.
type Stats struct {
	MemberCount               uint      `json:"member_count,omitempty"`
	UnsubscribeCount          uint      `json:"unsubscribe_count,omitempty"`
	CleanedCount              uint      `json:"cleaned_count,omitempty"`
	MemberCountSinceSend      uint      `json:"member_count_since_send,omitempty"`
	UnsubscribeCountSinceSend uint      `json:"unsubscribe_count_since_send,omitempty"`
	CleanedCountSinceSend     uint      `json:"cleaned_count_since_send,omitempty"`
	CampaignCount             uint      `json:"campaign_count,omitempty"`
	CampaignLastSent          time.Time `json:"campaign_last_sent,omitempty"`
	MergeFieldCount           uint      `json:"merge_field_count,omitempty"`
	AvgSubRate                float64   `json:"avg_sub_rate,omitempty"`
	AvgUnsubRate              float64   `json:"avg_unsub_rate,omitempty"`
	TargetSubRate             float64   `json:"target_sub_rate,omitempty"`
	OpenRate                  float32   `json:"open_rate,omitempty"`
	ClickRate                 float32   `json:"click_rate,omitempty"`
	LastSubDate               time.Time `json:"last_sub_date,omitempty"`
	LastUnsubDate             time.Time `json:"last_unsub_date,omitempty"`
}

// UnmarshalJSON handles custom JSON unmarshalling for the Stats object.
// Credit to http://choly.ca/post/go-json-marshalling/
func (s *Stats) UnmarshalJSON(data []byte) error {
	var err error
	type alias Stats

	aux := &struct {
		*alias
		CampaignLastSent string `json:"campaign_last_sent,omitempty"`
		LastSubDate      string `json:"last_sub_date,omitempty"`
		LastUnsubDate    string `json:"last_unsub_date,omitempty"`
	}{
		alias: (*alias)(s),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	if aux.CampaignLastSent != "" {
		if s.CampaignLastSent, err = time.Parse(timeFormat, aux.CampaignLastSent); err != nil {
			return err
		}
	}
	if aux.LastSubDate != "" {
		if s.LastSubDate, err = time.Parse(timeFormat, aux.LastSubDate); err != nil {
			return err
		}
	}
	if aux.LastUnsubDate != "" {
		if s.LastUnsubDate, err = time.Parse(timeFormat, aux.LastUnsubDate); err != nil {
			return err
		}
	}

	return nil
}

// List defines a list.
//
// Schema: https://api.mailchimp.com/schema/3.0/Lists/Instance.json
type List struct {
	ID                  string            `json:"id,omitempty"`
	Name                string            `json:"name"`
	Contact             *Contact          `json:"contact"`
	PermissionReminder  string            `json:"permission_reminder"`
	UseArchiveBar       bool              `json:"use_archive_bar,omitempty"`
	CampaignDefaults    *CampaignDefaults `json:"campaign_defaults"`
	NotifyOnSubscribe   string            `json:"notify_on_subscribe,omitempty"`
	NotifyOnUnsubscribe string            `json:"notify_on_unsubscribe,omitempty"`
	DateCreated         time.Time         `json:"date_created,omitempty"`
	ListRating          uint8             `json:"list_rating,omitempty"`
	EmailTypeOption     bool              `json:"email_type_option"`
	SubscribeUrlShort   string            `json:"subscribe_url_short,omitempty"`
	SubscribeUrlLong    string            `json:"subscribe_url_long,omitempty"`
	BeamerAddress       string            `json:"beamer_address,omitempty"`
	Visibility          Visibility        `json:"visibility"`
	Modules             []string          `json:"modules,omitempty"`
	Stats               *Stats            `json:"stats,omitempty"`
}

// UnmarshalJSON handles custom JSON unmarshalling for the List object.
// Credit to http://choly.ca/post/go-json-marshalling/
func (l *List) UnmarshalJSON(data []byte) error {
	var err error
	type alias List

	aux := &struct {
		*alias
		DateCreated string `json:"date_created,omitempty"`
	}{
		alias: (*alias)(l),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	if aux.DateCreated != "" {
		if l.DateCreated, err = time.Parse(timeFormat, aux.DateCreated); err != nil {
			return err
		}
	}

	return nil
}

// NewParams defines the available parameters that can be used when
// creating a new list via the New function.
type NewParams struct {
	Name                string            `json:"name"`
	Contact             *Contact          `json:"contact"`
	PermissionReminder  string            `json:"permission_reminder"`
	UseArchiveBar       bool              `json:"use_archive_bar,omitempty"`
	CampaignDefaults    *CampaignDefaults `json:"campaign_defaults"`
	NotifyOnSubscribe   string            `json:"notify_on_subscribe,omitempty"`
	NotifyOnUnsubscribe string            `json:"notify_on_unsubscribe,omitempty"`
	EmailTypeOption     bool              `json:"email_type_option"`
	Visibility          Visibility        `json:"visibility"`
}

// GetParams defines the available parameters that can be used when
// getting information about all lists via the Get function.
type GetParams struct {
	Fields                 string `url:"fields,omitempty"`
	ExcludeFields          string `url:"exclude_fields,omitempty"`
	Count                  int    `url:"count,omitempty"`
	Offset                 int    `url:"offset,omitempty"`
	BeforeDateCreated      string `url:"before_date_created,omitempty"`
	SinceDateCreated       string `url:"since_date_created,omitempty"`
	BeforeCampaignLastSent string `url:"before_campaign_last_sent,omitempty"`
	SinceCampaignLastSent  string `url:"since_campaign_last_sent,omitempty"`
	Email                  string `url:"email,omitempty"`
}

// New creates a new list.
//
// Method:     POST
// Resource:   /lists
// Definition: http://developer.mailchimp.com/documentation/mailchimp/reference/lists/#create-post_lists
func New(params *NewParams) (*List, error) {
	res := &List{}
	if err := mailchimp.Call("POST", "lists", nil, params, res); err != nil {
		return nil, err
	}
	return res, nil
}
