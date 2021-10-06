/*******************************************************************************
* IBM Cloud Kubernetes Service, 5737-D43
* (C) Copyright IBM Corp. 2021 All Rights Reserved.
*
* SPDX-License-Identifier: Apache2.0
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*******************************************************************************/

package ibm

import (
	"cloud.ibm.com/cloud-provider-vpc-controller/pkg/vpcctl"
)

// InitCloudVpc - Initialize the VPC cloud logic
func (c *Cloud) InitCloudVpc(enablePrivateEndpoint bool) error {
	if c.Vpc == nil {
		options := &vpcctl.CloudVpcOptions{
			ClusterID:     c.Config.Prov.ClusterID,
			EnablePrivate: enablePrivateEndpoint,
		}
		cloudVpc, err := vpcctl.NewCloudVpc(c.KubeClient, options)
		if err != nil {
			return err
		}
		c.Vpc = cloudVpc
	}
	return nil
}
