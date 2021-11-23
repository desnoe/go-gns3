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
		t.Error("Could not create a new Ethernet Switch node")
		t.Error(err)
	}

	return &n
}

func resetTestNodeQemu(t *testing.T) *Node {
	n := Node{
		ComputeID: "local",
		Name:      "PC1",
		NodeType:  "qemu",
		Project:   resetTestProject(t),
		Properties: NodeProperties{
			Platform: "x86_64",
		},
	}

	err := n.Create()

	if err != nil {
		t.Error("Could not create a new QEMU node")
		t.Error(err)
	}

	return &n
}

func resetTestNodeVpcs(t *testing.T) *Node {
	n := Node{
		ComputeID: "local",
		Name:      "PC1",
		NodeType:  "vpcs",
		Project:   resetTestProject(t),
	}

	err := n.Create()

	if err != nil {
		t.Error("Could not create a new VPCS node")
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

func TestNodeQemuCreate(t *testing.T) {
	n := Node{
		ComputeID:       "local",
		ConsoleType:     "telnet",
		FirstPortName:   "eth0",
		Name:            "QEMU1",
		NodeType:        "qemu",
		PortNameFormat:  "eth{port0}",
		PortSegmentSize: 0,
		Project:         resetTestProject(t),
		Properties: NodeProperties{
			Platform: "x86_64",
		},
		Symbol: ":/symbols/qemu_guest.svg",
		X:      100,
		Y:      100,
		Z:      1,
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
	if n.Name != "QEMU1" {
		t.Error("This node seems to be misconfigured (name != QEMU1)")
	}
	if n.NodeType != "qemu" {
		t.Error("This node seems to be misconfigured (node_type != qemu)")
	}
	if n.PortNameFormat != "eth{port0}" {
		t.Error("This node seems to be misconfigured (port_name_format != eth{port0})")
	}
	if n.PortSegmentSize != 0 {
		t.Error("This node seems to be misconfigured (port_segment_size != 0)")
	}
	if n.Symbol != ":/symbols/qemu_guest.svg" {
		t.Error("This node seems to be misconfigured (symbol != :/symbols/qemu_guest.svg)")
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

func TestNodeQemuDelete(t *testing.T) {
	n := *resetTestNodeQemu(t)
	// clear the UUID to make sure a Read is automatically performed before
	n.UUID = ""

	if err := n.Delete(); err != nil {
		t.Error("Could not delete an existing node")
		t.Error(err)
	}
}

func TestNodeQemuPropertiesUpdate(t *testing.T) {
	n := *resetTestNodeQemu(t)

	n.Properties = NodeProperties{
		AdapterType:       "virtio",
		Adapters:          4,
		BootPriority:      "cd",
		CPUThrottling:     1,
		CPUs:              2,
		HdaDiskInterface:  "virtio",
		KernelCommandLine: "noapic",
		LegacyNetworking:  true,
		MacAddress:        "01:01:01:01:01:01",
		Options:           "-nographic",
		Platform:          "x86_64",
		ProcessPriority:   "high",
		RAM:               512,
		Usage:             "This is a test node",
	}

	err := n.Update()

	n.Read()
	if err != nil {
		t.Error("Could not update an existing node")
	}
	if n.Properties.AdapterType != "virtio" {
		t.Error("This node property seems to be misconfigured (adapter_type != virtio)")
	}
	if n.Properties.Adapters != 4 {
		t.Error("This node property seems to be misconfigured (adapters != 4)")
	}
	if n.Properties.BootPriority != "cd" {
		t.Error("This node property seems to be misconfigured (boot_priority != cd)")
	}
	if n.Properties.CPUThrottling != 1 {
		t.Error("This node property seems to be misconfigured (cpu_throttling != 1)")
	}
	if n.Properties.CPUs != 2 {
		t.Error("This node property seems to be misconfigured (cpus != 2)")
	}
	if n.Properties.HdaDiskInterface != "virtio" {
		t.Error("This node property seems to be misconfigured (hda_disk_interface != virtio)")
	}
	if n.Properties.KernelCommandLine != "noapic" {
		t.Error("This node property seems to be misconfigured (kernel_command_line != noapic)")
	}
	if n.Properties.LegacyNetworking != true {
		t.Error("This node property seems to be misconfigured (legacy_networking != true)")
	}
	if n.Properties.MacAddress != "01:01:01:01:01:01" {
		t.Error("This node property seems to be misconfigured (mac_address != 01:01:01:01:01:01)")
	}
	if n.Properties.Options != "-nographic" {
		t.Error("This node property seems to be misconfigured (options != -nographic)")
	}
	if n.Properties.Platform != "x86_64" {
		t.Error("This node property seems to be misconfigured (platform != x86_64)")
	}
	if n.Properties.ProcessPriority != "high" {
		t.Error("This node property seems to be misconfigured (process_priority != high)")
	}
	if n.Properties.RAM != 512 {
		t.Error("This node property seems to be misconfigured (ram != 512)")
	}
	if n.Properties.Usage != "This is a test node" {
		t.Error("This node property seems to be misconfigured (usage != This is a test node)")
	}
}

func TestNodeQemuPropertiesUpdateImageError(t *testing.T) {
	n := *resetTestNodeQemu(t)

	n.Properties = NodeProperties{
		HdaDiskImage: "missing.qcow2",
	}

	err := n.Update()

	if err != nil {
		switch e := err.(type) {
		case *ServerError:
			if e.Status != 409 {
				t.Error(e)
			}
		default:
			t.Error(e)
		}
	}
}

func TestNodeVpcsCreate(t *testing.T) {
	n := Node{
		ComputeID:       "local",
		ConsoleType:     "telnet",
		FirstPortName:   "eth0",
		Name:            "PC1",
		NodeType:        "vpcs",
		PortNameFormat:  "eth{port0}",
		PortSegmentSize: 0,
		Project:         resetTestProject(t),
		Symbol:          ":/symbols/vpcs_guest.svg",
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
	if n.Name != "PC1" {
		t.Error("This node seems to be misconfigured (name != PC1)")
	}
	if n.NodeType != "vpcs" {
		t.Error("This node seems to be misconfigured (node_type != vpcs)")
	}
	if n.PortNameFormat != "eth{port0}" {
		t.Error("This node seems to be misconfigured (port_name_format != eth{port0})")
	}
	if n.PortSegmentSize != 0 {
		t.Error("This node seems to be misconfigured (port_segment_size != 0)")
	}
	if n.Symbol != ":/symbols/vpcs_guest.svg" {
		t.Error("This node seems to be misconfigured (symbol != :/symbols/vpcs_guest.svg)")
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

func TestNodeVpcsDelete(t *testing.T) {
	n := *resetTestNodeVpcs(t)
	// clear the UUID to make sure a Read is automatically performed before
	n.UUID = ""

	if err := n.Delete(); err != nil {
		t.Error("Could not delete an existing node")
		t.Error(err)
	}
}
