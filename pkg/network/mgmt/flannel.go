package mgmt

import (
	"fmt"

	"github.com/vishvananda/netlink"

	"github.com/harvester/harvester-network-controller/pkg/network/iface"
)

const flannelVtepName = "flannel.1"

type FlannelNetwork struct {
	vtep *netlink.Vxlan
	nic  iface.IFace
}

func NewFlannelNetwork() (*FlannelNetwork, error) {
	l, err := netlink.LinkByName(flannelVtepName)
	if err != nil {
		return nil, err
	}
	vtep, ok := l.(*netlink.Vxlan)
	if !ok {
		return nil, fmt.Errorf("got data of type %T but wanted *netlink.Vxlan", l)
	}

	nic, err := iface.GetLink(iface.EmptyName, vtep.VtepDevIndex)
	if err != nil {
		return nil, err
	}

	return &FlannelNetwork{
		vtep: vtep,
		nic:  nic,
	}, nil
}

func (f *FlannelNetwork) Type() string {
	return "flannel"
}

func (f *FlannelNetwork) Setup(nic string, vids []uint16) error {
	return nil
}

func (f *FlannelNetwork) Teardown() error {
	return nil
}

func (f *FlannelNetwork) NIC() iface.IFace {
	return f.nic
}
