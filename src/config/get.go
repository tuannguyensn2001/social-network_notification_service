package config

import _const "social-work_notification_service/src/const"

type IConfig interface {
	GetEnvironment() _const.Environment
	GetGrpcAddress() string
}

func (c *config) GetEnvironment() _const.Environment {
	return c.env
}

func (c *config) GetGrpcAddress() string {
	return c.grpcAddress
}
