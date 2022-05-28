// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type L2VPNXconnectGroupP2PNeighborIPv4 struct {
	Device          types.String `tfsdk:"device"`
	Id              types.String `tfsdk:"id"`
	GroupName       types.String `tfsdk:"group_name"`
	P2pXconnectName types.String `tfsdk:"p2p_xconnect_name"`
	Address         types.String `tfsdk:"address"`
	PwId            types.Int64  `tfsdk:"pw_id"`
	PwClass         types.String `tfsdk:"pw_class"`
}

func (data L2VPNXconnectGroupP2PNeighborIPv4) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/xconnect/groups/group[group-name=%s]/p2ps/p2p[p2p-xconnect-name=%s]/neighbor/ipv4s/ipv4[address=%s][pw-id=%v]", data.GroupName.Value, data.P2pXconnectName.Value, data.Address.Value, data.PwId.Value)
}

func (data L2VPNXconnectGroupP2PNeighborIPv4) toBody() string {
	body := "{}"
	if !data.PwClass.Null && !data.PwClass.Unknown {
		body, _ = sjson.Set(body, "pw-class", data.PwClass.Value)
	}
	return body
}

func (data *L2VPNXconnectGroupP2PNeighborIPv4) updateFromBody(res []byte) {
	if value := gjson.GetBytes(res, "pw-class"); value.Exists() {
		data.PwClass.Value = value.String()
	} else {
		data.PwClass.Null = true
	}
}

func (data *L2VPNXconnectGroupP2PNeighborIPv4) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "pw-class"); value.Exists() {
		data.PwClass.Value = value.String()
		data.PwClass.Null = false
	}
}

func (data *L2VPNXconnectGroupP2PNeighborIPv4) fromPlan(plan L2VPNXconnectGroupP2PNeighborIPv4) {
	data.Device = plan.Device
	data.GroupName.Value = plan.GroupName.Value
	data.P2pXconnectName.Value = plan.P2pXconnectName.Value
	data.Address.Value = plan.Address.Value
	data.PwId.Value = plan.PwId.Value
}

func (data *L2VPNXconnectGroupP2PNeighborIPv4) setUnknownValues() {
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
	if data.Address.Unknown {
		data.Address.Unknown = false
		data.Address.Null = true
	}
	if data.PwId.Unknown {
		data.PwId.Unknown = false
		data.PwId.Null = true
	}
	if data.PwClass.Unknown {
		data.PwClass.Unknown = false
		data.PwClass.Null = true
	}
}

func (data *L2VPNXconnectGroupP2PNeighborIPv4) getDeletedListItems(state L2VPNXconnectGroupP2PNeighborIPv4) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *L2VPNXconnectGroupP2PNeighborIPv4) getEmptyLeafsDelete() []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}
