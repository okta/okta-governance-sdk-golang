package governance

import (
	"fmt"

	"github.com/okta/okta-sdk-golang/v5/okta"
)

var (
	apiClient      *OktaGovernanceAPIClient
	idaasAPIClient *okta.APIClient
)

func init() {
	configuration, err := okta.NewConfiguration(okta.WithCache(false))
	if err != nil {
		fmt.Printf("Create new config should not be error %v", err)
	}
	configuration.Debug = false

	apiClient = NewAPIClient(configuration)
	idaasAPIClient = okta.NewAPIClient(configuration)
}
