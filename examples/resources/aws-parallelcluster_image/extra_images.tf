/** Copyright 2024 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

resource "aws-parallelcluster_image" "build-demo2" {
  image_id            = "imageBuilderDemo2"
  image_configuration = file("files/image-build-demo.yaml")
}

resource "aws-parallelcluster_image" "build-demo3" {
  image_id            = "imageBUilderDemo3"
  image_configuration = file("files/image-build-demo.yaml")
}
