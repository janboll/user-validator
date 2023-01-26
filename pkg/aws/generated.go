// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package aws

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/app-sre/go-qontract-reconcile/pkg/gql"
)

// GetAccountsAccountsAWSAccount_v1 includes the requested fields of the GraphQL type AWSAccount_v1.
type GetAccountsAccountsAWSAccount_v1 struct {
	Name                   string                                                        `json:"name"`
	ResourcesDefaultRegion string                                                        `json:"resourcesDefaultRegion"`
	AutomationToken        GetAccountsAccountsAWSAccount_v1AutomationTokenVaultSecret_v1 `json:"automationToken"`
}

// GetName returns GetAccountsAccountsAWSAccount_v1.Name, and is useful for accessing the field via an interface.
func (v *GetAccountsAccountsAWSAccount_v1) GetName() string { return v.Name }

// GetResourcesDefaultRegion returns GetAccountsAccountsAWSAccount_v1.ResourcesDefaultRegion, and is useful for accessing the field via an interface.
func (v *GetAccountsAccountsAWSAccount_v1) GetResourcesDefaultRegion() string {
	return v.ResourcesDefaultRegion
}

// GetAutomationToken returns GetAccountsAccountsAWSAccount_v1.AutomationToken, and is useful for accessing the field via an interface.
func (v *GetAccountsAccountsAWSAccount_v1) GetAutomationToken() GetAccountsAccountsAWSAccount_v1AutomationTokenVaultSecret_v1 {
	return v.AutomationToken
}

// GetAccountsAccountsAWSAccount_v1AutomationTokenVaultSecret_v1 includes the requested fields of the GraphQL type VaultSecret_v1.
type GetAccountsAccountsAWSAccount_v1AutomationTokenVaultSecret_v1 struct {
	Path    string `json:"path"`
	Field   string `json:"field"`
	Version int    `json:"version"`
	Format  string `json:"format"`
}

// GetPath returns GetAccountsAccountsAWSAccount_v1AutomationTokenVaultSecret_v1.Path, and is useful for accessing the field via an interface.
func (v *GetAccountsAccountsAWSAccount_v1AutomationTokenVaultSecret_v1) GetPath() string {
	return v.Path
}

// GetField returns GetAccountsAccountsAWSAccount_v1AutomationTokenVaultSecret_v1.Field, and is useful for accessing the field via an interface.
func (v *GetAccountsAccountsAWSAccount_v1AutomationTokenVaultSecret_v1) GetField() string {
	return v.Field
}

// GetVersion returns GetAccountsAccountsAWSAccount_v1AutomationTokenVaultSecret_v1.Version, and is useful for accessing the field via an interface.
func (v *GetAccountsAccountsAWSAccount_v1AutomationTokenVaultSecret_v1) GetVersion() int {
	return v.Version
}

// GetFormat returns GetAccountsAccountsAWSAccount_v1AutomationTokenVaultSecret_v1.Format, and is useful for accessing the field via an interface.
func (v *GetAccountsAccountsAWSAccount_v1AutomationTokenVaultSecret_v1) GetFormat() string {
	return v.Format
}

// GetAccountsResponse is returned by GetAccounts on success.
type GetAccountsResponse struct {
	Accounts []GetAccountsAccountsAWSAccount_v1 `json:"accounts"`
}

// GetAccounts returns GetAccountsResponse.Accounts, and is useful for accessing the field via an interface.
func (v *GetAccountsResponse) GetAccounts() []GetAccountsAccountsAWSAccount_v1 { return v.Accounts }

// __GetAccountsInput is used internally by genqlient
type __GetAccountsInput struct {
	AccountName string `json:"AccountName"`
}

// GetAccountName returns __GetAccountsInput.AccountName, and is useful for accessing the field via an interface.
func (v *__GetAccountsInput) GetAccountName() string { return v.AccountName }

func GetAccounts(
	ctx context.Context,
	AccountName string,
) (*GetAccountsResponse, error) {
	req := &graphql.Request{
		OpName: "GetAccounts",
		Query: `
query GetAccounts ($AccountName: String!) {
	accounts: awsaccounts_v1(name: $AccountName) {
		name
		resourcesDefaultRegion
		automationToken {
			path
			field
			version
			format
		}
	}
}
`,
		Variables: &__GetAccountsInput{
			AccountName: AccountName,
		},
	}
	var err error
	var client graphql.Client

	client, err = gql.NewQontractClient(ctx)
	if err != nil {
		return nil, err
	}

	var data GetAccountsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
