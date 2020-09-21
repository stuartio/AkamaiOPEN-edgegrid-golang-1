package papi

import (
	"context"
	"fmt"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v2/pkg/session"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/spf13/cast"
	"net/http"
)

type (
	// PropertyRules contains operations available on PropertyRule resource
	// See: https://developer.akamai.com/api/core_features/property_manager/v1.html#propertyversionrulesgroup
	PropertyRules interface {
		// GetRuleTree lists all available CP codes
		// See: https://developer.akamai.com/api/core_features/property_manager/v1.html#getpropertyversionrules
		GetRuleTree(context.Context, GetRuleTreeRequest) (*GetRuleTreeResponse, error)

		// UpdateRuleTree lists all available CP codes
		// See: https://developer.akamai.com/api/core_features/property_manager/v1.html#putpropertyversionrules
		UpdateRuleTree(context.Context, UpdateRulesRequest) (*UpdateRulesResponse, error)
	}

	// GetRuleTreeRequest contains path and query params necessary to perform GET /rules request
	GetRuleTreeRequest struct {
		PropertyID      string
		PropertyVersion int
		ContractID      string
		GroupID         string
		ValidateMode    string
		ValidateRules   bool
	}

	// GetRuleTreeResponse contains data returned by performing GET /rules request
	GetRuleTreeResponse struct {
		AccountID       string `json:"accountId"`
		ContractID      string `json:"contractId"`
		GroupID         string `json:"groupId"`
		PropertyID      string `json:"propertyId"`
		PropertyVersion int    `json:"propertyVersion"`
		Etag            string `json:"etag"`
		RuleFormat      string `json:"ruleFormat"`
		Rules           Rules  `json:"rules"`
	}

	// Rules contains Rule object
	Rules struct {
		AdvancedOverride string              `json:"advancedOverride"`
		Behaviors        []RuleBehavior      `json:"behaviors"`
		Children         []Rules             `json:"children"`
		Comment          string              `json:"comment"`
		Criteria         []RuleBehavior      `json:"criteria"`
		CriteriaLocked   bool                `json:"criteriaLocked"`
		CustomOverride   *RuleCustomOverride `json:"customOverride"`
		Name             string              `json:"name"`
		Options          *RuleOptions        `json:"options"`
		UUID             string              `json:"uuid"`
		Variables        []RuleVariable      `json:"variables"`
	}

	// RuleBehavior contains data for both rule behaviors and rule criteria
	RuleBehavior struct {
		Locked  string `json:"locked"`
		Name    string `json:"name"`
		Options map[string]interface{}
		UUID    string `json:"uuid"`
	}

	// RuleCustomOverride represents customOverride field from Rule resource
	RuleCustomOverride struct {
		Name       string `json:"name"`
		OverrideID string `json:"overrideId"`
	}

	// RuleOptions represents options field from Rule resource
	RuleOptions struct {
		IsSecure bool `json:"is_secure"`
	}

	// RuleVariable represents and entry in variables field from Rule resource
	RuleVariable struct {
		Description string `json:"description"`
		Hidden      bool   `json:"hidden"`
		Name        string `json:"name"`
		Sensitive   bool   `json:"sensitive"`
		Value       string `json:"value"`
	}

	// UpdateRulesRequest contains path and query params, as well as request body necessary to perform PUT /rules request
	UpdateRulesRequest struct {
		PropertyID      string
		PropertyVersion int
		ContractID      string
		DryRun          bool
		GroupID         string
		ValidateMode    string
		ValidateRules   bool
		Rules           Rules
	}
	// UpdateRulesResponse contains data returned by performing PUT /rules request
	UpdateRulesResponse struct {
		AccountID       string      `json:"accountId"`
		ContractID      string      `json:"contractId"`
		GroupID         string      `json:"groupId"`
		PropertyID      string      `json:"propertyId"`
		PropertyVersion int         `json:"propertyVersion"`
		Etag            string      `json:"etag"`
		RuleFormat      string      `json:"ruleFormat"`
		Rules           Rules       `json:"rules"`
		Errors          []RuleError `json:"errors"`
	}

	// RuleError represents and entry in error field from PUT /rules response body
	RuleError struct {
		Type         string `json:"type"`
		Title        string `json:"title"`
		Detail       string `json:"detail"`
		Instance     string `json:"instance"`
		BehaviorName string `json:"behaviorName"`
	}
)

const (
	// RuleValidateModeFast const
	RuleValidateModeFast = "fast"
	// RuleValidateModeFull const
	RuleValidateModeFull = "full"
)

