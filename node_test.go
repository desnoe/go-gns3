package gogns3

import (
	"testing"
)

func resetTestNodeEthernetSwitch(t *testing.T) *Node {
	n := Node{
		ComputeID: "local",
		Name:      "SW1",
		NodeType:  "ethernet_switch",
		Project:   resetTestProject(t),
	}

	err := n.Create()

	if err != nil {
		t.Error("Could not create a new node")
		t.Error(err)
	}

	return &n
}

func TestNodeEthernetSwitchCreate(t *testing.T) {
	n := Node{
		ComputeID:       "local",
		ConsoleType:     "telnet",
		FirstPortName:   "eth0",
		Name:            "SW1",
		NodeType:        "ethernet_switch",
		PortNameFormat:  "eth{port0}",
		PortSegmentSize: 0,
		Project:         resetTestProject(t),
		Symbol:          ":/symbols/voice_access_server.svg",
		X:               100,
		Y:               100,
		Z:               1,
	}

	err := n.Create()

	if err != nil {
		t.Error("Could not create a new node")
		t.Error(err)
	}
	n.Read()
	if n.ComputeID != "local" {
		t.Error("This node seems to be misconfigured (copute_id != local)")
	}
	if n.ConsoleType != "telnet" {
		t.Error("This node seems to be misconfigured (console_type != telnet)")
	}
	if n.FirstPortName != "eth0" {
		t.Error("This node seems to be misconfigured (first_port_name != eth0)")
	}
	if n.Name != "SW1" {
		t.Error("This node seems to be misconfigured (name != SW1)")
	}
	if n.NodeType != "ethernet_switch" {
		t.Error("This node seems to be misconfigured (node_type != ethernet_switch)")
	}
	if n.PortNameFormat != "eth{port0}" {
		t.Error("This node seems to be misconfigured (port_name_format != eth{port0})")
	}
	if n.PortSegmentSize != 0 {
		t.Error("This node seems to be misconfigured (port_segment_size != 0)")
	}
	if n.Symbol != ":/symbols/voice_access_server.svg" {
		t.Error("This node seems to be misconfigured (symbol != :/symbols/voice_access_server.svg)")
	}
	if n.X != 100 {
		t.Error("This node seems to be misconfigured (X != 100)")
	}
	if n.Y != 100 {
		t.Error("This node seems to be misconfigured (Y != 100)")
	}
	if n.Z != 1 {
		t.Error("This node seems to be misconfigured (Z != 1)")
	}
}

func TestNodeCreateEthernetSwitchError(t *testing.T) {
	n := Node{
		ComputeID: "local",
		Name:      "",
		NodeType:  "ethernet_switch",
		Project:   resetTestProject(t),
	}

	if err := n.Create(); err != nil {
		switch e := err.(type) {
		case *ServerError:
			if e.Status != 400 {
				t.Error(e)
			}
		default:
			t.Error(e)
		}
	}
}

func TestNodeEthernetSwitchExists(t *testing.T) {
	n := *resetTestNodeEthernetSwitch(t)

	b, err := n.Exists()
	if err != nil {
		t.Error(err)
	}
	if !b {
		t.Error("This node must exist on the server")
	}
	if n.X < 0 {
		t.Error("This node seems to be misconfigured (X < 0)")
	}
	if n.Y < 0 {
		t.Error("This node seems to be misconfigured (Y < 0)")
	}
	if n.Z < 0 {
		t.Error("This node seems to be misconfigured (Z < 0)")
	}
}

func TestNodeExistsEthernetSwitchError(t *testing.T) {
	n := *resetTestNodeEthernetSwitch(t)

	n.Name = "fakefakefake"

	if _, err := n.Exists(); err != nil {
		switch e := err.(type) {
		case *ServerError:
			if e.Status != 404 {
				t.Error(e)
			}
		default:
			t.Error(e)
		}
	}
}

func TestNodeEthernetSwitchUpdate(t *testing.T) {
	n := *resetTestNodeEthernetSwitch(t)

	n.FirstPortName = "e0"
	n.Name = "SW2"
	n.PortNameFormat = "e{port0}"
	n.PortSegmentSize = 3
	n.Symbol = ":/symbols/voice_access_server.svg"
	n.X = 123
	n.Y = 123
	n.Z = 3
	err := n.Update()

	n.Read()
	if err != nil {
		t.Error("Could not update an existing node")
	}
	if n.ComputeID != "local" {
		t.Error("This node seems to be misconfigured (copute_id != local)")
	}
	if n.FirstPortName != "e0" {
		t.Error("This node seems to be misconfigured (first_port_name != e0)")
	}
	if n.Name != "SW2" {
		t.Error("This node seems to be misconfigured (name != SW2)")
	}
	if n.NodeType != "ethernet_switch" {
		t.Error("This node seems to be misconfigured (node_type != ethernet_switch)")
	}
	if n.PortNameFormat != "e{port0}" {
		t.Error("This node seems to be misconfigured (port_name_format != e{port0})")
	}
	if n.PortSegmentSize != 3 {
		t.Error("This node seems to be misconfigured (port_segment_size != 3)")
	}
	if n.Symbol != ":/symbols/voice_access_server.svg" {
		t.Error("This node seems to be misconfigured (symbol != :/symbols/voice_access_server.svg)")
	}
	if n.X != 123 {
		t.Error("This node seems to be misconfigured (X != 123)")
	}
	if n.Y != 123 {
		t.Error("This node seems to be misconfigured (Y != 123)")
	}
	if n.Z != 3 {
		t.Error("This node seems to be misconfigured (Z != 3)")
	}

}

