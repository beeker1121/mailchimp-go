# mailchimp-go [![License](http://img.shields.io/badge/license-mit_bsd-blue.svg)](https://raw.githubusercontent.com/beeker1121/mailchimp-go/master/LICENSE) [![Go Report Card](https://img.shields.io/badge/go_report-A+-brightgreen.svg)](https://goreportcard.com/report/github.com/beeker1121/mailchimp-go) [![Build Status](https://travis-ci.org/beeker1121/mailchimp-go.svg?branch=master)](https://travis-ci.org/beeker1121/mailchimp-go)

mailchimp-go is a Go client for the MailChimp API v3.

While coverage of the MailChimp API is limited in the current state, the goal is to provide a basic structure that can be built upon to eventually have full coverage.

Contributing code to complete missing resources is greatly appreciated.

## API GoDoc Reference

Below is the main mailchimp-go GoDoc reference:

**mailchimp-go** - [http://godoc.org/github.com/beeker1121/mailchimp-go](http://godoc.org/github.com/beeker1121/mailchimp-go)

Each API resource is a separate package within mailchimp-go.

Below are the GoDoc references for each supported resource:

**Lists** - [https://godoc.org/github.com/beeker1121/mailchimp-go/lists](https://godoc.org/github.com/beeker1121/mailchimp-go/lists)  
**Lists/Members** - [https://godoc.org/github.com/beeker1121/mailchimp-go/lists/members](https://godoc.org/github.com/beeker1121/mailchimp-go/lists/members)

## Installation

Fetch the package from GitHub:

```sh
go get github.com/beeker1121/mailchimp-go
```

Import to your project:

```go
import mailchimp "github.com/beeker1121/mailchimp-go"
```

Import the API resources you wish to use. For example, to use the `Lists` resource:

```go
import "github.com/beeker1121/mailchimp-go/lists"
```

## Usage

At the moment, this library has minimal coverage of the MailChimp API.

### Set API Key

First, set your MailChimp API key:

```go
import mailchimp "github.com/beeker1121/mailchimp-go"
...
err := mailchimp.SetKey("YOUR-API-KEY")
...
```

### Create a list

```go
import "github.com/beeker1121/mailchimp-go/lists"
...

// Set request parameters.
params := &lists.NewParams{
	Name: "My List",
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

list, err := lists.New(params)
...
fmt.Printf("%+v\n", list)
```

### Add a member to a list

```go
import "github.com/beeker1121/mailchimp-go/lists/members"
...

// Set request parameters.
params := &members.NewParams{
	EmailAddress: "user@example.com",
	Status:       members.StatusSubscribed,
}

// Add member to list 123456.
member, err := members.New("123456", params)
...
fmt.Printf("%+v\n", member)
```

### Get list members

```go
import "github.com/beeker1121/mailchimp-go/lists/members"
...

// Set request parameters.
params := &members.GetParams{
	Status: members.StatusSubscribed,
}

// Get subscribed members of list 123456.
listMembers, err := members.Get("123456", params)
...
fmt.Printf("%+v\n", listMembers)
```

### Get a list member

```go
import "github.com/beeker1121/mailchimp-go/lists/members"
...

// Get member 123 from list 123456.
member, err := members.GetMember("123456", "123", nil)
...
fmt.Printf("%+v\n", member)
```

### Delete a list member

```go
import "github.com/beeker1121/mailchimp-go/lists/members"
...

// Delete member 123 from list 123456.
err := members.Delete("123456", "123")
...
```

## Testing

To run the tests, you must have a valid MailChimp account and API key.

Set the API key environment variable:

```sh
export MAILCHIMP_API_KEY=your-key
```

Run the tests from the mailchimp-go directory:

```sh
go test -v ./...
```

## Thanks

**The Go Devs** - For existing  
**Discord Gophers** ([Join Chat](https://discordapp.com/channels/118456055842734083/118456055842734083)) - Troubleshooting and advice  
**go-querystring** ([https://github.com/google/go-querystring](https://github.com/google/go-querystring)) - Query string library  
**Choly** ([http://choly.ca/](http://choly.ca/)) - custom JSON solution