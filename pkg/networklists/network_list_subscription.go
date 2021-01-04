package networklists

import (
	"context"
	"fmt"
	"net/http"
)

// NetworkListSubscription represents a collection of NetworkListSubscription
//
// See: NetworkListSubscription.GetNetworkListSubscription()
// API Docs: // network_lists v2
//
// https://developer.akamai.com/api/cloud_security/network_lists/v2.html

type (
	// NetworkListSubscription  contains operations available on NetworkListSubscription  resource
	// See: // network_lists v2
	//
	// https://developer.akamai.com/api/cloud_security/network_lists/v2.html#getnetworklistsubscription
	NetworkListSubscription interface {
		//GetNetworkListSubscriptions(ctx context.Context, params GetNetworkListSubscriptionsRequest) (*GetNetworkListSubscriptionsResponse, error)
		GetNetworkListSubscription(ctx context.Context, params GetNetworkListSubscriptionRequest) (*GetNetworkListSubscriptionResponse, error)
		UpdateNetworkListSubscription(ctx context.Context, params UpdateNetworkListSubscriptionRequest) (*UpdateNetworkListSubscriptionResponse, error)
		RemoveNetworkListSubscription(ctx context.Context, params RemoveNetworkListSubscriptionRequest) (*RemoveNetworkListSubscriptionResponse, error)
	}

	GetNetworkListSubscriptionRequest struct {
		Recipients []string `json:"-"`
		UniqueIds  []string `json:"-"`
	}

	GetNetworkListSubscriptionResponse struct {
		Links struct {
			Create struct {
				Href   string `json:"href"`
				Method string `json:"method"`
			} `json:"create"`
		} `json:"links"`
		NetworkLists []struct {
			ElementCount int `json:"elementCount"`
			Links        struct {
				ActivateInProduction struct {
					Href   string `json:"href"`
					Method string `json:"method"`
				} `json:"activateInProduction"`
				ActivateInStaging struct {
					Href   string `json:"href"`
					Method string `json:"method"`
				} `json:"activateInStaging"`
				AppendItems struct {
					Href   string `json:"href"`
					Method string `json:"method"`
				} `json:"appendItems"`
				Retrieve struct {
					Href string `json:"href"`
				} `json:"retrieve"`
				StatusInProduction struct {
					Href string `json:"href"`
				} `json:"statusInProduction"`
				StatusInStaging struct {
					Href string `json:"href"`
				} `json:"statusInStaging"`
				Update struct {
					Href   string `json:"href"`
					Method string `json:"method"`
				} `json:"update"`
			} `json:"links"`
			Name               string `json:"name"`
			NetworkListType    string `json:"networkListType"`
			ReadOnly           bool   `json:"readOnly"`
			Shared             bool   `json:"shared"`
			SyncPoint          int    `json:"syncPoint"`
			Type               string `json:"type"`
			UniqueID           string `json:"uniqueId"`
			AccessControlGroup string `json:"accessControlGroup,omitempty"`
			Description        string `json:"description,omitempty"`
		} `json:"networkLists"`
	}

	UpdateNetworkListSubscriptionRequest struct {
		Recipients []string `json:"recipients"`
		UniqueIds  []string `json:"uniqueIds"`
	}

	UpdateNetworkListSubscriptionResponse struct {
		Empty string `json:"-"`
	}

	RemoveNetworkListSubscriptionResponse struct {
		Empty string `json:"-"`
	}

	RemoveNetworkListSubscriptionRequest struct {
		Recipients []string `json:"recipients"`
		UniqueIds  []string `json:"uniqueIds"`
	}

	Recipients struct {
		Recipients string `json:"notificationRecipients"`
	}
)

/*
// Validate validates GetNetworkListSubscriptionRequest
func (v GetNetworkListSubscriptionRequest) Validate() error {
	return validation.Errors{
		"Name": validation.Validate(v.Name, validation.Required),
	}.Filter()
}
*/
/*
// Validate validates UpdateNetworkListSubscriptionRequest
func (v RemoveNetworkListSubscriptionRequest) Validate() error {
	return validation.Errors{
		"Name": validation.Validate(v.Name, validation.Required),
	}.Filter()
}
*/
func (p *networklists) GetNetworkListSubscription(ctx context.Context, params GetNetworkListSubscriptionRequest) (*GetNetworkListSubscriptionResponse, error) {
	/*if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}*/

	logger := p.Log(ctx)
	logger.Debug("GetNetworkListSubscription")

	var rval GetNetworkListSubscriptionResponse

	uri :=
		"/network-list/v2/notifications/subscriptions"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create getnetworklistsubscription request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("getnetworklistsubscription  request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil

}

// Update will update a NetworkListSubscription.
//
// API Docs: // network_lists v2
//
// https://developer.akamai.com/api/cloud_security/network_lists/v2.html#putnetworklistsubscription

func (p *networklists) UpdateNetworkListSubscription(ctx context.Context, params UpdateNetworkListSubscriptionRequest) (*UpdateNetworkListSubscriptionResponse, error) {
	/*if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}
	*/
	logger := p.Log(ctx)
	logger.Debug("UpdateNetworkListSubscription")

	postURL :=
		"/network-list/v2/notifications/subscribe"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, postURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create create NetworkListSubscriptionrequest: %w", err)
	}

	var rval UpdateNetworkListSubscriptionResponse
	resp, err := p.Exec(req, &rval, params)
	if err != nil {
		return nil, fmt.Errorf("remove NetworkListSubscription request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
		return nil, p.Error(resp)
	}

	return &rval, nil
}

// Remove will remove a NetworkListSubscription.
//
// API Docs: // network_lists v2
//
// https://developer.akamai.com/api/cloud_security/network_lists/v2.html#putnetworklistsubscription

func (p *networklists) RemoveNetworkListSubscription(ctx context.Context, params RemoveNetworkListSubscriptionRequest) (*RemoveNetworkListSubscriptionResponse, error) {
	/*if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}
	*/
	logger := p.Log(ctx)
	logger.Debug("UpdateNetworkListSubscription")

	postURL :=
		"/network-list/v2/notifications/unsubscribe"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, postURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create create NetworkListSubscriptionrequest: %w", err)
	}

	var rval RemoveNetworkListSubscriptionResponse
	resp, err := p.Exec(req, &rval, params)
	if err != nil {
		return nil, fmt.Errorf("remove NetworkListSubscription request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
		return nil, p.Error(resp)
	}

	return &rval, nil
}