package mailchimp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/beeker1121/mailchimp-go/query"
)

// The MailChimp API url structure.
const apiURL string = "https://%s.api.mailchimp.com/3.0/%s"

// key is the MailChimp API key.
var key string

// dc is the data center used for the given API key.
var dc string

// httpClient is the default http.Client used to make requests.
var httpClient = &http.Client{}

// SetKey sets the API key and updates the data center value accordingly.
func SetKey(apiKey string) error {
	// Get the data center from the key.
	split := strings.Split(apiKey, "-")
	if len(split) != 2 {
		return ErrAPIKeyFormat
	}

	// Set key and dc values.
	key = apiKey
	dc = split[1]

	return nil
}

// SetClient sets the http.Client used to make API requests.
func SetClient(client *http.Client) {
	httpClient = client
}

// Call issues a request to the MailChimp API.
func Call(method, path string, queryParams, bodyParams, v interface{}) error {
	// Check if the API key has been set.
	if key == "" {
		return ErrAPIKeyNotSet
	}

	// Build the API url using the given endpoint.
	u := fmt.Sprintf(apiURL, dc, path)

	// Handle the query parameters.
	if queryParams != nil {
		q, err := query.Encode(queryParams)
		if err != nil {
			return err
		}

		u += "?" + q
	}

	// Handle the body parameters.
	b := new(bytes.Buffer)
	if bodyParams != nil {
		if err := json.NewEncoder(b).Encode(bodyParams); err != nil {
			return err
		}
	}

	// Build the Request.
	req, err := http.NewRequest(method, u, b)
	if err != nil {
		return err
	}

	// Set headers.
	req.SetBasicAuth("", key)

	// Send request.
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle API error.
	if resp.StatusCode >= 400 {
		apiErr := new(APIError)
		if err = json.NewDecoder(resp.Body).Decode(apiErr); err != nil {
			return err
		}

		return apiErr
	}

	if v != nil {
		return json.NewDecoder(resp.Body).Decode(v)
	}

	return nil
}
