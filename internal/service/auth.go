package service

import "errors"

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	// 根据ID判断认证信息是否存在
	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist")
}
