package system

import (
	"common-web/server/global"
	"common-web/server/model/system"
)

type UserService struct {
}

func (u *UserService) FindUserByAccount(account string) (*system.User, error) {
	var user system.User
	err := global.CommonMysql.Where("account = ?", account).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
