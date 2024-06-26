---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aws-parallelcluster_image Resource - aws-parallelcluster"
subcategory: ""
description: |-
  Create a custom AWS ParallelCluster image in an AWS Region.
---

# aws-parallelcluster_image (Resource)

Create a custom AWS ParallelCluster image in an AWS Region.

## Example Usage

```terraform
resource "aws-parallelcluster_image" "build-demo" {
  image_id            = "imageBuilderDemo"
  rollback_on_failure = false
  image_configuration = file("files/image-build-demo.yaml")
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `image_configuration` (String) The image configuration as a YAML document.
- `image_id` (String) Id of the Image that will be built.

### Optional

- `region` (String) The AWS Region that the image is in.
- `rollback_on_failure` (Boolean) If set to true, image stack rollback occurs if the image fails to create. The default is true.
- `suppress_validators` (List of String) Identify one or more configuration validators to suppress.
- `validation_failure_level` (String) The minimum validation level that causes image create to fail. The default is ERROR.

### Read-Only

- `ami_id` (String) EC2 AMI Id
- `cloudformation_stack_arn` (String) The Amazon Resource Name (ARN) of the main CloudFormation stack.
- `cloudformation_stack_status` (String) Status of the cloudformation stack.
- `id` (String) Id of the Image that will be built.
- `image_build_status` (String) Status of the image.
- `version` (String) The AWS ParallelImage version that's used to create the image.
