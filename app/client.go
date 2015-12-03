package app

import (
	"dependency-injection/domain"
)

type Client struct {
	service domain.Service
}

func NewClient(service domain.Service) *Client {
	return &Client{service}
}

func (self *Client) Use() {
	self.service.Execute()
}
