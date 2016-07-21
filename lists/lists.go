package lists

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	mailchimp "github.com/beeker1121/mailchimp-go"
	"github.com/beeker1121/mailchimp-go/query"
)

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
		if s.CampaignLastSent, err = time.Parse(time.RFC3339, aux.CampaignLastSent); err != nil {
			return err
		}
	}
	if aux.LastSubDate != "" {
		if s.LastSubDate, err = time.Parse(time.RFC3339, aux.LastSubDate); err != nil {
			return err
		}
	}
	if aux.LastUnsubDate != "" {
		if s.LastUnsubDate, err = time.Parse(time.RFC3339, aux.LastUnsubDate); err != nil {
			return err
		}
	}

	return nil
}

// List defines a list.
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
		if l.DateCreated, err = time.Parse(time.RFC3339, aux.DateCreated); err != nil {
			return err
		}
	}

	return nil
}

// Lists defines a set of lists.
type Lists struct {
	Lists      []List `json:"lists,omitempty"`
	TotalItems int    `json:"total_items"`
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
	Fields                 []string  `url:"fields,omitempty"`
	ExcludeFields          []string  `url:"exclude_fields,omitempty"`
	Count                  int       `url:"count,omitempty"`
	Offset                 int       `url:"offset,omitempty"`
	BeforeDateCreated      time.Time `url:"before_date_created,omitempty"`
	SinceDateCreated       time.Time `url:"since_date_created,omitempty"`
	BeforeCampaignLastSent time.Time `url:"before_campaign_last_sent,omitempty"`
	SinceCampaignLastSent  time.Time `url:"since_campaign_last_sent,omitempty"`
	Email                  string    `url:"email,omitempty"`
}

// EncodeQueryString handles custom query string encoding for the
// GetParams object.
func (gp *GetParams) EncodeQueryString(v interface{}) (string, error) {
	var beforeDateCreated string
	var sinceDateCreated string
	var beforeCampaignLastSent string
	var sinceCampaignLastSent string

	if !gp.BeforeDateCreated.IsZero() {
		beforeDateCreated = gp.BeforeDateCreated.Format(time.RFC3339)
	}
	if !gp.SinceDateCreated.IsZero() {
		sinceDateCreated = gp.SinceDateCreated.Format(time.RFC3339)
	}
	if !gp.BeforeCampaignLastSent.IsZero() {
		beforeCampaignLastSent = gp.BeforeCampaignLastSent.Format(time.RFC3339)
	}
	if !gp.SinceCampaignLastSent.IsZero() {
		sinceCampaignLastSent = gp.SinceCampaignLastSent.Format(time.RFC3339)
	}

	return query.Encode(struct {
		Fields                 string `url:"fields,omitempty"`
		ExcludeFields          string `url:"exclude_fields,omitempty"`
		Count                  int    `url:"count,omitempty"`
		Offset                 int    `url:"offset,omitempty"`
		BeforeDateCreated      string `url:"before_date_created,omitempty"`
		SinceDateCreated       string `url:"since_date_created,omitempty"`
		BeforeCampaignLastSent string `url:"before_campaign_last_sent,omitempty"`
		SinceCampaignLastSent  string `url:"since_campaign_last_sent,omitempty"`
		Email                  string `url:"email,omitempty"`
	}{
		Fields:                 strings.Join(gp.Fields, ","),
		ExcludeFields:          strings.Join(gp.ExcludeFields, ","),
		Count:                  gp.Count,
		Offset:                 gp.Offset,
		BeforeDateCreated:      beforeDateCreated,
		SinceDateCreated:       sinceDateCreated,
		BeforeCampaignLastSent: beforeCampaignLastSent,
		SinceCampaignLastSent:  sinceCampaignLastSent,
		Email: gp.Email,
	})
}

// GetListParams defines the available parameters that can be used
// when getting information on a specific list via the GetList
// function.
type GetListParams struct {
	Fields        []string `url:"fields,omitempty"`
	ExcludeFields []string `url:"exclude_fields,omitempty"`
}

// EncodeQueryString handles custom query string encoding for the
// GetListParams object.
func (gmp *GetListParams) EncodeQueryString(v interface{}) (string, error) {
	return query.Encode(struct {
		Fields        string `url:"fields,omitempty"`
		ExcludeFields string `url:"exclude_fields,omitempty"`
	}{
		Fields:        strings.Join(gmp.Fields, ","),
		ExcludeFields: strings.Join(gmp.ExcludeFields, ","),
	})
}

// UpdateParams defines the available parameters that can be used when
// updating a list via the Update function.
type UpdateParams struct {
	Name                string            `json:"name,omitempty"`
	Contact             *Contact          `json:"contact,omitempty"`
	PermissionReminder  string            `json:"permission_reminder,omitempty"`
	UseArchiveBar       bool              `json:"use_archive_bar,omitempty"`
	CampaignDefaults    *CampaignDefaults `json:"campaign_defaults,omitempty"`
	NotifyOnSubscribe   string            `json:"notify_on_subscribe,omitempty"`
	NotifyOnUnsubscribe string            `json:"notify_on_unsubscribe,omitempty"`
	EmailTypeOption     bool              `json:"email_type_option,omitempty"`
	Visibility          Visibility        `json:"visibility,omitempty"`
}

// New creates a new list.
func New(params *NewParams) (*List, error) {
	res := &List{}

	if params == nil {
		if err := mailchimp.Call("POST", "lists", nil, nil, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if err := mailchimp.Call("POST", "lists", nil, params, res); err != nil {
		return nil, err
	}
	return res, nil
}

// Get retrieves information about all lists.
func Get(params *GetParams) (*Lists, error) {
	res := &Lists{}

	if params == nil {
		if err := mailchimp.Call("GET", "lists", nil, nil, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if err := mailchimp.Call("GET", "lists", params, nil, res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetList retrieves information about a specific list.
func GetList(listID string, params *GetListParams) (*List, error) {
	res := &List{}
	path := fmt.Sprintf("lists/%s", listID)

	if params == nil {
		if err := mailchimp.Call("GET", path, nil, nil, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if err := mailchimp.Call("GET", path, params, nil, res); err != nil {
		return nil, err
	}
	return res, nil
}

// Update updates a list.
func Update(listID string, params *UpdateParams) (*List, error) {
	res := &List{}
	path := fmt.Sprintf("lists/%s", listID)

	if params == nil {
		if err := mailchimp.Call("PATCH", path, nil, nil, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if err := mailchimp.Call("PATCH", path, nil, params, res); err != nil {
		return nil, err
	}
	return res, nil
}

// Delete deletes a list.
func Delete(listID string) error {
	path := fmt.Sprintf("lists/%s", listID)
	return mailchimp.Call("DELETE", path, nil, nil, nil)
}
