package gogns3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Server is a basic structure describing a GNS3 server
type Server struct {
	Host string
	Port int
}

func (s *Server) url() string {
	return "http://" + s.Host + ":" + strconv.Itoa(s.Port) + "/v2/projects"
}

// ServerError is the basic structure used for server related errors
type ServerError struct {
	Message string
	Method  string
	Path    string
	Request interface{}
	Status  int
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("Server error #%s: %s", strconv.Itoa(e.Status), e.Message)
}

// HTTPRequest executes any HTTP request to the server.
// Method, URL and body must be provided.
func (s *Server) HTTPRequest(method string, url string, body []byte) (int, []byte, error) {
	client := http.Client{Timeout: 5 * time.Second}
	req, _ := http.NewRequest(method, url, bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	resp, _ := client.Do(req)
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return resp.StatusCode, content, err
}

// Test is a simple HTTP GET request to the server to check it is alive
func (s *Server) Test() error {
	_, _, err := s.HTTPRequest("GET", s.url(), nil)
	return err
}

// GetProjects gets the list of all projects on the server
func (s *Server) GetProjects() ([]Project, error) {
	// Send the HTTP request and analyze errors and status code
	_, content, _ := s.HTTPRequest("GET", s.url(), nil)

	// Unmarshal the JSON-encoded project list
	projects := []Project{}
	json.Unmarshal(content, &projects)
	// Set the server for each project
	for idx := range projects {
		projects[idx].Server = s
	}

	return projects, nil
}
