// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type L2VPNXconnectGroupP2P struct {
	Device          types.String `tfsdk:"device"`
	Id              types.String `tfsdk:"id"`
	GroupName       types.String `tfsdk:"group_name"`
	P2pXconnectName types.String `tfsdk:"p2p_xconnect_name"`
	Description     types.String `tfsdk:"description"`
}

func (data L2VPNXconnectGroupP2P) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/xconnect/groups/group[group-name=%s]/p2ps/p2p[p2p-xconnect-name=%s]", data.GroupName.Value, data.P2pXconnectName.Value)
}

func (data L2VPNXconnectGroupP2P) toBody() string {
	body := "{}"
	if !data.Description.Null && !data.Description.Unknown {
		body, _ = sjson.Set(body, "description", data.Description.Value)
	}
	return body
}

func (data *L2VPNXconnectGroupP2P) updateFromBody(res []byte) {
	if value := gjson.GetBytes(res, "description"); value.Exists() {
		data.Description.Value = value.String()
	} else {
		data.Description.Null = true
	}
}

func (data *L2VPNXconnectGroupP2P) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "description"); value.Exists() {
		data.Description.Value = value.String()
		data.Description.Null = false
	}
}

func (data *L2VPNXconnectGroupP2P) fromPlan(plan L2VPNXconnectGroupP2P) {
	data.Device = plan.Device
	data.GroupName.Value = plan.GroupName.Value
	data.P2pXconnectName.Value = plan.P2pXconnectName.Value
}

func (data *L2VPNXconnectGroupP2P) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.GroupName.Unknown {
		data.GroupName.Unknown = false
		data.GroupName.Null = true
	}
	if data.P2pXconnectName.Unknown {
		data.P2pXconnectName.Unknown = false
		data.P2pXconnectName.Null = true
	}
	if data.Description.Unknown {
		data.Description.Unknown = false
		data.Description.Null = true
	}
}

func (data *L2VPNXconnectGroupP2P) getDeletedListItems(state L2VPNXconnectGroupP2P) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *L2VPNXconnectGroupP2P) getEmptyLeafsDelete() []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}
