package mgenerator

import "testing"

func TestGenStruct(t *testing.T) {
	GenStruct("vex_project", "Project")
}

func TestGenProtoMessage(t *testing.T) {
	GenProtoMessage("vex_project", "ProjectMessage")
}
