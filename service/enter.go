package service

type BaseGroup struct {
	EmailService
	ShortService
	WxPushService
	JwtService
}

var ServiceGroup = new(BaseGroup)
