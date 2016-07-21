package members

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	mailchimp "github.com/beeker1121/mailchimp-go"
	"github.com/beeker1121/mailchimp-go/query"
)

// EmailType defines the type of email a member asked to get.
type EmailType string

// The email type definitions.
const (
	EmailTypeHTML EmailType = "html"
	EmailTypeText           = "text"
)

// Status defines the subscription status for a given member
// within a List.
type Status string

// The subscription status definitions.
const (
	StatusSubscribed   Status = "subscribed"
	StatusUnsubscribed        = "unsubscribed"
	StatusCleaned             = "cleaned"
	StatusPending             = "pending"
)

// Stats defines the open and click rates for a member.
type Stats struct {
	AvgOpenRate  float32 `json:"avg_open_rate,omitempty"`
	AvgClickRate float32 `json:"avg_click_rate,omitempty"`
}

// Location defines a member's location information.
type Location struct {
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
	GMTOff      int     `json:"gmtoff,omitempty"`
	DSTOff      int     `json:"dstoff,omitempty"`
	CountryCode string  `json:"country_code,omitempty"`
	Timezone    string  `json:"timezone,omitempty"`
}

// Note defines a note about a member.
type Note struct {
	NoteID    int       `json:"note_id"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	Note      string    `json:"note"`
}

// UnmarshalJSON handles custom JSON unmarshalling for the Note object.
// Credit to http://choly.ca/post/go-json-marshalling/
func (n *Note) UnmarshalJSON(data []byte) error {
	var err error
	type alias Note

	aux := &struct {
		*alias
		CreatedAt string `json:"created_at"`
	}{
		alias: (*alias)(n),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	if aux.CreatedAt != "" {
		if n.CreatedAt, err = time.Parse(time.RFC3339, aux.CreatedAt); err != nil {
			return err
		}
	}

	return nil
}

// Member defines a single member within a list.
type Member struct {
	ID              string                 `json:"id"`
	EmailAddress    string                 `json:"email_address"`
	UniqueEmailID   string                 `json:"unique_email_id"`
	EmailType       EmailType              `json:"email_type,omitempty"`
	Status          Status                 `json:"status"`
	MergeFields     map[string]interface{} `json:"merge_fields,omitempty"`
	Interests       map[string]bool        `json:"interests,omitempty"`
	Stats           *Stats                 `json:"stats,omitempty"`
	IPSignup        string                 `json:"ip_signup,omitempty"`
	TimestampSignup time.Time              `json:"timestamp_signup,omitempty"`
	IPOpt           string                 `json:"ip_opt,omitempty"`
	TimestampOpt    time.Time              `json:"timestamp_opt,omitempty"`
	MemberRating    uint8                  `json:"member_rating,omitempty"`
	LastChanged     time.Time              `json:"last_changed,omitempty"`
	Language        string                 `json:"language,omitempty"`
	VIP             bool                   `json:"vip,omitempty"`
	EmailClient     string                 `json:"email_client,omitempty"`
	Location        *Location              `json:"location,omitempty"`
	LastNote        *Note                  `json:"last_note,omitempty"`
	ListID          string                 `json:"list_id"`
}

// UnmarshalJSON handles custom JSON unmarshalling for the Member object.
// Credit to http://choly.ca/post/go-json-marshalling/
func (m *Member) UnmarshalJSON(data []byte) error {
	var err error
	type alias Member

	aux := &struct {
		*alias
		TimestampSignup string `json:"timestamp_signup,omitempty"`
		TimestampOpt    string `json:"timestamp_opt,omitempty"`
		LastChanged     string `json:"last_changed,omitempty"`
	}{
		alias: (*alias)(m),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	if aux.TimestampSignup != "" {
		if m.TimestampSignup, err = time.Parse(time.RFC3339, aux.TimestampSignup); err != nil {
			return err
		}
	}
	if aux.TimestampOpt != "" {
		if m.TimestampOpt, err = time.Parse(time.RFC3339, aux.TimestampOpt); err != nil {
			return err
		}
	}
	if aux.LastChanged != "" {
		if m.LastChanged, err = time.Parse(time.RFC3339, aux.LastChanged); err != nil {
			return err
		}
	}

	return nil
}

// ListMembers defines a list of members.
type ListMembers struct {
	Members    []Member `json:"members,omitempty"`
	ListID     string   `json:"list_id"`
	TotalItems int      `json:"total_items"`
}

// NewParams defines the available parameters that can be used when
// adding a new list member via the New function.
type NewParams struct {
	EmailType       EmailType              `json:"email_type,omitempty"`
	Status          Status                 `json:"status"`
	MergeFields     map[string]interface{} `json:"merge_fields,omitempty"`
	Interests       map[string]bool        `json:"interests,omitempty"`
	Language        string                 `json:"language,omitempty"`
	VIP             bool                   `json:"vip,omitempty"`
	Location        *Location              `json:"location,omitempty"`
	IPSignup        string                 `json:"ip_signup,omitempty"`
	TimestampSignup time.Time              `json:"timestamp_signup,omitempty"`
	IPOpt           string                 `json:"ip_opt,omitempty"`
	TimestampOpt    time.Time              `json:"timestamp_opt,omitempty"`
	EmailAddress    string                 `json:"email_address"`
}

// MarshalJSON handles custom JSON marshalling for the NewParams object.
// Credit to http://choly.ca/post/go-json-marshalling/
func (np *NewParams) MarshalJSON() ([]byte, error) {
	var timestampSignup string
	var timestampOpt string

	if !np.TimestampSignup.IsZero() {
		timestampSignup = np.TimestampSignup.Format(time.RFC3339)
	}
	if !np.TimestampOpt.IsZero() {
		timestampOpt = np.TimestampOpt.Format(time.RFC3339)
	}

	type alias NewParams
	return json.Marshal(&struct {
		*alias
		TimestampSignup string `json:"timestamp_signup,omitempty"`
		TimestampOpt    string `json:"timestamp_opt,omitempty"`
	}{
		alias:           (*alias)(np),
		TimestampSignup: timestampSignup,
		TimestampOpt:    timestampOpt,
	})
}

// GetParams defines the available parameters that can be used when
// getting a list of members via the Get function.
type GetParams struct {
	Fields             []string  `url:"fields,omitempty"`
	ExcludeFields      []string  `url:"exclude_fields,omitempty"`
	Count              int       `url:"count,omitempty"`
	Offset             int       `url:"offset,omitempty"`
	EmailType          EmailType `url:"email_type,omitempty"`
	Status             Status    `url:"status,omitempty"`
	SinceTimestampOpt  time.Time `url:"since_timestamp_opt,omitempty"`
	BeforeTimestampOpt time.Time `url:"before_timestamp_opt,omitempty"`
	SinceLastChanged   time.Time `url:"since_last_changed,omitempty"`
	BeforeLastChanged  time.Time `url:"before_last_changed,omitempty"`
	UniqueEmailID      string    `url:"unique_email_id,omitempty"`
	VIPOnly            bool      `url:"vip_only,omitempty"`
}

// EncodeQueryString handles custom query string encoding for the
// GetParams object.
func (gp *GetParams) EncodeQueryString(v interface{}) (string, error) {
	var sinceTimestampOpt string
	var beforeTimestampOpt string
	var sinceLastChanged string
	var beforeLastChanged string

	if !gp.SinceTimestampOpt.IsZero() {
		sinceTimestampOpt = gp.SinceTimestampOpt.Format(time.RFC3339)
	}
	if !gp.BeforeTimestampOpt.IsZero() {
		beforeTimestampOpt = gp.BeforeTimestampOpt.Format(time.RFC3339)
	}
	if !gp.SinceLastChanged.IsZero() {
		sinceLastChanged = gp.SinceLastChanged.Format(time.RFC3339)
	}
	if !gp.BeforeLastChanged.IsZero() {
		beforeLastChanged = gp.BeforeLastChanged.Format(time.RFC3339)
	}

	return query.Encode(struct {
		Fields             string    `url:"fields,omitempty"`
		ExcludeFields      string    `url:"exclude_fields,omitempty"`
		Count              int       `url:"count,omitempty"`
		Offset             int       `url:"offset,omitempty"`
		EmailType          EmailType `url:"email_type,omitempty"`
		Status             Status    `url:"status,omitempty"`
		SinceTimestampOpt  string    `url:"since_timestamp_opt,omitempty"`
		BeforeTimestampOpt string    `url:"before_timestamp_opt,omitempty"`
		SinceLastChanged   string    `url:"since_last_changed,omitempty"`
		BeforeLastChanged  string    `url:"before_last_changed,omitempty"`
		UniqueEmailID      string    `url:"unique_email_id,omitempty"`
		VIPOnly            bool      `url:"vip_only,omitempty"`
	}{
		Fields:             strings.Join(gp.Fields, ","),
		ExcludeFields:      strings.Join(gp.ExcludeFields, ","),
		Count:              gp.Count,
		Offset:             gp.Offset,
		EmailType:          gp.EmailType,
		Status:             gp.Status,
		SinceTimestampOpt:  sinceTimestampOpt,
		BeforeTimestampOpt: beforeTimestampOpt,
		SinceLastChanged:   sinceLastChanged,
		BeforeLastChanged:  beforeLastChanged,
		UniqueEmailID:      gp.UniqueEmailID,
		VIPOnly:            gp.VIPOnly,
	})
}

// GetMemberParams defines the available parameters that can be used
// when getting information on a specific member via the GetMember
// function.
type GetMemberParams struct {
	Fields        []string `url:"fields,omitempty"`
	ExcludeFields []string `url:"exclude_fields,omitempty"`
}

// EncodeQueryString handles custom query string encoding for the
// GetMemberParams object.
func (gmp *GetMemberParams) EncodeQueryString(v interface{}) (string, error) {
	return query.Encode(struct {
		Fields        string `url:"fields,omitempty"`
		ExcludeFields string `url:"exclude_fields,omitempty"`
	}{
		Fields:        strings.Join(gmp.Fields, ","),
		ExcludeFields: strings.Join(gmp.ExcludeFields, ","),
	})
}

// UpdateParams defines the available parameters that can be used when
// updating a list member via the Update function.
type UpdateParams struct {
	EmailType       EmailType              `json:"email_type,omitempty"`
	Status          Status                 `json:"status,omitempty"`
	MergeFields     map[string]interface{} `json:"merge_fields,omitempty"`
	Interests       map[string]bool        `json:"interests,omitempty"`
	Language        string                 `json:"language,omitempty"`
	VIP             bool                   `json:"vip,omitempty"`
	Location        *Location              `json:"location,omitempty"`
	IPSignup        string                 `json:"ip_signup,omitempty"`
	TimestampSignup time.Time              `json:"timestamp_signup,omitempty"`
	IPOpt           string                 `json:"ip_opt,omitempty"`
	TimestampOpt    time.Time              `json:"timestamp_opt,omitempty"`
	EmailAddress    string                 `json:"email_address,omitempty"`
	StatusIfNew     string                 `json:"status_if_new,omitempty"`
}

// MarshalJSON handles custom JSON marshalling for the UpdateParams object.
// Credit to http://choly.ca/post/go-json-marshalling/
func (up *UpdateParams) MarshalJSON() ([]byte, error) {
	var timestampSignup string
	var timestampOpt string

	if !up.TimestampSignup.IsZero() {
		timestampSignup = up.TimestampSignup.Format(time.RFC3339)
	}
	if !up.TimestampOpt.IsZero() {
		timestampOpt = up.TimestampOpt.Format(time.RFC3339)
	}

	type alias UpdateParams
	return json.Marshal(&struct {
		*alias
		TimestampSignup string `json:"timestamp_signup,omitempty"`
		TimestampOpt    string `json:"timestamp_opt,omitempty"`
	}{
		alias:           (*alias)(up),
		TimestampSignup: timestampSignup,
		TimestampOpt:    timestampOpt,
	})
}

// New adds a new list member.
func New(listID string, params *NewParams) (*Member, error) {
	res := &Member{}
	path := fmt.Sprintf("lists/%s/members", listID)

	if params == nil {
		if err := mailchimp.Call("POST", path, nil, nil, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if err := mailchimp.Call("POST", path, nil, params, res); err != nil {
		return nil, err
	}
	return res, nil
}

// Get retrieves information about members in a list.
func Get(listID string, params *GetParams) (*ListMembers, error) {
	res := &ListMembers{}
	path := fmt.Sprintf("lists/%s/members", listID)

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

// GetMember retrieves information about a specific member within a list.
func GetMember(listID, hash string, params *GetMemberParams) (*Member, error) {
	res := &Member{}
	path := fmt.Sprintf("lists/%s/members/%s", listID, hash)

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

// Update updates a list member.
func Update(listID, hash string, params *UpdateParams) (*Member, error) {
	res := &Member{}
	path := fmt.Sprintf("lists/%s/members/%s", listID, hash)

	if params == nil {
		if err := mailchimp.Call("PUT", path, nil, nil, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if err := mailchimp.Call("PUT", path, nil, params, res); err != nil {
		return nil, err
	}
	return res, nil
}

// Delete deletes a list member.
func Delete(listID, hash string) error {
	path := fmt.Sprintf("lists/%s/members/%s", listID, hash)
	return mailchimp.Call("DELETE", path, nil, nil, nil)
}
