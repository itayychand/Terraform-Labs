package migration

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-provider-azurerm/helpers/azure"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/devtestlabs/parse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tags"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

var _ pluginsdk.StateUpgrade = DevTestVirtualNetworkUpgradeV0ToV1{}

type DevTestVirtualNetworkUpgradeV0ToV1 struct{}

func (DevTestVirtualNetworkUpgradeV0ToV1) Schema() map[string]*pluginsdk.Schema {
	return devTestVirtualNetworkSchemaForV0AndV1()
}

func (DevTestVirtualNetworkUpgradeV0ToV1) UpgradeFunc() pluginsdk.StateUpgraderFunc {
	return func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
		// old:
		// 	/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.devtestlab/labs/{labName}/virtualnetworks/{virtualNetworkName}
		// new:
		// 	/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualNetworks/{virtualNetworkName}

		oldId, err := azure.ParseAzureResourceID(strings.Replace(rawState["id"].(string), "/virtualnetworks/", "/virtualNetworks/", 1))
		if err != nil {
			return rawState, err
		}

		labName, err := oldId.PopSegment("labs")
		if err != nil {
			return rawState, err
		}

		vnetName, err := oldId.PopSegment("virtualNetworks")
		if err != nil {
			return rawState, err
		}

		newId := parse.NewDevTestVirtualNetworkID(oldId.SubscriptionID, oldId.ResourceGroup, labName, vnetName)

		log.Printf("[DEBUG] Updating ID from %q to %q", oldId, newId.ID())

		rawState["id"] = newId.ID()

		return rawState, nil
	}
}

func devTestVirtualNetworkSchemaForV0AndV1() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:     pluginsdk.TypeString,
			Required: true,
		},

		"lab_name": {
			Type:     pluginsdk.TypeString,
			Required: true,
		},

		"resource_group_name": azure.SchemaResourceGroupNameDiffSuppress(),

		"description": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"subnet": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			Computed: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"name": {
						Type:     pluginsdk.TypeString,
						Computed: true,
					},

					"use_in_virtual_machine_creation": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},

					"use_public_ip_address": {
						Type:     pluginsdk.TypeString,
						Optional: true,
					},
				},
			},
		},

		"tags": tags.Schema(),

		"unique_identifier": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},
	}
}
