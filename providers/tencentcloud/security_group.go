// Copyright 2021 The Terraformer Authors.
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

package tencentcloud

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"
)

type SecurityGroupGenerator struct {
	TencentCloudService
}

func (g *SecurityGroupGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := vpc.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	request := vpc.NewDescribeSecurityGroupsRequest()

	var offset int64 = 0
	var pageSize int64 = 50
	allSecurityGroups := make([]*vpc.SecurityGroup, 0)

	for {
		offsetString := fmt.Sprintf("%d", offset)
		limitString := fmt.Sprintf("%d", pageSize)
		request.Offset = &offsetString
		request.Limit = &limitString
		response, err := client.DescribeSecurityGroups(request)
		if err != nil {
			return err
		}

		allSecurityGroups = append(allSecurityGroups, response.Response.SecurityGroupSet...)
		if len(response.Response.SecurityGroupSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, securityGroup := range allSecurityGroups {
		resource := terraformutils.NewResource(
			*securityGroup.SecurityGroupId,
			*securityGroup.SecurityGroupName+"_"+*securityGroup.SecurityGroupId,
			"tencentcloud_security_group",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)

		sgLiteRuleResource := terraformutils.NewResource(
			*securityGroup.SecurityGroupId,
			*securityGroup.SecurityGroupName+"_"+*securityGroup.SecurityGroupId,
			"tencentcloud_security_group_lite_rule",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, sgLiteRuleResource)
	}

	fmt.Printf("************** sg resources %v\n", len(g.Resources))
	return nil
}
