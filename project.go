package gogns3

import (
	"encoding/json"
)

// Project is the basic structure used for a GNS3 project
type Project struct {
	AutoClose           bool    `json:"auto_close"`
	Name                string  `json:"name"`
	Server              *Server `json:"-"`
	UUID                string  `json:"project_id,omitempty"`
	SceneHeight         int     `json:"scene_height,omitempty"`
	SceneWidth          int     `json:"scene_width,omitempty"`
	ShowGrid            bool    `json:"show_grid"`
	ShowInterfaceLabels bool    `json:"show_interface_labels"`
	ShowLayers          bool    `json:"show_layers"`
	SnapToGgrid         bool    `json:"snap_to_grid"`
	Zoom                int     `json:"zoom,omitempty"`
}

func (p *Project) url() string {
	return p.Server.url() + string('/') + p.UUID
}

// Read reads an existing project on the server
func (p *Project) Read() error {
	// save the server information as it will be reset
	server := p.Server
	projects, _ := p.Server.GetProjects()
	for _, project := range projects {
		if project.Name == p.Name {
			*p = project
			// restore the server information
			p.Server = server
			return nil
		}
	}
	return &ServerError{Status: 404, Message: "Project does not exist on server"}
}

// Exists checks a project exists on the server
func (p *Project) Exists() (bool, error) {
	// Use a new struct because Read() will overwrite it
	project := Project{
		Name:   p.Name,
		Server: p.Server,
	}

	err := project.Read()
	return err == nil, err
}

// Create creates a project on the server
func (p *Project) Create() error {
	b, _ := json.Marshal(p)
	status, content, err := p.Server.HTTPRequest("POST", p.Server.url(), b)
	if !(status >= 200 && status < 300) {
		serverError := ServerError{}
		json.Unmarshal(content, &serverError)
		return &serverError
	}
	json.Unmarshal(content, p)
	return err
}

// Delete deletes a project on the server
// Read() may be called before a Delete() can be executed
func (p *Project) Delete() error {
	if p.UUID == "" {
		p.Read()
	}
	status, content, err := p.Server.HTTPRequest("DELETE", p.url(), nil)
	if !(status >= 200 && status < 300) {
		serverError := ServerError{}
		json.Unmarshal(content, &serverError)
		return &serverError
	}
	return err
}

// Update updates a project on the server
func (p *Project) Update() error {
	var UUID = p.UUID
	p.UUID = ""
	b, _ := json.Marshal(p)

	status, content, err := p.Server.HTTPRequest("PUT", p.url()+UUID, b)
	if !(status >= 200 && status < 300) {
		serverError := ServerError{}
		json.Unmarshal(content, &serverError)
		return &serverError
	}
	json.Unmarshal(content, p)
	return err
}

// GetNodes gets the list of all nodes of a project
func (p *Project) GetNodes() ([]Node, error) {
	// Send the HTTP request and analyze errors and status code
	status, content, _ := p.Server.HTTPRequest("GET", p.url()+"/nodes", nil)
	if !(status >= 200 && status < 300) {
		serverError := ServerError{}
		json.Unmarshal(content, &serverError)
		return nil, &serverError
	}

	// Unmarshal the JSON-encoded node list
	nodes := []Node{}
	json.Unmarshal(content, &nodes)
	// Set the project for each node
	for idx := range nodes {
		nodes[idx].Project = p
	}

	return nodes, nil
}

// GetLinks gets the list of all links of a project
func (p *Project) GetLinks() ([]Link, error) {
	// Send the HTTP request and analyze errors and status code
	status, content, _ := p.Server.HTTPRequest("GET", p.url()+"/links", nil)
	if !(status >= 200 && status < 300) {
		serverError := ServerError{}
		json.Unmarshal(content, &serverError)
		return nil, &serverError
	}

	// Unmarshal the JSON-encoded node list
	links := []Link{}
	json.Unmarshal(content, &links)
	// Set the project for each node
	for idx := range links {
		links[idx].Project = p
	}

	return links, nil
}
