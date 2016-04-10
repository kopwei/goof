package ovs

import "github.com/kopwei/goof"

// OpenVSwitch represents the openvswitch structure
type OpenVSwitch struct {
	brname   string
	dataPath *goof.DatapathID
	ofver    []uint8
}

// GetDatapathID returns the datapathid of an ovs bridge
func (ovs *OpenVSwitch) GetDatapathID() *goof.DatapathID {
	return ovs.dataPath
}

// DoesSupportOFVer returns whether the bridge supports the openflow version or not
func (ovs *OpenVSwitch) DoesSupportOFVer(ofpversion uint8) bool {
	for _, ver := range ovs.ofver {
		if ver == ofpversion {
			return true
		}
	}
	return false
}
