/** Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
*
* Licensed under the Apache License, Version 2.0 (the "License"). You may not
* use this file except in compliance with the License. A copy of the License is
* located at
*
* http://aws.amazon.com/apache2.0/
*
* or in the "LICENSE.txt" file accompanying this file. This file is distributed
* on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, express or
* implied. See the License for the specific language governing permissions and
* limitations under the License.
*/

#######################
#### PROVIDER VARS ####
#######################

variable "role_arn" {
  type        = string
  description = "The role used by the pcluster provider."
  default     = null
}

variable "endpoint" {
  type        = string
  description = "The endpoint used by the pcluster provider."
  default     = null
}

variable "profile" {
  type        = string
  description = "The aws profile to use."
  default     = null
}

variable "region" {
  type        = string
  description = "The region to create the cluster in"
  default     = "us-west-2"
}