// Validate validates GetRuleTreeRequest struct
func (r GetRuleTreeRequest) Validate() error {
	return validation.Errors{
		"PropertyID":      validation.Validate(r.PropertyID, validation.Required),
		"PropertyVersion": validation.Validate(r.PropertyVersion, validation.Required),
		"ValidateMode":    validation.Validate(r.ValidateMode, validation.In(RuleValidateModeFast, RuleValidateModeFull)),
	}.Filter()
}

// Validate validates UpdateRulesRequest struct
func (r UpdateRulesRequest) Validate() error {
	return validation.Errors{
		"PropertyID":      validation.Validate(r.PropertyID, validation.Required),
		"PropertyVersion": validation.Validate(r.PropertyVersion, validation.Required),
		"ValidateMode":    validation.Validate(r.ValidateMode, validation.In(RuleValidateModeFast, RuleValidateModeFull)),
		"Rules":           validation.Validate(r.Rules),
	}.Filter()
}

// Validate validates Rules struct
func (r Rules) Validate() error {
	return validation.Errors{
		"Behaviors":      validation.Validate(r.Behaviors, validation.Required),
		"Name":           validation.Validate(r.Name, validation.Required),
		"CustomOverride": validation.Validate(r.CustomOverride),
		"Criteria":       validation.Validate(r.Criteria),
		"Children":       validation.Validate(r.Children),
		"Variables":      validation.Validate(r.Variables),
	}.Filter()
}

// Validate validates RuleBehavior struct
func (b RuleBehavior) Validate() error {
	return validation.Errors{
		"Name":    validation.Validate(b.Name, validation.Required),
		"Options": validation.Validate(b.Options, validation.Required),
	}.Filter()
}

// Validate validates RuleCustomOverride struct
func (co RuleCustomOverride) Validate() error {
	return validation.Errors{
		"Name":       validation.Validate(co.Name, validation.Required),
		"OverrideID": validation.Validate(co.OverrideID, validation.Required),
	}.Filter()
}

// Validate validates RuleVariable struct
func (v RuleVariable) Validate() error {
	return validation.Errors{
		"Name": validation.Validate(v.Name, validation.Required),
	}.Filter()
}

func (p *papi) GetRuleTree(ctx context.Context, params GetRuleTreeRequest) (*GetRuleTreeResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("GetRuleTree")

	getURL := fmt.Sprintf(
		"/papi/v1/properties/%s/versions/%d/rules?contractId=%s&groupId=%s",
		params.PropertyID,
		params.PropertyVersion,
		params.ContractID,
		params.GroupID,
	)
	if params.ValidateMode != "" {
		getURL += fmt.Sprintf("&validateMode=%s", params.ValidateMode)
	}
	if !params.ValidateRules {
		getURL += fmt.Sprintf("&validateRules=%t", params.ValidateRules)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, getURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create getruletree request: %w", err)
	}

	req.Header.Set("PAPI-Use-Prefixes", cast.ToString(p.usePrefixes))
	var versions GetRuleTreeResponse
	resp, err := p.Exec(req, &versions)
	if err != nil {
		return nil, fmt.Errorf("getruletree request failed: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("%w: %s", session.ErrNotFound, getURL)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, session.NewAPIError(resp, logger)
	}

	return &versions, nil
}

func (p *papi) UpdateRuleTree(ctx context.Context, request UpdateRulesRequest) (*UpdateRulesResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("UpdateRuleTree")

	putURL := fmt.Sprintf(
		"/papi/v1/properties/%s/versions/%d/rules?contractId=%s&groupId=%s",
		request.PropertyID,
		request.PropertyVersion,
		request.ContractID,
		request.GroupID,
	)
	if request.ValidateMode != "" {
		putURL += fmt.Sprintf("&validateMode=%s", request.ValidateMode)
	}
	if !request.ValidateRules {
		putURL += fmt.Sprintf("&validateRules=%t", request.ValidateRules)
	}
	if request.DryRun {
		putURL += fmt.Sprintf("&dryRun=%t", request.DryRun)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, putURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create UpdateRuleTree request: %w", err)
	}

	req.Header.Set("PAPI-Use-Prefixes", cast.ToString(p.usePrefixes))
	var versions UpdateRulesResponse
	resp, err := p.Exec(req, &versions, request.Rules, request.Rules)
	if err != nil {
		return nil, fmt.Errorf("UpdateRuleTree request failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, session.NewAPIError(resp, logger)
	}

	return &versions, nil
}