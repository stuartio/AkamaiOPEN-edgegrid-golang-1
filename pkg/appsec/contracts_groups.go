package appsec

import (
	"context"
	"fmt"
	"net/http"
)

// ContractsGroups represents a collection of ContractsGroups
//
// See: ContractsGroups.GetContractsGroups()
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html

type (
	// ContractsGroups  contains operations available on ContractsGroups  resource
	// See: // appsec v1
	//
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getcontractsgroups
	ContractsGroups interface {
		GetContractsGroups(ctx context.Context, params GetContractsGroupsRequest) (*GetContractsGroupsResponse, error)
	}

	GetContractsGroupsRequest struct {
		ConfigID   int    `json:"-"`
		Version    int    `json:"-"`
		PolicyID   string `json:"-"`
		ContractID string `json:"-"`
		GroupID    int    `json:"-"`
	}

	GetContractsGroupsResponse struct {
		ContractGroups []struct {
			ContractID  string `json:"contractId"`
			DisplayName string `json:"displayName"`
			GroupID     int    `json:"groupId"`
		} `json:"contract_groups"`
	}
)

func (p *appsec) GetContractsGroups(ctx context.Context, params GetContractsGroupsRequest) (*GetContractsGroupsResponse, error) {

	logger := p.Log(ctx)
	logger.Debug("GetContractsGroups")

	var rval GetContractsGroupsResponse
	var rvalfiltered GetContractsGroupsResponse

	uri :=
		"/appsec/v1/contracts-groups"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create getcontractsgroups request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("getcontractsgroups  request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	if params.GroupID != 0 {
		for _, val := range rval.ContractGroups {
			if val.ContractID == params.ContractID && val.GroupID == params.GroupID {
				rvalfiltered.ContractGroups = append(rvalfiltered.ContractGroups, val)
			}
		}
	} else {
		rvalfiltered = rval
	}
	return &rvalfiltered, nil

}
