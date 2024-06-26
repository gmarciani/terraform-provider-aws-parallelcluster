######################
#### CLUSTER VARS ####
######################

variable "cluster_name" {
  type        = string
  description = "The name of the cluster."
  default     = "pcluster-example"
}

variable "max_nodes" {
  type        = string
  description = "The maximum number of compute nodes for the cluster."
  default     = "4"
}

variable "min_nodes" {
  type        = string
  description = "The minimum number of compute nodes for the cluster."
  default     = "1"
}

variable "subnet" {
  type        = string
  description = "The subnet to deploy a cluster to."
  default     = null
}

variable "key_pair" {
  type        = string
  description = "The keypair used to deploy a cluster."
  default     = null
}

variable "region" {
  type        = string
  description = "The region to create the cluster in"
  default     = "us-east-1"
}

variable "compute_node_type" {
  type        = string
  description = "The type of instance for compute nodes."
  default     = "t3.micro"
}

variable "head_node_type" {
  type        = string
  description = "The type of instance for head nodes."
  default     = "t3.micro"
}
#######################
#### PROVIDER VARS ####
#######################

variable "role_arn" {
  type        = string
  description = "The role used by the ParallelCluster provider."
  default     = null
}

variable "endpoint" {
  type        = string
  description = "The endpoint used by the ParallelCluster provider."
  default     = null
}
