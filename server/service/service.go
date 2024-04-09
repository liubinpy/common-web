package service

import "common-web/server/service/system"

type Group struct {
	SystemServiceGroup system.ServiceGroup
}

var GroupApp = new(Group)
