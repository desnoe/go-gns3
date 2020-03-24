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

// NodeProperties are specific to an emulator. For some emulators, some fields
// are required, some are forbidden and it depends on the node type. I can't see
// a neater solution than duplicating the type in order to specify which field
// must be exported and which field must not. This way, we only have a
// marshalling problem with some custom MarshalJSON() functions, but nothing
// special to do when unmarshalling. Note than the omitempty keyword is useless
// with bool type as missing value does not mean false for the GNS3 server API.
type NodeProperties struct {
	nodeType           string
	AcpiShutdown       bool                       `json:"acpi_shutdown"`
	AdapterType        string                     `json:"adapter_type"`
	Adapters           int                        `json:"adapters"`
	BiosImage          string                     `json:"bios_image"`
	BiosImageMd5sum    string                     `json:"bios_image_md5sum"`
	BootPriority       string                     `json:"boot_priority"`
	CdromImage         string                     `json:"cdrom_image"`
	CdromImageMd5sum   string                     `json:"cdrom_image_md5sum"`
	CPUThrottling      int                        `json:"cpu_throttling"`
	CPUs               int                        `json:"cpus"`
	HdaDiskImage       string                     `json:"hda_disk_image"`
	HdaDiskImageMd5sum string                     `json:"hda_disk_image_md5sum"`
	HdaDiskInterface   string                     `json:"hda_disk_interface"`
	HdbDiskImage       string                     `json:"hdb_disk_image"`
	HdbDiskImageMd5sum string                     `json:"hdb_disk_image_md5sum"`
	HdbDiskInterface   string                     `json:"hdb_disk_interface"`
	HdcDiskImage       string                     `json:"hdc_disk_image"`
	HdcDiskImageMd5sum string                     `json:"hdc_disk_image_md5sum"`
	HdcDiskInterface   string                     `json:"hdc_disk_interface"`
	HddDiskImage       string                     `json:"hdd_disk_image"`
	HddDiskImageMd5sum string                     `json:"hdd_disk_image_md5sum"`
	HddDiskInterface   string                     `json:"hdd_disk_interface"`
	Initrd             string                     `json:"initrd"`
	InitrdMd5sum       string                     `json:"initrd_md5sum"`
	KernelCommandLine  string                     `json:"kernel_command_line"`
	KernelImage        string                     `json:"kernel_image"`
	KernelImageMd5sum  string                     `json:"kernel_image_md5sum"`
	LegacyNetworking   bool                       `json:"legacy_networking"`
	MacAddress         string                     `json:"mac_address"`
	Options            string                     `json:"options"`
	PortsMapping       []NodeEthernetPortsMapping `json:"ports_mapping"`
	Platform           string                     `json:"platform"`
	ProcessPriority    string                     `json:"process_priority"`
	QemuPath           string                     `json:"qemu_path"`
	RAM                int                        `json:"ram"`
	Usage              string                     `json:"usage"`
}

type nodeEthernetSwitchProperties struct {
	nodeType           string
	AcpiShutdown       bool                       `json:"-"`
	AdapterType        string                     `json:"-"`
	Adapters           int                        `json:"-"`
	BiosImage          string                     `json:"-"`
	BiosImageMd5sum    string                     `json:"-"`
	BootPriority       string                     `json:"-"`
	CdromImage         string                     `json:"-"`
	CdromImageMd5sum   string                     `json:"-"`
	CPUThrottling      int                        `json:"-"`
	CPUs               int                        `json:"-"`
	HdaDiskImage       string                     `json:"-"`
	HdaDiskImageMd5sum string                     `json:"-"`
	HdaDiskInterface   string                     `json:"-"`
	HdbDiskImage       string                     `json:"-"`
	HdbDiskImageMd5sum string                     `json:"-"`
	HdbDiskInterface   string                     `json:"-"`
	HdcDiskImage       string                     `json:"-"`
	HdcDiskImageMd5sum string                     `json:"-"`
	HdcDiskInterface   string                     `json:"-"`
	HddDiskImage       string                     `json:"-"`
	HddDiskImageMd5sum string                     `json:"-"`
	HddDiskInterface   string                     `json:"-"`
	Initrd             string                     `json:"-"`
	InitrdMd5sum       string                     `json:"-"`
	KernelCommandLine  string                     `json:"-"`
	KernelImage        string                     `json:"-"`
	KernelImageMd5sum  string                     `json:"-"`
	LegacyNetworking   bool                       `json:"-"`
	MacAddress         string                     `json:"-"`
	Options            string                     `json:"-"`
	PortsMapping       []NodeEthernetPortsMapping `json:"ports_mapping,omitempty"`
	Platform           string                     `json:"-"`
	ProcessPriority    string                     `json:"-"`
	QemuPath           string                     `json:"-"`
	RAM                int                        `json:"-"`
	Usage              string                     `json:"-"`
}

