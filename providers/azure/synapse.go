package azure

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/preview/synapse/2019-06-01-preview/managedvirtualnetwork"
	"github.com/Azure/azure-sdk-for-go/services/synapse/mgmt/2020-12-01/synapse"
	// "github.com/Azure/azure-sdk-for-go/services/preview/synapse/2020-08-01-preview/accesscontrol"
)

type SynapseGenerator struct {
	AzureService
}

func (az *SynapseGenerator) listWorkspaces() ([]synapse.Workspace, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := synapse.NewWorkspacesClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator synapse.WorkspaceInfoListResultIterator
		err      error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		iterator, err = client.ListByResourceGroupComplete(ctx, resourceGroup)
	} else {
		iterator, err = client.ListComplete(ctx)
	}
	if err != nil {
		return nil, err
	}
	var resources []synapse.Workspace
	for iterator.NotDone() {
		item := iterator.Value()
		resources = append(resources, item)
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return resources, err
		}
	}
	return resources, nil
}

func (az *SynapseGenerator) appendWorkspace(workspace *synapse.Workspace) {
	az.AppendSimpleResource(*workspace.ID, *workspace.Name, "azurerm_synapse_workspace", "syn")
}

func (az *SynapseGenerator) appendSQLPools(workspace *synapse.Workspace, workspaceRg *ResourceID) error {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := synapse.NewSQLPoolsClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	iterator, err := client.ListByWorkspaceComplete(ctx, workspaceRg.ResourceGroup, *workspace.Name)
	if err != nil {
		return err
	}
	for iterator.NotDone() {
		item := iterator.Value()
		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_synapse_sql_pool", "synp")
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (az *SynapseGenerator) appendSparkPools(workspace *synapse.Workspace, workspaceRg *ResourceID) error {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := synapse.NewBigDataPoolsClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	iterator, err := client.ListByWorkspaceComplete(ctx, workspaceRg.ResourceGroup, *workspace.Name)
	if err != nil {
		return err
	}
	for iterator.NotDone() {
		item := iterator.Value()
		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_synapse_spark_pool", "synp")
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (az *SynapseGenerator) appendFirewallRule(workspace *synapse.Workspace, workspaceRg *ResourceID) error {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := synapse.NewIPFirewallRulesClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	iterator, err := client.ListByWorkspaceComplete(ctx, workspaceRg.ResourceGroup, *workspace.Name)
	if err != nil {
		return err
	}
	for iterator.NotDone() {
		item := iterator.Value()
		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_synapse_firewall_rule", "synf")
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (az *SynapseGenerator) appendManagedPrivateEndpoint(workspace *synapse.Workspace) error {

	if workspace.WorkspaceProperties == nil || workspace.WorkspaceProperties.ManagedVirtualNetwork == nil {
		return nil
	}
	virtualNetworkName := *workspace.WorkspaceProperties.ManagedVirtualNetwork
	subscriptionID, _, authorizer := az.getClientArgs()
	client := managedvirtualnetwork.NewManagedPrivateEndpointsClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	iterator, err := client.ListComplete(ctx, virtualNetworkName)
	if err != nil {
		return err
	}
	for iterator.NotDone() {
		item := iterator.Value()
		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_synapse_managed_private_endpoint", "syne")
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (az *SynapseGenerator) listPrivateLinkHubs() ([]synapse.PrivateLinkHub, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := synapse.NewPrivateLinkHubsClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator synapse.PrivateLinkHubInfoListResultIterator
		err      error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		iterator, err = client.ListByResourceGroupComplete(ctx, resourceGroup)
	} else {
		iterator, err = client.ListComplete(ctx)
	}
	if err != nil {
		return nil, err
	}
	var resources []synapse.PrivateLinkHub
	for iterator.NotDone() {
		item := iterator.Value()
		resources = append(resources, item)
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return resources, err
		}
	}
	return resources, nil
}

func (az *SynapseGenerator) appendtPrivateLinkHubs(workspace *synapse.PrivateLinkHub) {
	az.AppendSimpleResource(*workspace.ID, *workspace.Name, "azurerm_synapse_private_link_hub", "synl")
}

func (az *SynapseGenerator) InitResources() error {

	workspaces, err := az.listWorkspaces()
	if err != nil {
		return err
	}
	for _, workspace := range workspaces {
		az.appendWorkspace(&workspace)
		workspaceRg, err := ParseAzureResourceID(*workspace.ID)
		if err != nil {
			return err
		}
		az.appendSQLPools(&workspace, workspaceRg)
		az.appendSparkPools(&workspace, workspaceRg)
		az.appendFirewallRule(&workspace, workspaceRg)
		az.appendManagedPrivateEndpoint(&workspace)
	}

	hubs, err := az.listPrivateLinkHubs()
	if err != nil {
		return err
	}
	for _, hub := range hubs {
		az.appendtPrivateLinkHubs(&hub)
	}
	return nil
}
