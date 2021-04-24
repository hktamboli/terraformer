// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aws

import (
	"os"
	"strconv"

	"github.com/hashicorp/terraform/addrs"
	"github.com/pkg/errors"
	"github.com/zclconf/go-cty/cty"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type AWSProvider struct { //nolint
	terraformutils.Provider
	region  string
	profile string
}

const GlobalRegion = "aws-global"
const MainRegionPublicPartition = "us-east-1"
const NoRegion = ""

// SupportedGlobalResources should be bound to a default region. AWS doesn't specify in which region default services are
// placed (see  https://docs.aws.amazon.com/general/latest/gr/rande.html), so we shouldn't assume any region as well
//
// AWS WAF V2 if added, should not be included in this list since it is a composition of regional and global resources.
var SupportedGlobalResources = []string{
	"budgets",
	"cloudfront",
	"ecrpublic",
	"iam",
	"organization",
	"route53",
	"waf",
}

func (p AWSProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"alb": {
			"sg":     []string{"security_groups", "id"},
			"subnet": []string{"subnets", "id"},
			"alb": []string{
				"load_balancer_arn", "id",
				"listener_arn", "id",
				// TF ALB TG attachment logic doesn't work well with references (doesn't interpolate)
			},
		},
		"auto_scaling": {
			"sg":     []string{"security_groups", "id"},
			"subnet": []string{"vpc_zone_identifier", "id"},
		},
		"ec2_instance": {
			"sg":     []string{"vpc_security_group_ids", "id"},
			"subnet": []string{"subnet_id", "id"},
			"ebs":    []string{"ebs_block_device", "id"},
		},
		"elasticache": {
			"vpc":    []string{"vpc_id", "id"},
			"subnet": []string{"subnet_ids", "id"},
			"sg":     []string{"security_group_ids", "id"},
		},
		"ebs": {
			// TF EBS attachment logic doesn't work well with references (doesn't interpolate)
		},
		"ecs": {
			// ECS is not able anymore to support references (doesn't interpolate)
			"subnet": []string{"network_configuration.subnets", "id"},
			"sg":     []string{"network_configuration.security_groups", "id"},
		},
		"eks": {
			"subnet": []string{"vpc_config.subnet_ids", "id"},
			"sg":     []string{"vpc_config.security_group_ids", "id"},
		},
		"elb": {
			"sg":     []string{"security_groups", "id"},
			"subnet": []string{"subnets", "id"},
		},
		"igw": {"vpc": []string{"vpc_id", "id"}},
		"msk": {
			"subnet": []string{"broker_node_group_info.client_subnets", "id"},
			"sg":     []string{"broker_node_group_info.security_groups", "id"},
		},
		"nacl": {
			"subnet": []string{"subnet_ids", "id"},
			"vpc":    []string{"vpc_id", "id"},
		},
		"organization": {
			"organization": []string{
				"policy_id", "id",
				"parent_id", "id",
				"target_id", "id",
			},
		},
		"rds": {
			"subnet": []string{"subnet_ids", "id"},
			"sg":     []string{"vpc_security_group_ids", "id"},
		},
		"route_table": {
			"route_table": []string{"route_table_id", "id"},
			"subnet":      []string{"subnet_id", "id"},
			"vpc":         []string{"vpc_id", "id"},
		},
		"sns": {
			"sns": []string{"topic_arn", "id"},
			"sqs": []string{"endpoint", "arn"},
		},
		"sg": {
			"sg": []string{
				"egress.security_groups", "id",
				"ingress.security_groups", "id",
				"security_group_id", "id",
				"source_security_group_id", "id",
			},
		},
		"subnet": {"vpc": []string{"vpc_id", "id"}},
		"transit_gateway": {
			"vpc":             []string{"vpc_id", "id"},
			"transit_gateway": []string{"transit_gateway_id", "id"},
			"subnet":          []string{"subnet_ids", "id"},
			"vpn_connection":  []string{"vpn_connection_id", "id"},
		},
		"vpn_gateway": {"vpc": []string{"vpc_id", "id"}},
		"vpn_connection": {
			"customer_gateway": []string{"customer_gateway_id", "id"},
			"vpn_gateway":      []string{"vpn_gateway_id", "id"},
		},
	}
}

func (p AWSProvider) GetProviderData(arg ...string) map[string]interface{} {
	awsConfig := map[string]interface{}{}

	if p.region == GlobalRegion {
		awsConfig["region"] = MainRegionPublicPartition // For TF to workaround terraform-providers/terraform-provider-aws#1043
	} else if p.region != NoRegion {
		awsConfig["region"] = p.region
	}

	return map[string]interface{}{
		"provider": map[string]interface{}{
			"aws": awsConfig,
		},
	}
}

func (p *AWSProvider) GetConfig() cty.Value {
	if p.region != GlobalRegion {
		return cty.ObjectVal(map[string]cty.Value{
			"region":                 cty.StringVal(p.region),
			"skip_region_validation": cty.True,
		})
	}
	return cty.ObjectVal(map[string]cty.Value{
		"region":                 cty.StringVal(""),
		"skip_region_validation": cty.True,
	})
}

func (p *AWSProvider) GetBasicConfig() cty.Value {
	return p.GetConfig()
}

// check projectName in env params
func (p *AWSProvider) Init(args []string) error {
	p.region = args[0]
	p.profile = args[1]

	// Terraformer accepts region and profile configuration, so we must detect what env variables to adjust to make Go SDK rely on them. AWS_SDK_LOAD_CONFIG here must be checked to determine correct variable to set.
	enableSharedConfig, _ := strconv.ParseBool(os.Getenv("AWS_SDK_LOAD_CONFIG"))
	var err error
	if p.region != GlobalRegion && p.region != NoRegion {
		if enableSharedConfig {
			err = os.Setenv("AWS_DEFAULT_REGION", p.region)
		} else {
			err = os.Setenv("AWS_REGION", p.region)
		}
		if err != nil {
			return err
		}
	}

	if p.profile != "default" && p.profile != "" {
		if enableSharedConfig {
			err = os.Setenv("AWS_DEFAULT_PROFILE", p.profile)
		} else {
			err = os.Setenv("AWS_PROFILE", p.profile)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *AWSProvider) GetName() string {
	return "aws"
}

func (p *AWSProvider) GetProviderSource() addrs.Provider {
	return addrs.NewDefaultProvider("aws")
}

func (p *AWSProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("aws: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"region":                 p.region,
		"profile":                p.profile,
		"skip_region_validation": true,
	})
	return nil
}

// GetAWSSupportService return map of support service for AWS
func (p *AWSProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"ec2_instance":      &AwsFacade{service: &Ec2Generator{}},
	}
}

func StringValue(value *string) string {
	if value != nil {
		return *value
	}
	return ""
}
