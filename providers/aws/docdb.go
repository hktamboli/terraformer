// Copyright 2019 The Terraformer Authors.
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
	"context"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
)

var docDBAllowEmptyValues = []string{"tags."}

type DocDBGenerator struct {
	AWSService
}

func (g *DocDBGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := docdb.NewFromConfig(config)

	clusterPaginator := docdb.NewDescribeDBClustersPaginator(svc, &docdb.DescribeDBClustersInput{})
	for clusterPaginator.HasMorePages() {
		page, err := clusterPaginator.NextPage(context.TODO())
		if err != nil {
			return err
		}

		for _, cluster := range page.DBClusters {

			resourceName := StringValue(cluster.DBClusterIdentifier)

			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_docdb_cluster",
				"aws",
				docDBAllowEmptyValues))

			for _, member := range cluster.DBClusterMembers {
				instanceName := StringValue(member.DBInstanceIdentifier)
				g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
					instanceName,
					instanceName,
					"aws_docdb_cluster_instance",
					"aws",
					docDBAllowEmptyValues))
			}

		}
	}

	return nil
}

// PostConvertHook for add policy json as heredoc
func (g *DocDBGenerator) PostConvertHook() error {
	return nil
}
