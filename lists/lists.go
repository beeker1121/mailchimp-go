package lists

// Contact defines the contact information of the list owner, which
// is displayed in footer of campaigns.
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

// Stats defines statistics for a list.
type Stats struct {
	MemberCount               uint    `json:"member_count,omitempty"`
	UnsubscribeCount          uint    `json:"unsubscribe_count,omitempty"`
	CleanedCount              uint    `json:"cleaned_count,omitempty"`
	MemberCountSinceSend      uint    `json:"member_count_since_send,omitempty"`
	UnsubscribeCountSinceSend uint    `json:"unsubscribe_count_since_send,omitempty"`
	CleanedCountSinceSend     uint    `json:"cleaned_count_since_send,omitempty"`
	CampaignCount             uint    `json:"campaign_count,omitempty"`
	CampaignLastSent          string  `json:"campaign_last_sent,omitempty"`
	MergeFieldCount           uint    `json:"merge_field_count,omitempty"`
	AvgSubRate                float64 `json:"avg_sub_rate,omitempty"`
	AvgUnsubRate              float64 `json:"avg_unsub_rate,omitempty"`
	TargetSubRate             float64 `json:"target_sub_rate,omitempty"`
	OpenRate                  float32 `json:"open_rate,omitempty"`
	ClickRate                 float32 `json:"click_rate,omitempty"`
	LastSubDate               string  `json:"last_sub_date,omitempty"`
	LastUnsubDate             string  `json:"last_unsub_date,omitempty"`
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
	DateCreated         string            `json:"date_created,omitempty"`
	ListRating          uint8             `json:"list_rating,omitempty"`
	EmailTypeOption     bool              `json:"email_type_option"`
	SubscribeUrlShort   string            `json:"subscribe_url_short,omitempty"`
	SubscribeUrlLong    string            `json:"subscribe_url_long,omitempty"`
	BeamerAddress       string            `json:"beamer_address,omitempty"`
	Visibility          string            `json:"visibility"`
	Modules             []string          `json:"modules,omitempty"`
	Stats               *Stats            `json:"stats,omitempty"`
}
