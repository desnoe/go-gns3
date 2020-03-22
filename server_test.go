package gogns3

import (
	"os"
	"strconv"
	"testing"
)

func getTestServer(t *testing.T) *Server {
	hostEnv, ok := os.LookupEnv("GNS3_HOST")
	if !ok {
		t.Error("GNS3_HOST environement variable is not set!")
	}
	portEnv, ok := os.LookupEnv("GNS3_PORT")
	if !ok {
		t.Error("GNS3_PORT environement variable is not set!")
	}
	p, _ := strconv.Atoi(portEnv)
	s := Server{
		Host: hostEnv,
		Port: p,
	}
	return &s
}

func TestServerOK(t *testing.T) {
	s := getTestServer(t)

	if s.Test() != nil {
		t.Error("This server must have answered")
	}
}

func TestServerKO(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping testing in short mode")
	}

	s := getTestServer(t)

	if err := s.Test(); err != nil {
		switch e := err.(type) {
		case *ServerError:
			if e.Status != 1 {
				t.Error(e)
			}
		default:
			t.Error(e)
		}
	}
}

func TestServerGetProjects(t *testing.T) {
	s := getTestServer(t)

	projects, _ := s.GetProjects()
	if s.Test() != nil {
		t.Errorf("%+v\n", projects)
	}
}

func TestServerErrorString(t *testing.T) {
	e := ServerError{
		Status:  1,
		Message: "Test",
	}

	if e.Error() != "Server error #1: Test" {
		t.Error("Error string different than expected")
	}
}
