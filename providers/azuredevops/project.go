package azuredevpos

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
)

type ProjectGenerator struct {
	AzureDevOpsService
}

func (az *ProjectGenerator) listResources() ([]core.TeamProjectReference, error) {

	client, err := az.getCoreClient()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	projects, err := client.GetProjects(ctx, core.GetProjectsArgs{})
	if err != nil {
		return nil, err
	}
	var resources []core.TeamProjectReference
	for projects != nil {
		resources = append(resources, (*projects).Value...)
		if projects.ContinuationToken != "" {
			// Get next page of team projects
			projectArgs := core.GetProjectsArgs{
				ContinuationToken: &projects.ContinuationToken,
			}
			projects, err = client.GetProjects(ctx, projectArgs)
			if err != nil {
				return nil, err
			}
		} else {
			projects = nil
		}
	}
	return resources, nil
}

func (az *ProjectGenerator) appendResource(project *core.TeamProjectReference) {
	az.appendSimpleResource((*project.Id).String(), *project.Name, "azuredevops_project")
}

func (az *ProjectGenerator) InitResources() error {

	projects, err := az.listResources()
	if err != nil {
		return err
	}
	for _, project := range projects {
		az.appendResource(&project)
	}
	return nil
}
