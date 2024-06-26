// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
func _TestAccObjectResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
resource "observability_object" "test" {
	type_name = "anzen:cloudConnection"
	object_id = "just-terraform-testing"
	layer_type = "TENANT"
	layer_id = "0eb4e853-34fb-4f77-b3fc-b9cd3b462366"
	import_id = "anzen:cloudConnection|just-terraform-testing|TENANT|0eb4e853-34fb-4f77-b3fc-b9cd3b462366"
	data = jsonencode(
		{
			"cloudType": "AWS",
			"connectionName": "just-terraform-testing",
			"region": "us-east-2",
			"accessKey": "**********",
			"secretAccessKey": "**********",
			"s3AccessLogBucket": "s3://s3-sanity-logging/",
			"athenaOutputBucket": "s3://s3-sanity-athena-logs/",
			"createTimestamp": "",
			"accountId": "81892134343434"
		}
	)
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("observability_object.test", "type_name", "anzen:cloudConnection"),

					resource.TestCheckResourceAttr("observability_object.test",
						"import_id",
						"anzen:cloudConnection|just-terraform-testing|TENANT|0eb4e853-34fb-4f77-b3fc-b9cd3b462366"),
					resource.TestCheckResourceAttr("observability_object.test", "layer_id", "0eb4e853-34fb-4f77-b3fc-b9cd3b462366"),
					resource.TestCheckResourceAttr("observability_object.test", "layer_type", "TENANT"),
					resource.TestCheckResourceAttr("observability_object.test", "object_id", "just-terraform-testing"),
					//nolint:lll // Due to payload nature of data field
					resource.TestCheckResourceAttr("observability_object.test", "data", "{\"accessKey\":\"**********\",\"accountId\":\"81892134343434\",\"athenaOutputBucket\":\"s3://s3-sanity-athena-logs/\",\"cloudType\":\"AWS\",\"connectionName\":\"just-terraform-testing\",\"createTimestamp\":\"\",\"region\":\"us-east-2\",\"s3AccessLogBucket\":\"s3://s3-sanity-logging/\",\"secretAccessKey\":\"**********\"}"),

					// Verify placeholder id attribute
					resource.TestCheckResourceAttrSet("observability_object.test", "id"),
					resource.TestCheckResourceAttr("observability_object.test", "id", "placeholder"),
				),
			},
			// ImportState testing
			// TO DO add this after the plugin is plublished
			// {
			// 	ResourceName:      "observability_object.test",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
			// Update and Read testing
			{
				Config: providerConfig + `
resource "observability_object" "test" {
				type_name = "anzen:cloudConnection"
				object_id = "just-terraform-testing"
				layer_type = "TENANT"
				layer_id = "0eb4e853-34fb-4f77-b3fc-b9cd3b462366"
				import_id = "anzen:cloudConnection|just-terraform-testing|TENANT|0eb4e853-34fb-4f77-b3fc-b9cd3b462366"
				data = jsonencode(
					{
						"cloudType": "GCP",
						"connectionName": "just-terraform-testing",
						"region": "us-west-2",
						"accessKey": "**********",
						"secretAccessKey": "**********",
						"s3AccessLogBucket": "s3://s3-sanity-logging/",
						"athenaOutputBucket": "s3://s3-sanity-athena-logs/",
						"createTimestamp": "",
						"accountId": "81892134343434"
					}
				)
			}
			`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("observability_object.test", "type_name", "anzen:cloudConnection"),
					resource.TestCheckResourceAttr("observability_object.test",
						"import_id",
						"anzen:cloudConnection|just-terraform-testing|TENANT|0eb4e853-34fb-4f77-b3fc-b9cd3b462366"),
					resource.TestCheckResourceAttr("observability_object.test", "layer_id", "0eb4e853-34fb-4f77-b3fc-b9cd3b462366"),
					resource.TestCheckResourceAttr("observability_object.test", "layer_type", "TENANT"),
					resource.TestCheckResourceAttr("observability_object.test", "object_id", "just-terraform-testing"),
					//nolint:lll // Due to payload nature of data field
					resource.TestCheckResourceAttr("observability_object.test", "data", "{\"accessKey\":\"**********\",\"accountId\":\"81892134343434\",\"athenaOutputBucket\":\"s3://s3-sanity-athena-logs/\",\"cloudType\":\"GCP\",\"connectionName\":\"just-terraform-testing\",\"createTimestamp\":\"\",\"region\":\"us-west-2\",\"s3AccessLogBucket\":\"s3://s3-sanity-logging/\",\"secretAccessKey\":\"**********\"}"),

					// Verify placeholder id attribute
					resource.TestCheckResourceAttrSet("observability_object.test", "id"),
					resource.TestCheckResourceAttr("observability_object.test", "id", "placeholder"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
