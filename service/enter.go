package service

type BaseServiceGroup struct {
	EmailService
	ShortService
	WxPushService
	JwtService
}

var ServiceGroup = new(BaseServiceGroup)
