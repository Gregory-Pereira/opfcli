package models

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type OnboardNamespace struct {
	Enablemonitoring bool
	Name             string
	Quota            string
}

type OnboardingRequest struct {
	Env                string `yaml:",omitempty"`
	Namespaces         []OnboardNamespace
	ProjectDescription string `yaml:"project_description"`
	Targetcluster      string `yaml:"target_cluster"`
	Teamname           string `yaml:"team_name"`
	Template           string
	Users              []string `yaml:",omitempty"`
}

func NewOnboardingNamespace(projectName string, namespaceQuota string, customQuota []string) OnboardNamespace {
	rscs := OnboardNamespace{
		Enablemonitoring: true,
		Name:             projectName,
		Quota:            namespaceQuota,
	}
	return rscs
}

func OnboardRequestFromYAMLPath(path string) (OnboardingRequest, error) {

	var onboard_request OnboardingRequest

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return OnboardingRequest{}, err
	}

	err = yaml.Unmarshal(content, &onboard_request)
	if err != nil {
		return OnboardingRequest{}, err
	}

	return onboard_request, nil
}
