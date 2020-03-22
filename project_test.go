package gogns3

import (
	"testing"
)

func resetTestProject(t *testing.T) *Project {
	p := Project{
		Name:   "gogns3",
		Server: getTestServer(t),
	}

	// Check it exists and delete it if so
	b, _ := p.Exists()
	if b {
		p.Delete()
	}

	err := p.Create()
	if err != nil {
		t.Error("Could not create test project")
		t.Error(err)
	}

	return &p
}

func TestProjectCreate(t *testing.T) {
	p := Project{
		Name:   "gogns3",
		Server: getTestServer(t),
	}

	// Check it exists and delete it if so
	b, _ := p.Exists()
	if b {
		p.Delete()
	}

	p = Project{
		AutoClose:           true,
		Name:                "gogns3",
		Server:              getTestServer(t),
		SceneHeight:         4000,
		SceneWidth:          4000,
		ShowGrid:            true,
		ShowInterfaceLabels: true,
		ShowLayers:          true,
		SnapToGgrid:         true,
		Zoom:                200,
	}

	err := p.Create()

	if err != nil {
		t.Error("Could not create a new project")
		t.Error(err)
	}
	p.Read()
	if p.AutoClose != true {
		t.Error("This project seems to be misconfigured (auto_close != true)")
	}
	if p.SceneHeight != 4000 {
		t.Error("This project seems to be misconfigured (scene_height != 4000)")
	}
	if p.SceneWidth != 4000 {
		t.Error("This project seems to be misconfigured (scene_width != 4000)")
	}
	if p.ShowGrid != true {
		t.Error("This project seems to be misconfigured (show_grid != true)")
	}
	if p.ShowInterfaceLabels != true {
		t.Error("This project seems to be misconfigured (show_interface_labels != true)")
	}
	if p.ShowLayers != true {
		t.Error("This project seems to be misconfigured (show_layers != true)")
	}
	if p.SnapToGgrid != true {
		t.Error("This project seems to be misconfigured (snap_to_grid != true)")
	}
	if p.Zoom != 200 {
		t.Error("This project seems to be misconfigured (zoom != 200)")
	}
}

func TestProjectCreateError(t *testing.T) {
	p := Project{
		Name:   "",
		Server: getTestServer(t),
	}

	if err := p.Create(); err != nil {
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

func TestProjectExists(t *testing.T) {
	p := *resetTestProject(t)

	b, err := p.Exists()
	if err != nil {
		t.Error(err)
	}
	if !b {
		t.Error("This project must exist on the server")
	}
	if p.SceneHeight <= 0 {
		t.Error("This project seems to be misconfigured (scene_height <= 0)")
	}
	if p.SceneWidth <= 0 {
		t.Error("This project seems to be misconfigured (scene_width <= 0)")
	}
}

func TestProjectExistsError(t *testing.T) {
	p := Project{
		Name:   "fakefakefake",
		Server: getTestServer(t),
	}

	if _, err := p.Exists(); err != nil {
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

func TestProjectUpdate(t *testing.T) {
	p := *resetTestProject(t)

	p.AutoClose = false
	p.SceneHeight = 1234
	p.SceneWidth = 1234
	p.ShowGrid = false
	p.ShowInterfaceLabels = false
	p.ShowLayers = false
	p.SnapToGgrid = false
	p.Zoom = 123
	err := p.Update()

	if err != nil {
		t.Error("Could not update an existing project")
	}
	if p.AutoClose != false {
		t.Error("This project seems to be misconfigured (auto_close != false)")
	}
	if p.SceneHeight != 1234 {
		t.Error("This project seems to be misconfigured (scene_height != 1234)")
	}
	if p.SceneWidth != 1234 {
		t.Error("This project seems to be misconfigured (scene_width != 1234)")
	}
	if p.ShowGrid != false {
		t.Error("This project seems to be misconfigured (show_grid != false)")
	}
	if p.ShowInterfaceLabels != false {
		t.Error("This project seems to be misconfigured (show_interface_labels != false)")
	}
	if p.ShowLayers != false {
		t.Error("This project seems to be misconfigured (show_layers != false)")
	}
	if p.SnapToGgrid != false {
		t.Error("This project seems to be misconfigured (snap_to_grid != false)")
	}
	if p.Zoom != 123 {
		t.Error("This project seems to be misconfigured (zoom != 123)")
	}

}

func TestProjectUpdateError(t *testing.T) {
	p := *resetTestProject(t)

	p.Name = ""
	err := p.Update()

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

func TestProjectDelete(t *testing.T) {
	p := *resetTestProject(t)
	// clear the UUID to make sure a Read is automatically performed before
	p.UUID = ""

	if err := p.Delete(); err != nil {
		t.Error("Could not delete an existing project")
	}
}

func TestProjectDeleteError(t *testing.T) {
	p := *resetTestProject(t)
	p.UUID = "11111111-1111-1111-1111-111111111111"

	if err := p.Delete(); err != nil {
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

func TestProjectGetNodes(t *testing.T) {
	p := *resetTestProject(t)

	_, err := p.GetNodes()
	if err != nil {
		t.Error(err)
	}
}

func TestProjectGetNodesError(t *testing.T) {
	p := *resetTestProject(t)
	p.UUID = "11111111-1111-1111-1111-111111111111"

	if _, err := p.GetNodes(); err != nil {
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

func TestProjectGetLinks(t *testing.T) {
	p := *resetTestProject(t)

	_, err := p.GetLinks()
	if err != nil {
		t.Error(err)
	}
}

func TestProjectGetLinksError(t *testing.T) {
	p := *resetTestProject(t)
	p.UUID = "11111111-1111-1111-1111-111111111111"

	if _, err := p.GetLinks(); err != nil {
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
