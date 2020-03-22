package gogns3

import "encoding/json"

// Label of a node or link
type Label struct {
	Rotation int    `json:"rotation,omitempty"`
	Style    string `json:"style,omitempty"`
	Text     string `json:"text"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}

// NodeProperties are specific to an emulator
type NodeProperties struct {
	PortsMapping []NodeEthernetPortsMapping `json:"ports_mapping,omitempty"`
}

// NodeEthernetPortsMapping are specific to an Ethernet switch
type NodeEthernetPortsMapping struct {
	Name       string `json:"name"`
	PortNumber int    `json:"port_number"`
	Type       string `json:"type"`
	Vlan       int    `json:"vlan"`
	EtherType  string `json:"ethertype,omitempty"`
}

// Node is the basic structure used for a GNS3 node
type Node struct {
	CommandLine     string         `json:"command_line,omitempty"`
	ComputeID       string         `json:"compute_id"`
	Console         int            `json:"console,omitempty"`
	ConsoleType     string         `json:"console_type,omitempty"`
	FirstPortName   string         `json:"first_port_name,omitempty"`
	Label           *Label         `json:"label,omitempty"`
	Name            string         `json:"name"`
	NodeType        string         `json:"node_type"`
	PortNameFormat  string         `json:"port_name_format,omitempty"`
	PortSegmentSize int            `json:"port_segment_size,omitempty"`
	Project         *Project       `json:"-"`
	Properties      NodeProperties `json:"properties,omitempty"`
	Status          string         `json:"status,omitempty"`
	Symbol          string         `json:"symbol,omitempty"`
	UUID            string         `json:"node_id,omitempty"`
	X               int            `json:"x,omitempty"`
	Y               int            `json:"y,omitempty"`
	Z               int            `json:"z,omitempty"`
}

func (n *Node) url() string {
	return n.Project.url() + "/nodes/" + n.UUID
}

// Read reads an existing node in the project
func (n *Node) Read() error {
	// save the project information as it will be reset
	project := n.Project
	nodes, _ := n.Project.GetNodes()
	for _, node := range nodes {
		if node.Name == n.Name {
			*n = node
			// restore the project information
			n.Project = project
			return nil
		}
	}
	return &ServerError{Status: 404, Message: "Node does not exist in the project"}
}

// Exists checks a node exists in the project
func (n *Node) Exists() (bool, error) {
	// Use a new struct because Read() will overwrite it
	node := Node{
		Name:    n.Name,
		Project: n.Project,
	}

	err := node.Read()
	return err == nil, err
}

// Create creates a node in the project
func (n *Node) Create() error {
	b, _ := json.Marshal(n)
	status, content, err := n.Project.Server.HTTPRequest("POST", n.Project.url()+"/nodes", b)
	if !(status >= 200 && status < 300) {
		serverError := ServerError{}
		json.Unmarshal(content, &serverError)
		return &serverError
	}
	json.Unmarshal(content, n)
	return err
}

// Delete deletes a node in the project
// Read() may be called before a Delete() can be executed
func (n *Node) Delete() error {
	if n.UUID == "" {
		n.Read()
	}
	status, content, err := n.Project.Server.HTTPRequest("DELETE", n.url(), nil)
	if !(status >= 200 && status < 300) {
		serverError := ServerError{}
		json.Unmarshal(content, &serverError)
		return &serverError
	}
	return err
}

// Update updates a node in the project
func (n *Node) Update() error {
	var UUID = n.UUID
	n.UUID = ""
	b, _ := json.Marshal(n)

	status, content, err := n.Project.Server.HTTPRequest("PUT", n.url()+UUID, b)
	if !(status >= 200 && status < 300) {
		serverError := ServerError{}
		json.Unmarshal(content, &serverError)
		return &serverError
	}
	json.Unmarshal(content, n)

	return err
}
