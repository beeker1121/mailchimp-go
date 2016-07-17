package members

import (
	"fmt"

	mailchimp "github.com/beeker1121/mailchimp-go"
)

// EmailType defines the type of email a member asked to get.
type EmailType string

const (
	EmailTypeHTML EmailType = "html"
	EmailTypeText           = "text"
)

// Status defines the subscription status for a given member
// within a List.
type Status string

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
	NoteID    int    `json:"note_id"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	Note      string `json:"note"`
}

// Member defines a single member within a list.
//
// Schema: https://api.mailchimp.com/schema/3.0/Lists/Members/Instance.json
type Member struct {
	ID              string                 `json:"id"`
	EmailAddress    string                 `json:"email_address"`
	UniqueEmailID   string                 `json:"unique_email_id"`
	EmailType       EmailType              `json:"email_type,omitempty"`
	Status          Status                 `json:"status"`
	MergeFields     map[string]interface{} `json:"merge_fields,omitempty"`
	Interests       map[string]bool        `json:"interests,omitempty"`
	Stats           *Stats                 `json:"stats,omitempty"`
	IPSignup        string                 `json:ip_signup,omitempty"`
	TimestampSignup string                 `json:"timestamp_signup,omitempty"`
	IPOpt           string                 `json:"ip_opt,omitempty"`
	TimestampOpt    string                 `json:"timestamp_opt,omitempty"`
	MemberRating    uint8                  `json:"member_rating,omitempty"`
	LastChanged     string                 `json:"last_changed,omitempty"`
	Language        string                 `json:"language,omitempty"`
	VIP             bool                   `json:"vip,omitempty"`
	EmailClient     string                 `json:"email_client,omitempty"`
	Location        *Location              `json:"location,omitempty"`
	LastNote        *Note                  `json:"last_note,omitempty"`
	ListID          string                 `json:"list_id"`
}

// ListMembers defines a list of members.
//
// Schema: https://api.mailchimp.com/schema/3.0/Lists/Members/Collection.json
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
	TimestampSignup string                 `json:"timestamp_signup,omitempty"`
	IPOpt           string                 `json:"ip_opt,omitempty"`
	TimestampOpt    string                 `json:"timestamp_opt,omitempty"`
	EmailAddress    string                 `json:"email_address"`
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
	SinceTimestampOpt  string    `url:"since_timestamp_opt,omitempty"`
	BeforeTimestampOpt string    `url:"before_timestamp_opt,omitempty"`
	SinceLastChanged   string    `url:"since_last_changed,omitempty"`
	BeforeLastChanged  string    `url:"before_last_changed,omitempty"`
	UniqueEmailID      string    `url:"unique_email_id,omitempty"`
	VIPOnly            bool      `url:"vip_only,omitempty"`
}

// New adds a new list member.
//
// Method:   POST
// Resource: lists/{list_id}/members
//
// Definition: http://developer.mailchimp.com/documentation/mailchimp/reference/lists/members/#create-post_lists_list_id_members
func New(listID string, params *NewParams) (*Member, error) {
	res := &Member{}

	path := fmt.Sprintf("lists/%s/members", listID)

	if err := mailchimp.Call("POST", path, nil, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Get retrieves information about members in a list.
//
// Method: GET
// Resource: /lists/{list_id}/members
//
// Definition: http://developer.mailchimp.com/documentation/mailchimp/reference/lists/members/#read-get_lists_list_id_members
//func Get(listID string, params *GetParams) (*ListMembers, error) {}

// GetMember retrieves information about a specific member within a list.
//
// Method: GET
// Resource: /lists/{list_id}/members/{subscriber_hash}
//
// Definition: http://developer.mailchimp.com/documentation/mailchimp/reference/lists/members/#read-get_lists_list_id_members_subscriber_hash
//func GetMember(listID, hash string, params *GetMemberParams) (*Member, error) {}
