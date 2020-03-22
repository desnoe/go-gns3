package gogns3

import "encoding/json"

// Link is the basic structure used for a GNS3 link
type Link struct {
	CaptureFileName string     `json:"capture_file_name,omitempty"`
	CaptureFilePath string     `json:"capture_file_path,omitempty"`
	Capturing       bool       `json:"capturing"`
	LinkType        string     `json:"link_type,omitempty"`
	Nodes           []LinkNode `json:"nodes,omitempty"`
	Project         *Project   `json:"-"`
	Suspend         bool       `json:"suspend"`
	UUID            string     `json:"link_id,omitempty"`
}

// LinkNode is the structure used to identify the end of a link
type LinkNode struct {
	AdapterNumber int    `json:"adapter_number"`
	Label         *Label `json:"label,omitempty"`
	NodeID        string `json:"node_id,omitempty"`
	PortNumber    int    `json:"port_number"`
}

func (l *Link) url() string {
	return l.Project.url() + "/links/" + l.UUID
}

// Read reads an existing node in the project
func (l *Link) Read() error {
	// save the project information as it will be reset
	project := l.Project
	links, _ := l.Project.GetLinks()
	for _, link := range links {
		// Links are not named in GNS3, so read will be based on UUID and not name
		if link.UUID == l.UUID {
			*l = link
			// restore the project information
			l.Project = project
			return nil
		}
	}
	return &ServerError{Status: 404, Message: "Link does not exist in the project"}
}

// Exists checks a link exists in the project
func (l *Link) Exists() (bool, error) {
	// Use a new struct because Read() will overwrite it
	link := Link{
		UUID:    l.UUID,
		Project: l.Project,
	}

	err := link.Read()
	return err == nil, err
}

// Create creates a link in the project
func (l *Link) Create() error {
	b, _ := json.Marshal(l)
	status, content, err := l.Project.Server.HTTPRequest("POST", l.Project.url()+"/links", b)
	if !(status >= 200 && status < 300) {
		serverError := ServerError{}
		json.Unmarshal(content, &serverError)
		return &serverError
	}
	json.Unmarshal(content, l)
	return err
}

// Delete deletes a link in the project
func (l *Link) Delete() error {
	status, content, err := l.Project.Server.HTTPRequest("DELETE", l.url(), nil)
	if !(status >= 200 && status < 300) {
		serverError := ServerError{}
		json.Unmarshal(content, &serverError)
		return &serverError
	}
	return err
}

// Update updates a link in the project
func (l *Link) Update() error {
	var UUID = l.UUID
	l.UUID = ""
	b, _ := json.Marshal(l)

	status, content, err := l.Project.Server.HTTPRequest("PUT", l.url()+UUID, b)
	if !(status >= 200 && status < 300) {
		serverError := ServerError{}
		json.Unmarshal(content, &serverError)
		return &serverError
	}
	json.Unmarshal(content, l)

	return err
}
