package infra

import (
	"dependency-injection/domain"
	"fmt"
)

type ServiceImpl struct {
}

func NewServiceImpl() domain.Service {
	return &ServiceImpl{}
}

func (self *ServiceImpl) Execute() {
	fmt.Println("implement in infra")
}
