package service

import "blog/server/service/system"

type ServiceGroup struct {
	ServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
