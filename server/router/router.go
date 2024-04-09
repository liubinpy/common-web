package router

import "common-web/server/router/system"

type Group struct {
	System system.RouterGroup
}

var GroupApp = new(Group)
