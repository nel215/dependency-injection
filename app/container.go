package app

import (
	"dependency-injection/domain"
	"dependency-injection/infra"
)

type Factory func() interface{}

type Container struct {
	factories map[string]Factory
}

func NewContainer() *Container {
	container := &Container{make(map[string]Factory)}
	container.register("service", func() interface{} {
		return infra.NewServiceImpl()
	})
	container.register("client", func() interface{} {
		service := container.get("service").(domain.Service)
		return NewClient(service)
	})
	return container
}

func (self *Container) register(name string, factory Factory) {
	self.factories[name] = factory
}

func (self *Container) get(name string) interface{} {
	return self.factories[name]()
}
