package router

import (
	c "github.com/goushuyun/taobao-erp/gateway/controller"
	m "github.com/goushuyun/taobao-erp/gateway/middleware"
)

//SetRouterV1 设置seller的router
func SetRouterV1() *m.Router {
	v1 := m.NewWithPrefix("/v1")

	// users
	v1.Register("/users/register", m.Wrap(c.Register))

	// sms
	v1.Register("/sms/send_sms", m.Wrap(c.SendSms))

	return v1
}
