package system

import "common-web/server/service"

type ApiGroup struct {
	UserApi
}

var (
	userService = service.GroupApp.SystemServiceGroup.UserService
)
