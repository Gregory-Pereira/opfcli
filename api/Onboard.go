package api

import (
	"fmt"
	"github.com/operate-first/opfcli/models"
	"strings"
)

func (api *API) Onboard(path string, projectDisplayName string, disableLimitrange bool) error {
	onboardRequest, err := models.OnboardRequestFromYAMLPath(path)

	fmt.Printf("onboardRequest, %v\n", onboardRequest)
	if err := api.CreateGroup(
		onboardRequest.Teamname,
		true,
	); err != nil {
		return err
	}

	if err := api.CreateRoleBinding(onboardRequest.Teamname, "admin"); err != nil {
		return err
	}

	var tempQuota string
	for i := 0; i < len(onboardRequest.Namespaces); i++ {
		if onboardRequest.Namespaces[i].Quota != "" && onboardRequest.Namespaces[i].Quota != "_No response_" {
			tempQuota = strings.ToLower(onboardRequest.Namespaces[i].Quota)
			if err := api.ValidateQuota(tempQuota); err != nil {
				return err
			}
		}
		api.CreateNamespace(
			onboardRequest.Namespaces[i].Name,
			onboardRequest.Teamname,
			projectDisplayName,
			tempQuota,
			disableLimitrange,
			true,
		)
	}

	if err != nil {
		return err
	}
	return nil

}
