package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type NetworkLoadBalancerForwardingRuleGenerator struct {
	IonosCloudService
}

func (g *NetworkLoadBalancerForwardingRuleGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	resource_type := "ionoscloud_networkloadbalancer_forwardingrule"

	datacenters, err := helpers.GetAllDatacenters(*cloudApiClient)
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		networkLoadBalancerResponse, _, err := cloudApiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersGet(context.TODO(), *datacenter.Id).Execute()
		if err != nil {
			return err
		}
		if networkLoadBalancerResponse.Items == nil {
			log.Printf(
				"[WARNING] expected a response containing network load balancers but received 'nil' instead, skipping search for datacenter with ID: %v.",
				*datacenter.Id)
			continue
		}
		networkLoadBalancers := *networkLoadBalancerResponse.Items
		for _, nlb := range networkLoadBalancers {
			forwardingRulesResponse, _, err := cloudApiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesGet(context.TODO(), *datacenter.Id, *nlb.Id).Depth(1).Execute()
			if err != nil {
				return err
			}
			if forwardingRulesResponse.Items == nil {
				log.Printf(
					"[WARNING] expected a response containing forwarding rules but received 'nil' instead, skipping search for NLB with ID: %v, datacenter ID: %v",
					*nlb.Id,
					*datacenter.Id)
				continue
			}
			forwardingRules := *forwardingRulesResponse.Items
			for _, fr := range forwardingRules {
				if fr.Properties == nil || fr.Properties.Name == nil {
					log.Printf(
						"[WARNING] 'nil' values in the response for the forwarding rule with ID %v, NLB ID: %v, datacenter ID: %v",
						*fr.Id,
						*nlb.Id,
						*datacenter.Id)
					continue
				}
				g.Resources = append(g.Resources, terraformutils.NewResource(
					*fr.Id,
					*fr.Properties.Name+"-"+*fr.Id,
					resource_type,
					helpers.Ionos,
					map[string]string{helpers.DcId: *datacenter.Id, "networkloadbalancer_id": *nlb.Id},
					[]string{},
					map[string]interface{}{}))
			}
		}
	}
	return nil
}
