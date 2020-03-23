package gogns3

import (
	"testing"
)

func resetTestLinkEthernetSwitchLab(t *testing.T) (*Project, *Node, *Node) {
	p := resetTestProject(t)
	n1 := &Node{
		ComputeID: "local",
		Name:      "SW1",
		NodeType:  "ethernet_switch",
		Project:   p,
		X:         0,
	}

	n1.Create()

	n2 := &Node{
		ComputeID: "local",
		Name:      "SW2",
		NodeType:  "ethernet_switch",
		Project:   p,
		X:         100,
	}

	n2.Create()

	return p, n1, n2
}

func resetTestLinkEthernetSwitchLab2(t *testing.T) (*Project, *Node, *Node, *Link) {
	p, n1, n2 := resetTestLinkEthernetSwitchLab(t)

	l := &Link{
		Project: p,
		Nodes: []LinkNode{
			{
				AdapterNumber: 0,
				NodeID:        n1.UUID,
				PortNumber:    0,
			},
			{
				AdapterNumber: 0,
				NodeID:        n2.UUID,
				PortNumber:    0,
			},
		},
		Capturing: false,
		LinkType:  "ethernet",
		Suspend:   false,
	}
	l.Create()

	return p, n1, n2, l
}

func resetTestLinkVpcsLab(t *testing.T) (*Project, *Node, *Node) {
	p := resetTestProject(t)
	n1 := &Node{
		ComputeID: "local",
		Name:      "PC1",
		NodeType:  "vpcs",
		Project:   p,
		X:         0,
	}

	n1.Create()

	n2 := &Node{
		ComputeID: "local",
		Name:      "PC2",
		NodeType:  "vpcs",
		Project:   p,
		X:         100,
	}

	n2.Create()

	return p, n1, n2
}

func createAndTestLink(t *testing.T, p *Project, n1 *Node, n2 *Node) {
	l := Link{
		Project: p,
		Nodes: []LinkNode{
			{
				AdapterNumber: 0,
				NodeID:        n1.UUID,
				PortNumber:    0,
			},
			{
				AdapterNumber: 0,
				NodeID:        n2.UUID,
				PortNumber:    0,
			},
		},
		Capturing: false,
		LinkType:  "ethernet",
		Suspend:   false,
	}

	err := l.Create()

	if err != nil {
		t.Error("Could not create a new link")
		t.Error(err)
	}
	l.Read()
	if l.Capturing != false {
		t.Error("This link seems to be misconfigured (capturing != false)")
	}
	if l.LinkType != "ethernet" {
		t.Error("This link seems to be misconfigured (link_type != ethernet)")
	}
	if l.Suspend != false {
		t.Error("This link seems to be misconfigured (suspend != false)")
	}
}

func TestNodeEthernetSwitchLinkCreate(t *testing.T) {
	p, n1, n2 := resetTestLinkEthernetSwitchLab(t)

	createAndTestLink(t, p, n1, n2)
}

func TestNodeCreateEthernetSwitchLinkError(t *testing.T) {
	p, n1, n2 := resetTestLinkEthernetSwitchLab(t)

	l := Link{
		Project: p,
		Nodes: []LinkNode{
			{
				AdapterNumber: 99,
				NodeID:        n1.UUID,
				PortNumber:    0,
			},
			{
				AdapterNumber: 99,
				NodeID:        n2.UUID,
				PortNumber:    0,
			},
		},
		Capturing: false,
		LinkType:  "ethernet",
		Suspend:   false,
	}

	if err := l.Create(); err != nil {
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

func TestNodeEthernetSwitchLinkExists(t *testing.T) {
	_, _, _, l := resetTestLinkEthernetSwitchLab2(t)

	b, err := l.Exists()
	if err != nil {
		t.Error(err)
	}
	if !b {
		t.Error("This link must exist on the server")
	}
}

func TestNodeExistsEthernetSwitchLinkError(t *testing.T) {
	_, _, _, l := resetTestLinkEthernetSwitchLab2(t)

	l.UUID = "11111111-1111-1111-1111-111111111111"

	if _, err := l.Exists(); err != nil {
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

func TestNodeEthernetSwitchLinkUpdate(t *testing.T) {
	_, _, _, l := resetTestLinkEthernetSwitchLab2(t)

	l.Suspend = true
	err := l.Update()

	l.Read()
	if err != nil {
		t.Error("Could not update an existing link")
	}
	if l.Suspend != true {
		t.Error("This node seems to be misconfigured (suspend != true)")
	}

}

func TestNodeEthernetSwitchLinkUpdateError(t *testing.T) {
	_, _, _, l := resetTestLinkEthernetSwitchLab2(t)

	l.LinkType = "foo"
	err := l.Update()

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

func TestNodeEthernetSwitchLinkLabelUpdate(t *testing.T) {
	_, _, _, l := resetTestLinkEthernetSwitchLab2(t)

	l.Nodes[0].Label = &Label{
		Rotation: -45,
		Style:    "font-family: Times;font-size: 12;fill: #FF0000;fill-opacity: 1.0;",
		X:        123,
		Y:        123,
	}
	err := l.Update()

	l.Read()
	if err != nil {
		t.Error("Could not update an existing link")
	}
	if l.Nodes[0].Label.Rotation != -45 {
		t.Error("This node label seems to be misconfigured (rotation != -45)")
	}
	if l.Nodes[0].Label.Style != "font-family: Times;font-size: 12;fill: #FF0000;fill-opacity: 1.0;" {
		t.Error("This node label seems to be misconfigured (style)")
	}
	if l.Nodes[0].Label.X != 123 {
		t.Error("This node label seems to be misconfigured (x != 123)")
	}
	if l.Nodes[0].Label.Y != 123 {
		t.Error("This node label seems to be misconfigured (y != 123)")
	}

}

func TestNodeEthernetSwitchLinkDelete(t *testing.T) {
	_, _, _, l := resetTestLinkEthernetSwitchLab2(t)

	if err := l.Delete(); err != nil {
		t.Error("Could not delete an existing node")
		t.Error(err)
	}
}

func TestNodeEthernetSwitchLinkDeleteError(t *testing.T) {
	_, _, _, l := resetTestLinkEthernetSwitchLab2(t)
	l.UUID = "11111111-1111-1111-1111-111111111111"

	if err := l.Delete(); err != nil {
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

func TestNodeVpcsLinkCreate(t *testing.T) {
	p, n1, n2 := resetTestLinkVpcsLab(t)

	createAndTestLink(t, p, n1, n2)
}