type nodeQemuProperties struct {
	nodeType           string
	AcpiShutdown       bool                       `json:"acpi_shutdown,omitempty"`
	AdapterType        string                     `json:"adapter_type,omitempty"`
	Adapters           int                        `json:"adapters,omitempty"`
	BiosImage          string                     `json:"bios_image,omitempty"`
	BiosImageMd5sum    string                     `json:"bios_image_md5sum,omitempty"`
	BootPriority       string                     `json:"boot_priority,omitempty"`
	CdromImage         string                     `json:"cdrom_image,omitempty"`
	CdromImageMd5sum   string                     `json:"cdrom_image_md5sum,omitempty"`
	CPUThrottling      int                        `json:"cpu_throttling,omitempty"`
	CPUs               int                        `json:"cpus,omitempty"`
	HdaDiskImage       string                     `json:"hda_disk_image,omitempty"`
	HdaDiskImageMd5sum string                     `json:"hda_disk_image_md5sum,omitempty"`
	HdaDiskInterface   string                     `json:"hda_disk_interface,omitempty"`
	HdbDiskImage       string                     `json:"hdb_disk_image,omitempty"`
	HdbDiskImageMd5sum string                     `json:"hdb_disk_image_md5sum,omitempty"`
	HdbDiskInterface   string                     `json:"hdb_disk_interface,omitempty"`
	HdcDiskImage       string                     `json:"hdc_disk_image,omitempty"`
	HdcDiskImageMd5sum string                     `json:"hdc_disk_image_md5sum,omitempty"`
	HdcDiskInterface   string                     `json:"hdc_disk_interface,omitempty"`
	HddDiskImage       string                     `json:"hdd_disk_image,omitempty"`
	HddDiskImageMd5sum string                     `json:"hdd_disk_image_md5sum,omitempty"`
	HddDiskInterface   string                     `json:"hdd_disk_interface,omitempty"`
	Initrd             string                     `json:"initrd,omitempty"`
	InitrdMd5sum       string                     `json:"initrd_md5sum,omitempty"`
	KernelCommandLine  string                     `json:"kernel_command_line,omitempty"`
	KernelImage        string                     `json:"kernel_image,omitempty"`
	KernelImageMd5sum  string                     `json:"kernel_image_md5sum,omitempty"`
	LegacyNetworking   bool                       `json:"legacy_networking,omitempty"`
	MacAddress         string                     `json:"mac_address,omitempty"`
	Options            string                     `json:"options,omitempty"`
	PortsMapping       []NodeEthernetPortsMapping `json:"-"`
	Platform           string                     `json:"platform,omitempty"`
	ProcessPriority    string                     `json:"process_priority,omitempty"`
	QemuPath           string                     `json:"qemu_path,omitempty"`
	RAM                int                        `json:"ram,omitempty"`
	Usage              string                     `json:"usage,omitempty"`
}

type nodeVpcsProperties struct {
	nodeType           string
	AcpiShutdown       bool                       `json:"-"`
	AdapterType        string                     `json:"-"`
	Adapters           int                        `json:"-"`
	BiosImage          string                     `json:"-"`
	BiosImageMd5sum    string                     `json:"-"`
	BootPriority       string                     `json:"-"`
	CdromImage         string                     `json:"-"`
	CdromImageMd5sum   string                     `json:"-"`
	CPUThrottling      int                        `json:"-"`
	CPUs               int                        `json:"-"`
	HdaDiskImage       string                     `json:"-"`
	HdaDiskImageMd5sum string                     `json:"-"`
	HdaDiskInterface   string                     `json:"-"`
	HdbDiskImage       string                     `json:"-"`
	HdbDiskImageMd5sum string                     `json:"-"`
	HdbDiskInterface   string                     `json:"-"`
	HdcDiskImage       string                     `json:"-"`
	HdcDiskImageMd5sum string                     `json:"-"`
	HdcDiskInterface   string                     `json:"-"`
	HddDiskImage       string                     `json:"-"`
	HddDiskImageMd5sum string                     `json:"-"`
	HddDiskInterface   string                     `json:"-"`
	Initrd             string                     `json:"-"`
	InitrdMd5sum       string                     `json:"-"`
	KernelCommandLine  string                     `json:"-"`
	KernelImage        string                     `json:"-"`
	KernelImageMd5sum  string                     `json:"-"`
	LegacyNetworking   bool                       `json:"-"`
	MacAddress         string                     `json:"-"`
	Options            string                     `json:"-"`
	PortsMapping       []NodeEthernetPortsMapping `json:"-"`
	Platform           string                     `json:"-"`
	ProcessPriority    string                     `json:"-"`
	QemuPath           string                     `json:"-"`
	RAM                int                        `json:"-"`
	Usage              string                     `json:"-"`
}

// MarshalJSON allows to customize the JSON output The design of the GNS3 server
// does not permit unncessary fields, we thus have to filter them. For each node
// type the properties structure is defined multiple times and the right cast
// into the relevant specific type is done at marshalling time.
func (p NodeProperties) MarshalJSON() ([]byte, error) {
	switch p.nodeType {
	case "ethernet_switch":
		return json.Marshal(nodeEthernetSwitchProperties(p))
	case "qemu":
		return json.Marshal(nodeQemuProperties(p))
	case "vpcs":
		return json.Marshal(nodeVpcsProperties(p))
	}
	return json.Marshal(nil)
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

type nodeAlias Node

// MarshalJSON allows to copy the nodeType of the Node structure to the child
// NodeProperties structure that will need this information to select the right
// object type for casting. The nodeAlias prevents infinite loop recursions.
func (n Node) MarshalJSON() ([]byte, error) {
	n.Properties.nodeType = n.NodeType
	return json.Marshal(nodeAlias(n))
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