func TestNodeEthernetSwitchUpdateError(t *testing.T) {
	n := *resetTestNodeEthernetSwitch(t)

	n.Name = ""
	err := n.Update()

	if err != nil {
		switch e := err.(type) {
		case *ServerError:
			if e.Status != 400 {
				t.Error(e)
			}
		default:
			t.Error(e)
		}
	}
}

func TestNodeEthernetSwitchLabelUpdate(t *testing.T) {
	n := *resetTestNodeEthernetSwitch(t)

	n.Label = &Label{
		Rotation: -45,
		Style:    "font-family: Times;font-size: 12;fill: #FF0000;fill-opacity: 1.0;",
		X:        123,
		Y:        123,
	}
	err := n.Update()

	n.Read()
	if err != nil {
		t.Error("Could not update an existing node")
	}
	if n.Label.Rotation != -45 {
		t.Error("This node label seems to be misconfigured (rotation != -45)")
	}
	if n.Label.Style != "font-family: Times;font-size: 12;fill: #FF0000;fill-opacity: 1.0;" {
		t.Error("This node label seems to be misconfigured (style)")
	}
	if n.Label.X != 123 {
		t.Error("This node label seems to be misconfigured (x != 123)")
	}
	if n.Label.Y != 123 {
		t.Error("This node label seems to be misconfigured (y != 123)")
	}
}

func TestNodeEthernetSwitchPropertiesUpdate(t *testing.T) {
	n := *resetTestNodeEthernetSwitch(t)

	n.Properties = NodeProperties{
		PortsMapping: []NodeEthernetPortsMapping{
			{
				Name:       "ignored",
				PortNumber: 0,
				Type:       "access",
				Vlan:       123,
			},
			{
				Name:       "ignored",
				PortNumber: 1,
				Type:       "dot1q",
				Vlan:       234,
			},
			{
				Name:       "ignored",
				PortNumber: 1,
				Type:       "qinq",
				Vlan:       345,
				EtherType:  "0x88A8",
			},
		}}
	err := n.Update()

	n.Read()
	if err != nil {
		t.Error("Could not update an existing node")
	}
	if n.Properties.PortsMapping[0].PortNumber != 0 {
		t.Error("This node ports mapping #0 seems to be misconfigured (port_number != 0)")
	}
	if n.Properties.PortsMapping[0].Type != "access" {
		t.Error("This node ports mapping #0 seems to be misconfigured (type != access)")
	}
	if n.Properties.PortsMapping[0].Vlan != 123 {
		t.Error("This node ports mapping #0 seems to be misconfigured (vlan != 123)")
	}
	if n.Properties.PortsMapping[1].PortNumber != 1 {
		t.Error("This node ports mapping #1 seems to be misconfigured (port_number != 1)")
	}
	if n.Properties.PortsMapping[1].Type != "dot1q" {
		t.Error("This node ports mapping #2 seems to be misconfigured (type != dot1q)")
	}
	if n.Properties.PortsMapping[1].Vlan != 234 {
		t.Error("This node ports mapping #2 seems to be misconfigured (vlan != 234)")
	}
	if n.Properties.PortsMapping[2].PortNumber != 2 {
		t.Error("This node ports mapping #2 seems to be misconfigured (port_number != 2)")
	}
	if n.Properties.PortsMapping[2].Type != "qinq" {
		t.Error("This node ports mapping #3 seems to be misconfigured (type != qinq)")
	}
	if n.Properties.PortsMapping[2].Vlan != 345 {
		t.Error("This node ports mapping #3 seems to be misconfigured (vlan != 345)")
	}
}

func TestNodeEthernetSwitchDelete(t *testing.T) {
	n := *resetTestNodeEthernetSwitch(t)
	// clear the UUID to make sure a Read is automatically performed before
	n.UUID = ""

	if err := n.Delete(); err != nil {
		t.Error("Could not delete an existing node")
		t.Error(err)
	}
}

func TestNodeEthernetSwitchDeleteError(t *testing.T) {
	n := *resetTestNodeEthernetSwitch(t)
	n.UUID = "11111111-1111-1111-1111-111111111111"

	if err := n.Delete(); err != nil {
		switch e := err.(type) {
		case *ServerError:
			if e.Status != 404 {
				t.Error(e)
			}
		default:
			t.Error(e)
		}
	}
}
