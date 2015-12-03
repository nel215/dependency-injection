package app

import (
	"dependency-injection/infra"
	"testing"
)

func TestManual(t *testing.T) {
	service := infra.NewServiceImpl()
	client := NewClient(service)
	client.Use()
}

func TestContainer(t *testing.T) {
	container := NewContainer()
	client := container.get("client").(*Client)
	client.Use()
}
