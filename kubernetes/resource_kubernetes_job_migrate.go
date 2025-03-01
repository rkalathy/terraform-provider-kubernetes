// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceKubernetesJobV0() *schema.Resource {
	schemaV1 := resourceKubernetesJobV1Schema()
	schemaV0 := patchTemplatePodSpecWithResourcesFieldV0(schemaV1)
	return &schema.Resource{Schema: schemaV0}
}

func resourceKubernetesJobUpgradeV0(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	return upgradeTemplatePodSpecWithResourcesFieldV0(ctx, rawState, meta)
}
