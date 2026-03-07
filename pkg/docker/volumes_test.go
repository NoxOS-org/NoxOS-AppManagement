package docker_test

import (
	"fmt"
	"testing"

	"github.com/Nox-OS/NoxOS-AppManagement/pkg/docker"
)

func TestGetDir(t *testing.T) {
	fmt.Println(docker.GetDir("", "config"))
}
