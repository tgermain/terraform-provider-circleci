package circleci

import (
	"log"

	circleciapi "github.com/jszwedko/go-circleci"
)

// ProviderClient is a thin commodity wrapper on top of circleciapi
type ProviderClient struct {
	client       *circleciapi.Client
	vcsType      string
	organization string
}

// NewConfig initialize circleci API client and returns a new config object
func NewConfig(token, vscType, organization string) *ProviderClient {
	log.Printf("NewConfig with vscType '%s' , orga '%s'", vscType, organization)
	return &ProviderClient{
		client: &circleciapi.Client{
			Token: token,
		},
		vcsType:      vscType,
		organization: organization,
	}
}

// GetEnvVar get the environment variable with given name
// It returns an empty structure if no environment variable exists with that name
func (pv *ProviderClient) GetEnvVar(projectName, envVarName string) (*circleciapi.EnvVar, error) {
	log.Printf("GetEnvVar with vscType '%s' , orga '%s', projectName '%s', envvarname '%s'", pv.vcsType, pv.organization, projectName, envVarName)
	return pv.client.GetEnvVar(pv.vcsType, pv.organization, projectName, envVarName)
}

// EnvVarExists check if environment variable exists with given name
func (pv *ProviderClient) EnvVarExists(projectName, envVarName string) (bool, error) {
	log.Printf("EnvVarExists with vscType '%s'  orga '%s' projectName '%s'  envvarname '%s'", pv.vcsType, pv.organization, projectName, envVarName)
	envVar, err := pv.client.GetEnvVar(pv.vcsType, pv.organization, projectName, envVarName)
	if err != nil {
		log.Printf("Got error '%s'", err.Error())
		return false, err
	}
	return bool(envVar.Name != ""), nil
}

// AddEnvVar create an environment variable with given name and value
func (pv *ProviderClient) AddEnvVar(projectName, envVarName, envVarValue string) (*circleciapi.EnvVar, error) {
	log.Printf("AddEnvVar with vscType '%s' , orga '%s', projectName '%s', envvarname '%s'", pv.vcsType, pv.organization, projectName, envVarName)
	return pv.client.AddEnvVar(pv.vcsType, pv.organization, projectName, envVarName, envVarValue)
}

// DeleteEnvVar delete the environment variable with given name
func (pv *ProviderClient) DeleteEnvVar(projectName, envVarName string) error {
	log.Printf("AddEnvVar with vscType '%s' , orga '%s', projectName '%s', envvarname '%s'", pv.vcsType, pv.organization, projectName, envVarName)
	return pv.client.DeleteEnvVar(pv.vcsType, pv.organization, projectName, envVarName)
}
